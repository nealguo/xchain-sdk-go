package config

// RPC请求的状态码
const (
	Success = 200
	Error   = 500
)

// 合约类型
const (
	SystemContract = "1"
	UserContract   = "2"
)

// 语言类型
const (
	Golang = "1"
	Java   = "2"
	Python = "3"
	Nodejs = "4"
)

// 共识算法
const (
	Kafka = "kafka"
	Pbft  = "pbft"
	Solo  = "solo"
	Raft  = "raft"
)

// 签名算法
const (
	RSA = "RSA"
	SM2 = "SM2"
)
