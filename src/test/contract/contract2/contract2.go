package contract2

import (
	"errors"
	"log"
	"xchain-contract-go/src/xchain/contract"
	"xchain-contract-go/src/xchain/reply"
	"xchain-contract-go/src/xchain/vo"
)

// 创建账户
func InitAmount(stub *contract.Stub, vo *vo.ContractVo) (*vo.Response, error) {
	payload := vo.Data
	if payload == "" {
		log.Printf("empty payload when calling " + vo.Method)
		return nil, errors.New("empty payload when calling " + vo.Method)
	}
	log.Printf("payload:%s\n", payload)

	// 调用其他合约
	resp, err := stub.InvokeOtherContract("contract1", "1.0", "InitAmount", payload)
	if err != nil {
		log.Printf("参数类型错误, payload:%s, err:%v\n", payload, err)
		return reply.ContractResponseError("参数类型错误"), err
	}

	// 返回结果
	log.Printf("账户初始化成功, payload:%s", payload)
	return reply.ContractResponse(resp.Code, resp.Message, resp.Payload), nil
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

	// 调用其他合约
	resp, err := stub.InvokeOtherContract("contract1", "1.0", "QueryAmount", payload)
	if err != nil {
		log.Printf("参数类型错误, payload:%s, err:%v\n", payload, err)
		return reply.ContractResponseError("参数类型错误"), err
	}

	// 返回结果
	log.Printf("查询账户成功, payload:%s", payload)
	return reply.ContractResponse(resp.Code, resp.Message, resp.Payload), nil
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

	// 调用其他合约
	resp, err := stub.InvokeOtherContract("contract1", "1.0", "TransferAmount", payload)
	if err != nil {
		log.Printf("参数类型错误, payload:%s, err:%v\n", payload, err)
		return reply.ContractResponseError("参数类型错误"), err
	}

	// 返回结果
	log.Printf("转账成功, payload:%s", payload)
	return reply.ContractResponse(resp.Code, resp.Message, resp.Payload), nil
}
