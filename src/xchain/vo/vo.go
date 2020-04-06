package vo

type ContractVo struct {
	Method string
	Data   string
	appId  string
}

type Response struct {
	Code    int
	Message string
	Payload string
}
