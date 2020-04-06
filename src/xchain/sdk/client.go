package sdk

import (
	"io/ioutil"
	"log"
	"path/filepath"
	"xchain-sdk-go/src/xchain/config"
	"xchain-sdk-go/src/xchain/reply"
	"xchain-sdk-go/src/xchain/util"
	"xchain-sdk-go/src/xchain/vo"
)

type Client struct {
	channel *Channel
}

func NewClient(channel *Channel) *Client {
	return &Client{channel: channel}
}

const (
	paramIsNil     = "param is empty"
	failedToInvoke = "failed to invoke"
)

func (c *Client) InitContractWithResult(identity, version, path string) *vo.Response {
	if c.channel == nil {
		log.Printf("wrong, channel is nil")
		reply.ResponseError(paramIsNil)
	}
	if identity == "" || version == "" || path == "" {
		log.Printf("wrong, params is empty")
		reply.ResponseError(paramIsNil)
	}
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("read %s wrong, err:%v", path, err)
		reply.ResponseError(failedToInvoke)
	}
	filename := filepath.Base(path)
	contractId := util.BuildContractID(identity, version, config.Golang, false)
	contractSpec := util.BuildContractSpec(contractId, filename, bytes)
	resp, err := c.channel.initContractWithResult(contractSpec)
	if err != nil {
		log.Printf("initContractWithResult %s-%s wrong, err:%v", identity, version, err)
		reply.ResponseError(failedToInvoke + ", err:" + err.Error())
	}
	return reply.BuildResponse(resp)
}

func (c *Client) Invoke(identity, version, method, params string) *vo.Response {
	if c.channel == nil {
		log.Printf("wrong, channel is nil")
		reply.ResponseError(paramIsNil)
	}
	if identity == "" || version == "" || method == "" || params == "" {
		log.Printf("wrong, params is empty")
		reply.ResponseError(paramIsNil)
	}
	contractId := util.BuildContractID(identity, version, config.Golang, false)
	req := util.BuildSdkPeerRequest(contractId, method, params)
	resp, err := c.channel.invoke(req)
	if err != nil {
		log.Printf("invoke %s-%s wrong, err:%v", identity, version, err)
		reply.ResponseError(failedToInvoke + ", err:" + err.Error())
	}
	return reply.BuildResponse(resp)
}

func (c *Client) InvokeSystemContract(identity, version, method, params string) *vo.Response {
	if c.channel == nil {
		log.Printf("wrong, channel is nil")
		reply.ResponseError(paramIsNil)
	}
	if identity == "" || version == "" || method == "" || params == "" {
		log.Printf("wrong, params is empty")
		reply.ResponseError(paramIsNil)
	}
	contractId := util.BuildContractID(identity, version, config.Java, true)
	req := util.BuildSdkPeerRequest(contractId, method, params)
	resp, err := c.channel.invoke(req)
	if err != nil {
		log.Printf("invoke %s-%s wrong, err:%v", identity, version, err)
		reply.ResponseError(failedToInvoke + ", err:" + err.Error())
	}
	return reply.BuildResponse(resp)
}
