package main

import (
	"fmt"
	"time"
	"xchain-sdk-go/src/test/config"
	"xchain-sdk-go/src/test/service"
	"xchain-sdk-go/src/xchain/util"
)

var svc *service.ContractService

func init() {
	config.ReadYaml()
}

func main() {
	// new service
	svc = service.NewContractService()

	// init contract
	path := "./src/test/contract/contract1/contract1.go"
	initContract(path)
	time.Sleep(120 * time.Second)

	// call method "InitAmount" in contract
	username01 := "user001"
	amount01 := 100
	initAmount(username01, amount01)
	username02 := "user002"
	amount02 := 100
	initAmount(username02, amount02)
	time.Sleep(3 * time.Second)

	// call method "QueryAmount" in contract
	queryAmount(username01)
	queryAmount(username02)

	// call method "TransferAmount" in contract
	amount03 := 20
	transferAmount(username01, username02, amount03)
	time.Sleep(3 * time.Second)

	// call method "QueryAmount" in contract
	queryAmount(username01)
	queryAmount(username02)
}

func initContract(path string) {
	resp := svc.Init(path)
	fmt.Println(resp)
}

func initAmount(username string, amount int) {
	kv := make(map[string]interface{}, 2)
	kv["userId"] = username
	kv["amount"] = amount
	method := "InitAmount"
	params := util.ToJson(kv)
	resp := svc.Invoke(method, params)
	fmt.Println(resp)
}

func queryAmount(username string) {
	kv := make(map[string]interface{}, 1)
	kv["userId"] = username
	method := "QueryAmount"
	params := util.ToJson(kv)
	resp := svc.Invoke(method, params)
	fmt.Println(resp)
}

func transferAmount(fromUserName, toUsername string, amount int) {
	kv := make(map[string]interface{}, 3)
	kv["fromUserId"] = fromUserName
	kv["toUserId"] = toUsername
	kv["transferAmount"] = amount
	method := "TransferAmount"
	params := util.ToJson(kv)
	resp := svc.Invoke(method, params)
	fmt.Println(resp)
}
