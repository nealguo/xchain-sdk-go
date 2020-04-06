package util

import (
	"log"
	"xchain-sdk-go/src/xchain/config"
)

func Sign(content string) string {
	privateKeyPath := config.AppConf.PrivateKeyPath
	if privateKeyPath == "" {
		log.Printf("sign wrong, privateKeyPath config is emtpy")
		return ""
	}
	signAlgorithm := config.SignConf.Algorithm
	if signAlgorithm == "" {
		log.Printf("sign wrong, signAlgorithm config is emtpy")
		return ""
	}
	if content == "" {
		log.Printf("sign wrong, content is emtpy")
		return ""
	}
	value, err := SignWithKeyPath(content, privateKeyPath, signAlgorithm)
	if err != nil {
		log.Printf("sign wrong, err:%v", err)
		return ""
	}
	return value
}

func Verify(content, signature, certContent string) bool {
	signAlgorithm := config.SignConf.Algorithm
	if signAlgorithm == "" {
		log.Printf("verify wrong, signAlgorithm config is emtpy")
		return false
	}
	if content == "" || signature == "" || certContent == "" {
		log.Printf("verify wrong, one of params is emtpy")
		return false
	}
	value, err := VerifyWithKeyContent(content, signature, certContent, signAlgorithm)
	if err != nil {
		log.Printf("verify wrong, err:%v", err)
		return false
	}
	return value
}
