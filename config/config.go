package config

import "os"

const (
	//restApi config
	BaseUrl       = "https://api.bitget.com"
	TimeoutSecond = 30

	//websocket config
	WsUrl = "wss://ws.bitgetapi.com/spot/v1/stream"
)


// credentials
type ApiCreds struct {
	ApiKey string
	SecretKey string
	PASSPHRASE string
}

func GetDefaultCredentials() *ApiCreds {
	return &ApiCreds{
		ApiKey:     os.Getenv("BITGET_API_KEY"),
		SecretKey:  os.Getenv("BITGET_API_SECRET"),
		PASSPHRASE: os.Getenv("BITGET_API_PASS"),
	}
}
