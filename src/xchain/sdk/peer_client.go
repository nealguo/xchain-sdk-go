package sdk

import (
	"context"
	"errors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
	"xchain-sdk-go/src/xchain/config"
	pb "xchain-sdk-go/src/xchain/proto"
)

type PeerClient struct {
	node   *config.Node
	client pb.SdkPeerServiceClient
}

func NewPeerClient(node *config.Node) (*PeerClient, error) {
	addr := node.Address()
	if addr == "" {
		msg := "failed to get peer addr, addr is empty"
		log.Printf(msg)
		return nil, errors.New(msg)
	}
	var conn *grpc.ClientConn
	if config.SslConf.SslEnable {
		// serverNameOverride使用RSA证书生成时填写的"Common Name"值
		creds, err := credentials.NewClientTLSFromFile(config.SslConf.SslCertFilePath, "xchain")
		if err != nil {
			log.Printf("failed to create SSL credentials:%v", err)
			return nil, err
		}
		conn, err = grpc.Dial(addr, grpc.WithTransportCredentials(creds))
		if err != nil {
			log.Printf("failed to connect to peer, addr:%s, err:%v", addr, err)
			return nil, err
		}
	} else {
		var err error
		conn, err = grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Printf("failed to connect to peer, addr:%s, err:%v", addr, err)
			return nil, err
		}
	}
	client := pb.NewSdkPeerServiceClient(conn)
	peerClient := &PeerClient{
		node:   node,
		client: client,
	}
	log.Printf("client started, connected to peer:%s With SSL", addr)
	return peerClient, nil
}

func (c *PeerClient) InitContract(req *pb.ContractSpec) (*pb.RpcReply, error) {
	return c.client.InitContract(context.Background(), req)
}

func (c *PeerClient) StopContract(req *pb.SdkPeerRequest) (*pb.RpcReply, error) {
	return c.client.StopContract(context.Background(), req)
}
func (c *PeerClient) RemoveContract(req *pb.SdkPeerRequest) (*pb.RpcReply, error) {
	return c.client.RemoveContract(context.Background(), req)
}

func (c *PeerClient) Invoke(req *pb.SdkPeerRequest) (*pb.RpcReply, error) {
	return c.client.Invoke(context.Background(), req)
}

func (c *PeerClient) Query(req *pb.SdkPeerRequest) (*pb.RpcReply, error) {
	return c.client.Query(context.Background(), req)
}
