package service

import (
	"xchain-sdk-go/src/test/config"
	"xchain-sdk-go/src/xchain/sdk"
	"xchain-sdk-go/src/xchain/vo"
)

type ContractService struct {
	sdk *sdk.Client
}

func NewContractService() *ContractService {
	// chain config
	baas := config.Conf.BaaSConf
	channel := baas.Channel
	consensus := baas.Consensus
	order := baas.Order
	peer := baas.Peer
	key := baas.Sdk.PrivateKeyPath
	appId := baas.Sdk.AppId
	algorithm := baas.Algorithm
	// ssl config
	ssl := false
	sslCert := baas.Sdk.SslCertFilePath
	sslKey := baas.Sdk.SslPrivateKeyPath
	sslTrust := ""
	if sslCert != "" && sslKey != "" {
		ssl = true
	}
	// contract service
	c := sdk.NewChannel(channel, consensus, order, peer, key, appId, algorithm, ssl, sslCert, sslKey, sslTrust)
	s := sdk.NewClient(c)
	return &ContractService{sdk: s}
}

func (s *ContractService) InitContract(identity, version, path string) *vo.Response {
	return s.sdk.InitContractWithResult(identity, version, path)
}

func (s *ContractService) InvokeContract(identity, version, method, params string) *vo.Response {
	return s.sdk.Invoke(identity, version, method, params)
}

func (s *ContractService) Invoke(method, params string) *vo.Response {
	// use the first config of "contract" in yaml
	demo := config.Conf.ContractConf["demo"]
	identity := demo.Identity
	version := demo.Version
	return s.sdk.Invoke(identity, version, method, params)
}

func (s *ContractService) Init(path string) *vo.Response {
	// use the first config of "contract" in yaml
	demo := config.Conf.ContractConf["demo"]
	identity := demo.Identity
	version := demo.Version
	return s.sdk.InitContractWithResult(identity, version, path)
}
