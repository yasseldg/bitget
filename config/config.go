package config

import "os"

const (
	//restApi config
	BaseUrl       = "https://api.bitget.com"
	TimeoutSecond = 30

	//websocket config
	WsUrl = "wss://ws.bitgetapi.com/spot/v1/stream"
)

func GetDefaultCredentials() *ApiCreds {
	return &ApiCreds{
		ApiKey:     os.Getenv("BITGET_API_KEY"),
		SecretKey:  os.Getenv("BITGET_API_SECRET"),
		PASSPHRASE: os.Getenv("BITGET_API_PASS"),
	}
}

type InterApiCreds interface {
	Key() string
	Secret() string
	Pass() string
}

// credentials
type ApiCreds struct {
	ApiKey     string
	SecretKey  string
	PASSPHRASE string
}

func (p *ApiCreds) Key() string {
	return p.ApiKey
}

func (p *ApiCreds) Secret() string {
	return p.SecretKey
}

func (p *ApiCreds) Pass() string {
	return p.PASSPHRASE
}
