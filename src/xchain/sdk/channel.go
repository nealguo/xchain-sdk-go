package sdk

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
	"xchain-sdk-go/src/xchain/config"
	"xchain-sdk-go/src/xchain/model"
	pb "xchain-sdk-go/src/xchain/proto"
	"xchain-sdk-go/src/xchain/reply"
	"xchain-sdk-go/src/xchain/util"
)

const (
	NextLine          = "]\r\n"
	WaitResultTimeout = 30 * time.Second
)

type Channel struct {
	sdkPeerClients  []*PeerClient
	sdkOrderClients []*OrderClient
}

func NewChannel(channel, consensus, order, peer, key, appId, algorithm string,
	ssl bool, sslCert, sslKey, sslTrust string) *Channel {
	// check
	if channel == "" || consensus == "" {
		log.Fatalf("NewChannel wrong, channel config is emtpy")
	}
	if order == "" || peer == "" {
		log.Fatalf("NewChannel wrong, order or peer config is emtpy")
	}
	if key == "" || appId == "" {
		log.Fatalf("NewChannel wrong, key or appId of sdk config is emtpy")
	}
	if algorithm == "" {
		log.Fatalf("NewChannel wrong, algorithm for signing is empty, now support SM2 and RSA")
	}

	// split order and peer
	orders := util.SplitToNodes(order)
	if orders == nil || len(orders) == 0 {
		log.Fatalf("NewChannel wrong, order config is wrong, order:%s", order)
	}
	peers := util.SplitToNodes(peer)
	if peers == nil || len(peers) == 0 {
		log.Fatalf("NewChannel wrong, peer config is wrong, order:%s", order)
	}

	// init config first
	config.InitConfigs(channel, consensus, order, peer, key, appId, algorithm, ssl, sslCert, sslKey, sslTrust)

	// init channel
	var peerClients []*PeerClient
	for _, peer := range peers {
		client, err := NewPeerClient(peer)
		if err != nil {
			log.Fatalf("NewChannel wrong, NewPeerClient err:%v", err)
		}
		peerClients = append(peerClients, client)
	}
	var orderClients []*OrderClient
	for _, order := range orders {
		client, err := NewOrderClient(order)
		if err != nil {
			log.Fatalf("NewChannel wrong, NewOrderClient err:%v", err)
		}
		orderClients = append(orderClients, client)
	}
	return &Channel{
		sdkPeerClients:  peerClients,
		sdkOrderClients: orderClients,
	}
}

func (ch *Channel) initContractWithResult(req *pb.ContractSpec) (*pb.RpcReply, error) {
	clients := ch.sdkPeerClients
	if clients == nil || len(clients) == 0 {
		msg := "peer client dose not exist"
		return reply.RpcReplyError(msg), errors.New(msg)
	}
	buf := bytes.NewBufferString("")
	wg := sync.WaitGroup{}
	wg.Add(len(clients))
	for _, client := range clients {
		go func(wg *sync.WaitGroup, buf *bytes.Buffer) {
			resp, err := util.RunWithTimeout(client.InitContract, "InitContract", req, WaitResultTimeout)
			if err != nil {
				log.Printf("init contract wrong, err:%v", err)
				buf.WriteString("等待peer[")
				buf.WriteString(client.node.Address())
				buf.WriteString("]返回结果时出现异常，异常信息[")
				buf.WriteString(err.Error())
				buf.WriteString(NextLine)
			}
			if resp == nil || resp.GetCode() != config.Success {
				buf.WriteString("peer[")
				buf.WriteString(client.node.Address())
				buf.WriteString("]初始化合约失败，失败信息[")
				buf.WriteString(resp.GetMessage())
				buf.WriteString(NextLine)
			}
			wg.Done()
		}(&wg, buf)
	}
	wg.Wait()

	errMsg := buf.String()
	if errMsg != "" {
		return reply.RpcReplyError(errMsg), errors.New(errMsg)
	}
	return reply.RpcReplySuccess(""), nil
}

