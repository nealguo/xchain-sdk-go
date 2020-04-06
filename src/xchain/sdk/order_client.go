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

type OrderClient struct {
	node   *config.Node
	client pb.SdkOrdererServiceClient
}

func NewOrderClient(node *config.Node) (*OrderClient, error) {
	addr := node.Address()
	if addr == "" {
		msg := "failed to get order addr, addr is empty"
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
			log.Printf("failed to connect to order, addr:%s, err:%v", addr, err)
			return nil, err
		}
	} else {
		var err error
		conn, err = grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Printf("failed to connect to order, addr:%s, err:%v", addr, err)
			return nil, err
		}
	}
	client := pb.NewSdkOrdererServiceClient(conn)
	orderClient := &OrderClient{
		node:   node,
		client: client,
	}
	log.Printf("client started, connected to order:%s With SSL", addr)
	return orderClient, nil
}

func (c *OrderClient) SendTransaction(req *pb.SdkOrdererRequest) (*pb.RpcReply, error) {
	return c.client.SendTransaction(context.Background(), req)
}
