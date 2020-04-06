package util

import (
	"xchain-sdk-go/src/xchain/config"
	pb "xchain-sdk-go/src/xchain/proto"
)

func BuildContractID(identity, version, language string, system bool) *pb.ContractID {
	var contractType string
	if system {
		contractType = config.SystemContract
	} else {
		contractType = config.UserContract
	}
	return &pb.ContractID{
		Name:         identity,
		Version:      version,
		Type:         contractType,
		LanguageType: language,
	}
}

func BuildContractSpec(contractId *pb.ContractID, filename string, fileContent []byte) *pb.ContractSpec {
	content := string(fileContent)
	sign := Sign(content)
	return &pb.ContractSpec{
		ContractId:    contractId,
		ContractInput: content,
		FileName:      filename,
		ChannelName:   config.ChannelConf.Channel,
		AppId:         config.AppConf.AppId,
		Sign:          sign,
	}
}

func BuildSdkPeerRequest(contractId *pb.ContractID, method, params string) *pb.SdkPeerRequest {
	var sign string
	if contractId.GetType() == config.UserContract {
		sign = Sign(params)
	}
	return &pb.SdkPeerRequest{
		ContractId:  contractId,
		Method:      method,
		Payload:     params,
		ChannelName: config.ChannelConf.Channel,
		AppId:       config.AppConf.AppId,
		Sign:        sign,
	}
}

func BuildSdkOrderRequest(payload string) *pb.SdkOrdererRequest {
	sign := Sign(payload)
	return &pb.SdkOrdererRequest{
		Payload: payload,
		AppId:   config.AppConf.AppId,
		Sign:    sign,
	}
}
