package util

import (
	"errors"
	"fmt"
	"time"
	pb "xchain-sdk-go/src/xchain/proto"
)

func RunWithTimeout(f func(req *pb.ContractSpec) (*pb.RpcReply, error), funcName string,
	req *pb.ContractSpec, t time.Duration) (*pb.RpcReply, error) {
	done := make(chan struct{})
	var resp *pb.RpcReply
	var err error
	go func() {
		resp, err = f(req)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return resp, err
	case <-time.After(t):
		return nil, errors.New(fmt.Sprintf("call method:%s timeout", funcName))
	}
}

func RunWithTimeout2(f func(req *pb.SdkPeerRequest) (*pb.RpcReply, error), funcName string,
	req *pb.SdkPeerRequest, t time.Duration) (*pb.RpcReply, error) {
	done := make(chan struct{})
	var resp *pb.RpcReply
	var err error
	go func() {
		resp, err = f(req)
		done <- struct{}{}
	}()

	select {
	case <-done:
		return resp, err
	case <-time.After(t):
		return nil, errors.New(fmt.Sprintf("call method:%s timeout", funcName))
	}
}
