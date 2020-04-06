package contract1

import (
	"errors"
	"log"
	"xchain-contract-go/src/xchain/contract"
	"xchain-contract-go/src/xchain/reply"
	"xchain-contract-go/src/xchain/util"
	"xchain-contract-go/src/xchain/vo"
)

// 创建账户
func InitAmount(stub *contract.Stub, vo *vo.ContractVo) (*vo.Response, error) {
	// 检查参数
	payload := vo.Data
	if payload == "" {
		log.Printf("empty payload when calling " + vo.Method)
		return nil, errors.New("empty payload when calling " + vo.Method)
	}
	log.Printf("payload:%s\n", payload)
	params := util.JsonToMap(payload)
	if params == nil {
		log.Printf("参数类型错误, payload:%s", payload)
		return reply.ContractResponseError("参数类型错误"), nil
	}

	// 检查账户
	key := util.ToString(params["userId"])
	if key == "" {
		log.Printf("该账户为空或类型错误, payload:%s", payload)
		return reply.ContractResponseError("该账户为空或类型错误"), nil
	}
	state := stub.GetState(key)
	if state != "" {
		log.Printf("账户已存在, userId:%s", key)
		return reply.ContractResponseError("账户已存在"), nil
	}

	// 创建帐户
	amount := util.ToFloat64(params["amount"])
	if amount < 0 {
		log.Printf("该账户金额为负或类型错误, payload:%s", payload)
		return reply.ContractResponseError("该账户金额为负或类型错误"), nil
	}
	value := util.Float64ToStr(amount)
	stub.PutState(key, value)

	// 返回结果
	log.Printf("账户初始化成功, userId:%s, amount:%s", key, value)
	return reply.ContractResponseSuccess("账户初始化成功"), nil
}

// 查询多个账户
func QueryAmount(stub *contract.Stub, vo *vo.ContractVo) (*vo.Response, error) {
	// 检查参数
	payload := vo.Data
	if payload == "" {
		log.Printf("empty payload when calling " + vo.Method)
		return nil, errors.New("empty payload when calling " + vo.Method)
	}
	log.Printf("payload:%s\n", payload)
	params := util.JsonToMap(payload)
	if params == nil {
		log.Printf("参数类型错误, payload:%s", payload)
		return reply.ContractResponseError("参数类型错误"), nil
	}

	// 查询账户
	key := util.ToString(params["userId"])
	if key == "" {
		log.Printf("该账户为空或类型错误, payload:%s", payload)
		return reply.ContractResponseError("该账户为空或类型错误"), nil
	}
	value := stub.GetState(key)

	// 返回结果
	log.Printf("查询账户成功, userId:%s, value:%s", key, value)
	return reply.ContractResponseSuccess(value), nil
}

// 转账
func TransferAmount(stub *contract.Stub, vo *vo.ContractVo) (*vo.Response, error) {
	// 检查参数
	payload := vo.Data
	if payload == "" {
		log.Printf("empty payload when calling " + vo.Method)
		return nil, errors.New("empty payload when calling " + vo.Method)
	}
	log.Printf("payload:%s\n", payload)
	params := util.JsonToMap(payload)
	if params == nil {
		log.Printf("参数类型错误, payload:%s", payload)
		return reply.ContractResponseError("参数类型错误"), nil
	}
	transferAmount := params["transferAmount"]
	if transferAmount == nil {
		log.Printf("转账资金为空, payload:%s", payload)
		return reply.ContractResponseError("转账资金为空"), nil
	}
	tAmt := util.ToFloat64(params["transferAmount"])
	if tAmt <= 0 {
		log.Printf("转账资金必须大于0, payload:%s", payload)
		return reply.ContractResponseError("转账资金必须大于0"), nil
	}

	// 检查转账账户和收款账户
	fromUserId := util.ToString(params["fromUserId"])
	if fromUserId == "" {
		log.Printf("转账账户为空或类型错误, payload:%s", payload)
		return reply.ContractResponseError("转账账户为空或类型错误"), nil
	}
	toUserId := util.ToString(params["toUserId"])
	if toUserId == "" {
		log.Printf("收款账户为空或类型错误, payload:%s", payload)
		return reply.ContractResponseError("收款账户为空或类型错误"), nil
	}
	if fromUserId == toUserId {
		log.Printf("转账账户和收款账户不能相同, payload:%s", payload)
		return reply.ContractResponseError("转账账户和收款账户不能相同"), nil
	}

	// 检查账户的资金
	fromUserAmount := stub.GetState(fromUserId)
	if fromUserAmount == "" {
		log.Printf("转账用户的资金账户不存在, payload:%s", payload)
		return reply.ContractResponseError("转账用户的资金账户不存在"), nil
	}
	toUserAmount := stub.GetState(toUserId)
	if toUserAmount == "" {
		log.Printf("收款用户的资金账户不存在, payload:%s", payload)
		return reply.ContractResponseError("收款用户的资金账户不存在"), nil
	}
	fromUserIdAmount := util.StrToFloat64(fromUserAmount)
	if fromUserIdAmount <= 0 {
		log.Printf("转账用户的资金账户余额不足, payload:%s", payload)
		return reply.ContractResponseError("转账用户的资金账户余额不足"), nil
	}
	if fromUserIdAmount < tAmt {
		log.Printf("转账用户的当前余额小于转账金额, payload:%s", payload)
		return reply.ContractResponseError("转账用户的当前余额小于转账金额"), nil
	}
	toUserIdAmount := util.StrToFloat64(toUserAmount)
	if toUserIdAmount < 0 {
		log.Printf("收款用户的资金账户余额为负, payload:%s", payload)
		return reply.ContractResponseError("收款用户的资金账户余额为负"), nil
	}

	// 转账
	fromUserIdValue := util.Float64ToStr(fromUserIdAmount - tAmt)
	stub.PutState(fromUserId, fromUserIdValue)
	toUserIdValue := util.Float64ToStr(toUserIdAmount + tAmt)
	stub.PutState(toUserId, toUserIdValue)

	// 返回结果
	log.Printf("转账成功, %s:%s, %s:%s", fromUserId, fromUserIdValue, toUserId, toUserIdValue)
	return reply.ContractResponseSuccess("转账成功"), nil
}
