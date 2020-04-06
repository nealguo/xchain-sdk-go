package config

var ChannelConf *ChannelConfig
var AppConf *AppConfig
var SslConf *SslConfig
var SignConf *SignConfig

type ChannelConfig struct {
	Channel   string
	Consensus string
	Order     string
	Peer      string
}

type AppConfig struct {
	PrivateKeyPath string
	AppId          string
}

type SslConfig struct {
	SslEnable            bool
	SslMutual            bool
	SslCertFilePath      string
	SslPrivateKeyPath    string
	SslTrustCertFilePath string
}

type SignConfig struct {
	Algorithm string
}

func InitConfigs(channel, consensus, order, peer, key, appId, algorithm string, ssl bool, sslCert, sslKey, sslTrust string) {
	ChannelConf = &ChannelConfig{
		Channel:   channel,
		Consensus: consensus,
		Order:     order,
		Peer:      peer,
	}
	AppConf = &AppConfig{
		PrivateKeyPath: key,
		AppId:          appId,
	}
	SignConf = &SignConfig{
		Algorithm: algorithm,
	}
	SslConf = &SslConfig{
		SslEnable:            ssl,
		SslMutual:            false,
		SslCertFilePath:      sslCert,
		SslPrivateKeyPath:    sslKey,
		SslTrustCertFilePath: sslTrust,
	}
}
