package config

type ContractConfig struct {
	Identity string `yaml:"identity"`
	Version  string `yaml:"version"`
}

type BaaSConfig struct {
	Sdk       SdkConfig `yaml:"sdk"`
	Algorithm string    `yaml:"algorithm"`
	Consensus string    `yaml:"consensus"`
	Channel   string    `yaml:"channel"`
	Order     string    `yaml:"order"`
	Peer      string    `yaml:"peer"`
}

type SdkConfig struct {
	AppId             string `yaml:"appId"`
	PrivateKeyPath    string `yaml:"privateKeyPath"`
	SslCertFilePath   string `yaml:"sslCertFilePath"`
	SslPrivateKeyPath string `yaml:"sslPrivateKeyPath"`
}

type Yaml struct {
	ContractConf map[string]ContractConfig `yaml:"contract"`
	BaaSConf     BaaSConfig                `yaml:"baas"`
}
