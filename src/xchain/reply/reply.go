package reply

import (
	"xchain-sdk-go/src/xchain/config"
	pb "xchain-sdk-go/src/xchain/proto"
	"xchain-sdk-go/src/xchain/vo"
)

func NewResponse(status int, msg, payload string) *vo.Response {
	return &vo.Response{
		Code:    status,
		Message: msg,
		Payload: payload,
	}
}

func ResponseSuccess(payload string) *vo.Response {
	return NewResponse(config.Success, "success", payload)
}

func ResponseError(msg string) *vo.Response {
	return NewResponse(config.Error, msg, "payload")
}

func BuildResponse(resp *pb.RpcReply) *vo.Response {
	var payload string
	hash := resp.GetTransactionHash()
	if hash == "" {
		payload = resp.GetPayload()
	} else {
		payload = hash
	}
	return &vo.Response{
		Code:    int(resp.GetCode()),
		Message: resp.GetMessage(),
		Payload: payload,
	}
}

func RpcReply(status int32, msg, payload, hash string) *pb.RpcReply {
	return &pb.RpcReply{
		Code:            status,
		Message:         msg,
		Payload:         payload,
		TransactionHash: hash,
		File:            nil}
}

func RpcReplySuccess(payload string) *pb.RpcReply {
	return RpcReply(config.Success, "success", payload, "")
}

func RpcReplyError(msg string) *pb.RpcReply {
	return RpcReply(config.Error, msg, "payload", "")
}
