package model

// 版本模型
type Version struct {
	BlockNum int `json:"blockNum"`
	TxNum    int `json:"txNum"`
}

// 写集合模型
type RwSetWrite struct {
	// 字段名按字典序排列，否则在Order中写集合验签会失败
	AppId           string `json:"appId"`
	Collection      string `json:"collection"`
	ContractVersion string `json:"contractVersion"`
	Delete          bool   `json:"delete"`
	Key             string `json:"key"`
	Value           string `json:"value"`
}

// 读集合模型
type RwSetRead struct {
	Key             string  `json:"key"`
	Version         Version `json:"version"`
	Collection      string  `json:"collection"`
	ContractVersion string  `json:"contractVersion"`
}

// 合约调用轨迹
type ContractInvokeTrace struct {
	Invoke           string                 `json:"invoke"`
	ContractIdentity string                 `json:"contractIdentity"`
	ContractVersion  string                 `json:"contractVersion"`
	SubTraceList     []*ContractInvokeTrace `json:"subTraceList"`
}

// 交易模型
type Transaction struct {
	TxId                int64                `json:"txId"`
	Invoke              string               `json:"invoke"`
	ContractIdentity    string               `json:"contractIdentity"`
	ContractVersion     string               `json:"contractVersion"`
	Writes              []*RwSetWrite        `json:"writes"`
	Reads               []*RwSetRead         `json:"reads"`
	Version             Version              `json:"version"`
	ChannelName         string               `json:"channelName"`
	Offset              int64                `json:"offset"`
	Hash                string               `json:"hash"`
	Timestamp           string               `json:"timestamp"`
	EndorsePeerName     string               `json:"endorsePeerName"`
	EndorseSign         string               `json:"endorseSign"`
	ContractInvokeTrace *ContractInvokeTrace `json:"contractInvokeTrace"`
}

// 合约调用后返回值对应的模型
type ResponseModel struct {
	Value       string       `json:"value"`
	Transaction *Transaction `json:"transaction"`
}

// 存储在状态库中的状态的模型
type StateModel struct {
	Value   string  `json:"value"`
	Version Version `json:"version"`
}