func (ch *Channel) removeContractWithResult(req *pb.SdkPeerRequest) (*pb.RpcReply, error) {
	clients := ch.sdkPeerClients
	if clients == nil || len(clients) == 0 {
		msg := "peer client dose not exist"
		return reply.RpcReplyError(msg), errors.New(msg)
	}
	buf := bytes.NewBufferString("")
	wg := sync.WaitGroup{}
	wg.Add(len(clients))
	for _, client := range clients {
		go func(wg *sync.WaitGroup, buf *bytes.Buffer) {
			resp, err := util.RunWithTimeout2(client.RemoveContract, "RemoveContract", req, WaitResultTimeout)
			if err != nil {
				log.Printf("init contract wrong, err:%v", err)
				buf.WriteString("等待peer[")
				buf.WriteString(client.node.Address())
				buf.WriteString("]卸载合约结果时出现异常，异常信息[")
				buf.WriteString(err.Error())
				buf.WriteString(NextLine)
			}
			if resp == nil || resp.GetCode() != config.Success {
				buf.WriteString("peer[")
				buf.WriteString(client.node.Address())
				buf.WriteString("]卸载合约失败，失败信息[")
				buf.WriteString(resp.GetMessage())
				buf.WriteString(NextLine)
			}
			wg.Done()
		}(&wg, buf)
	}
	wg.Wait()

	errMsg := buf.String()
	if errMsg != "" {
		return reply.RpcReplyError(errMsg), errors.New(errMsg)
	}
	return reply.RpcReplySuccess(""), nil
}

func (ch *Channel) invoke(req *pb.SdkPeerRequest) (*pb.RpcReply, error) {
	resp := ch.handlePeerResult(req)
	if resp != nil && resp.GetCode() == config.Success {
		resp = ch.sendToOrder(resp)
	}
	return resp, nil
}

func (ch *Channel) sendToOrder(resp *pb.RpcReply) *pb.RpcReply {
	if resp == nil {
		msg := "sendToOrder wrong, resp is nil"
		log.Printf(msg)
		return reply.RpcReplyError(msg)
	}
	payload := resp.GetPayload()
	if payload == "" {
		msg := "sendToOrder wrong, payload in resp is empty"
		log.Printf(msg)
		return reply.RpcReplyError(msg)
	}
	var responseModel model.ResponseModel
	err := json.Unmarshal([]byte(payload), &responseModel)
	if err != nil {
		msg := "sendToOrder wrong when deserializing json"
		log.Printf(msg)
		return reply.RpcReplyError(msg)
	}
	tx := responseModel.Transaction
	if tx != nil && tx.Writes != nil && len(tx.Writes) > 0 {
		//solo不发送给Order
		if config.ChannelConf.Consensus != config.Solo {
			orderClient := PickOrderClient(ch.sdkOrderClients)
			if orderClient == nil {
				msg := "sendToOrder wrong, order client is nil"
				log.Printf(msg)
				return reply.RpcReplyError(msg)
			}
			req := util.BuildSdkOrderRequest(payload)
			resp2, err := orderClient.SendTransaction(req)
			if err != nil {
				msg := fmt.Sprintf("sendToOrder wrong, when sending tx to order, err:%v", err)
				log.Printf(msg)
				return reply.RpcReplyError(msg)
			}
			if resp2 != nil && resp2.GetCode() != config.Success {
				return reply.RpcReply(resp2.GetCode(), resp2.GetMessage(), responseModel.Value, tx.Hash)
			}
		}
		return reply.RpcReply(resp.GetCode(), resp.GetMessage(), responseModel.Value, tx.Hash)
	}
	return reply.RpcReply(resp.GetCode(), resp.GetMessage(), responseModel.Value, "")
}

