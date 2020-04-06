package util

import (
	"encoding/base64"
	"io/ioutil"
	"log"
	"strings"
	"xchain-sdk-go/src/xchain/crypto"
)

var user = []byte("tanghuanyou@163.com")

func ReadSm2PrivateKey(privateKeyPath string) string {
	if privateKeyPath == "" {
		log.Printf("ReadSm2PrivateKey wrong, private key dose not exist")
		return ""
	}
	bytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Printf("ReadSm2PrivateKey wrong, read private key file wrong, err:%v", err)
		return ""
	}
	content := string(bytes)
	if content == "" {
		log.Printf("ReadSm2PrivateKey wrong, private key file is empty")
		return ""
	}
	head := `-----BEGIN PRIVATE KEY-----`
	foot := `-----END PRIVATE KEY-----`
	content = strings.Replace(content, head, "", -1)
	content = strings.Replace(content, foot, "", -1)
	return strings.TrimSpace(content)
}

func SignWithSm2(content string, key string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		log.Printf("SignWithSm2 wrong, decode private key wrong, err:%v", err)
		return "", err
	}
	privateKey, err := crypto.RawBytesToPrivateKey(keyBytes)
	if err != nil {
		log.Printf("SignWithSm2 wrong, read private key wrong, err:%v", err)
		return "", err
	}
	sign, err := crypto.Sign(privateKey, user, []byte(content))
	if err != nil {
		log.Printf("SignWithSm2 wrong, sign wrong, err:%v", err)
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(sign)
	return out, nil
}

func VerifyWithSm2(content, signature, certContent string) (bool, error) {
	//TODO read public key content from certificate file content

	// public key content is encoded with base64
	publicKeyContent := `BE+wuL4Tih369HVpCbbk1dg5LNLPNnDWF0BFjqnhZgGq8Ib7jrWbbIMH3lZOrJ20ALEa0Cy/XvTvMg54AlIoQZU=`
	publicKeyBytes, err := base64.StdEncoding.DecodeString(publicKeyContent)
	if err != nil {
		log.Printf("VerifyWithSm2 wrong, decode public key wrong, err:%v", err)
		return false, err
	}
	publicKey, err := crypto.RawBytesToPublicKey(publicKeyBytes)
	if err != nil {
		log.Printf("VerifyWithSm2 wrong, read public key wrong, err:%v", err)
		return false, err
	}
	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		log.Printf("VerifyWithSm2 wrong, decode signature wrong, err:%v", err)
		return false, err
	}

	result := crypto.Verify(publicKey, user, []byte(content), sign)
	return result, nil
}
