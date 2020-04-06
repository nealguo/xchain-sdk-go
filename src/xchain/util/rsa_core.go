package util

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func ReadRsaPrivateKey(privateKeyPath string) string {
	if privateKeyPath == "" {
		log.Printf("ReadRsaPrivateKey wrong, private key dose not exist")
		return ""
	}
	bytes, err := ioutil.ReadFile(privateKeyPath)
	if err != nil {
		log.Printf("ReadRsaPrivateKey wrong, read private key file wrong, err:%v", err)
		return ""
	}
	content := string(bytes)
	if content == "" {
		log.Printf("ReadRsaPrivateKey wrong, private key file is empty")
		return ""
	}
	head := `-----BEGIN PRIVATE KEY-----`
	foot := `-----END PRIVATE KEY-----`
	content = strings.Replace(content, head, "", -1)
	content = strings.Replace(content, foot, "", -1)
	return strings.TrimSpace(content)
}

func SignWithRsa(content string, key string) (string, error) {
	keyBytes, err := base64.StdEncoding.DecodeString(key)
	if err != nil {
		fmt.Printf("SignWithRsa wrong, decode private key cotent wrong, err:%v", err)
		return "", err
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(keyBytes)
	if err != nil {
		fmt.Printf("SignWithRsa wrong, parse private key wrong, err:%v", err)
		return "", err
	}
	h := sha256.New()
	_, err = h.Write([]byte(content))
	if err != nil {
		fmt.Printf("SignWithRsa wrong, write content to sha256 wrong, err:%v", err)
		return "", err
	}
	hash := h.Sum(nil)
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey.(*rsa.PrivateKey), crypto.SHA256, hash[:])
	if err != nil {
		fmt.Printf("SignWithRsa wrong, sign wrong, err:%v", err)
		return "", err
	}
	out := base64.StdEncoding.EncodeToString(signature)
	return out, nil
}

func VerifyWithRsa(content, signature, certContent string) (bool, error) {
	sign, err := base64.StdEncoding.DecodeString(signature)
	if err != nil {
		fmt.Printf("VerifyWithRsa wrong, decode signature wrong, err:%v", err)
		return false, err
	}

	block, _ := pem.Decode([]byte(certContent))
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		fmt.Printf("VerifyWithRsa wrong, decode public key cotent wrong, err:%v", err)
		return false, err
	}
	publicKey := cert.PublicKey.(*rsa.PublicKey)

	hash := sha256.New()
	_, err = hash.Write([]byte(content))
	if err != nil {
		fmt.Printf("VerifyWithRsa wrong, write content to sha256 wrong, err:%v", err)
		return false, err
	}
	err = rsa.VerifyPKCS1v15(publicKey, crypto.SHA256, hash.Sum(nil), sign)
	if err != nil {
		fmt.Printf("VerifyWithRsa wrong, verify wrong, err:%v", err)
		return false, err
	}
	return true, nil
}

func ReadRsaCert(certPath string) string {
	if certPath == "" {
		log.Printf("ReadRsaCert wrong, private key dose not exist")
		return ""
	}
	bytes, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Printf("ReadRsaCert wrong, read private key file wrong, err:%v", err)
		return ""
	}
	content := string(bytes)
	if content == "" {
		log.Printf("ReadRsaCert wrong, private key file is empty")
		return ""
	}
	return content
}