func (ch *Channel) handlePeerResult(req *pb.SdkPeerRequest) *pb.RpcReply {
	consensus := config.ChannelConf.Consensus
	switch consensus {
	case config.Raft, config.Pbft:
		peerClient := PickPeerClient(ch.sdkPeerClients)
		if peerClient == nil {
			msg := fmt.Sprintf("handlePeerResult wrong, peer client is nil, consensus:%s", consensus)
			log.Printf(msg)
			return reply.RpcReplyError(msg)
		}
		resp, err := peerClient.Invoke(req)
		if err != nil {
			msg := fmt.Sprintf("handlePeerResult wrong, when invoking contract, consensus:%s, err:%v", consensus, err)
			log.Printf(msg)
			return reply.RpcReplyError(msg)
		}
		if resp != nil && resp.GetCode() == config.Success {
			var responseModel model.ResponseModel
			err := json.Unmarshal([]byte(resp.GetPayload()), &responseModel)
			if err != nil {
				msg := "handlePeerResult wrong when deserializing json"
				log.Printf(msg)
				return reply.RpcReplyError(msg)
			}
			tx := responseModel.Transaction
			endorseSign := util.Sign(util.ToJson(tx.Writes))
			tx.EndorsePeerName = config.AppConf.AppId
			tx.EndorseSign = endorseSign
			resp.Payload = util.ToJson(responseModel)
			return resp
		} else {
			return resp
		}
	case config.Solo:
		peerClient := PickPeerClient(ch.sdkPeerClients)
		if peerClient == nil {
			msg := fmt.Sprintf("handlePeerResult wrong, peer client is nil, consensus:%s", consensus)
			log.Printf(msg)
			return reply.RpcReplyError(msg)
		}
		resp, err := peerClient.Invoke(req)
		if err != nil {
			msg := fmt.Sprintf("handlePeerResult wrong, when invoking contract, consensus:%s, err:%v", consensus, err)
			log.Printf(msg)
			return reply.RpcReplyError(msg)
		}
		return resp
	case config.Kafka:
		successList, failList := ch.invokeAll(req)
		if len(successList)*2 <= len(ch.sdkPeerClients) {
			errMsg := failList[0].GetMessage()
			msg := fmt.Sprintf("peer预共识失败，只获得%s的peer节点响应成功，预共识失败节点返回错误信息：%s", len(successList), errMsg)
			return reply.RpcReplyError(msg)
		}
		resp := getSameWritesReply(successList, 0, len(ch.sdkPeerClients))
		if resp == nil {
			return reply.RpcReplyError("peer预共识写集校验失败")
		}
		return resp
	}
	msg := fmt.Sprintf("unsupported consensus:%s", consensus)
	log.Printf(msg)
	return reply.RpcReplyError(msg)
}

func (ch *Channel) invokeAll(req *pb.SdkPeerRequest) ([]*pb.RpcReply, []*pb.RpcReply) {
	successList := make([]*pb.RpcReply, 0)
	failList := make([]*pb.RpcReply, 0)

	clients := ch.sdkPeerClients
	wg := sync.WaitGroup{}
	wg.Add(len(clients))

	for _, client := range clients {
		go func(wg *sync.WaitGroup) {
			resp, err := client.Invoke(req)
			if err != nil {
				log.Printf("invoke contract wrong, err:%v", err)
				failList = append(failList, resp)
			}
			if resp == nil || resp.GetCode() != config.Success {
				failList = append(failList, resp)
			} else {
				successList = append(successList, resp)
			}
			wg.Done()
		}(&wg)
	}
	wg.Wait()

	return successList, failList
}

// 收集写集，判断是否有过半peer返回的交易写集一致
func getSameWritesReply(successList []*pb.RpcReply, beginIndex, peers int) *pb.RpcReply {
	if beginIndex == len(successList) {
		return nil
	}
	resp := successList[beginIndex]
	if len(successList) == 1 {
		return resp
	}
	payload := resp.GetPayload()
	// 为了省去json序列化和反序列化的时间，这里直接比对交易写集的json字符串是否相等
	if payload != "" && checkMoreThanHalf(successList, payload, beginIndex, peers) {
		return resp
	}
	return getSameWritesReply(successList, beginIndex+1, peers)
}

// 检查预共识节点是否过半
func checkMoreThanHalf(successList []*pb.RpcReply, payload string, beginIndex, peers int) bool {
	sameCount := 1
	writes := getWritesJson(payload)
	for i := beginIndex + 1; i < len(successList); i++ {
		nextResp := successList[i]
		nextPayload := nextResp.GetPayload()
		if nextPayload != "" {
			nextWrites := getWritesJson(nextPayload)
			if writes != "" && writes == nextWrites {
				sameCount++
			}
			if sameCount*2 > peers {
				return true
			}
		}
	}
	return false
}

// 获取交易中写集的json字符串
func getWritesJson(payload string) string {
	// 交易中写集的属性名称对应的json字符串
	writesAttr := "\"writes\":"
	jsonList := strings.Split(payload, writesAttr)
	// 如果没有写集，则是查询，直接返回
	if len(jsonList) < 2 {
		return payload
	}
	index := strings.Index(jsonList[1], "]")
	return jsonList[1][0 : index+1]
}
