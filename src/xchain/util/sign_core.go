package util

import (
	"errors"
	"fmt"
	"log"
	"sync"
	"xchain-sdk-go/src/xchain/config"
)

var sm2Pks = make(map[string]string, 0)
var rsaPks = make(map[string]string, 0)

var sm2PksMutex sync.Mutex
var rsaPksMutex sync.Mutex

func SignWithKeyPath(content, privateKeyPath, signAlgorithm string) (string, error) {
	if config.RSA == signAlgorithm {
		rsaPksMutex.Lock()
		pk, ok := rsaPks[privateKeyPath]
		if !ok || pk == "" {
			pk = ReadRsaPrivateKey(privateKeyPath)
			rsaPks[privateKeyPath] = pk
		}
		rsaPksMutex.Unlock()
		return SignWithRsa(content, pk)
	} else if config.SM2 == signAlgorithm {
		sm2PksMutex.Lock()
		pk, ok := sm2Pks[privateKeyPath]
		if !ok || pk == "" {
			pk = ReadSm2PrivateKey(privateKeyPath)
			sm2Pks[privateKeyPath] = pk
		}
		sm2PksMutex.Unlock()
		return SignWithSm2(content, pk)
	} else {
		msg := fmt.Sprintf("unsupported sign algorithm:%s", signAlgorithm)
		log.Printf(msg)
		return "", errors.New(msg)
	}
}

func VerifyWithKeyContent(content, signature, certContent, signAlgorithm string) (bool, error) {
	if config.RSA == signAlgorithm {
		return VerifyWithRsa(content, signature, certContent)
	} else if config.SM2 == signAlgorithm {
		return VerifyWithSm2(content, signature, certContent)
	} else {
		msg := fmt.Sprintf("unsupported sign algorithm:%s", signAlgorithm)
		log.Printf(msg)
		return false, errors.New(msg)
	}
}
