package config

import "os"

const (
	//restApi config
	BaseUrl       = "https://api.bitget.com"
	TimeoutSecond = 30

	//websocket config
	WsUrl = "wss://ws.bitgetapi.com/spot/v1/stream"

	WsUrlPublic  = "wss://ws.bitget.com/v2/ws/public"
	WsUrlPrivate = "wss://ws.bitget.com/v2/ws/private"
)

func GetDefaultCredentials(creds InterApiCreds) *ApiCreds {
	_creds := &ApiCreds{
		ApiKey:     os.Getenv("BITGET_API_KEY"),
		SecretKey:  os.Getenv("BITGET_API_SECRET"),
		PASSPHRASE: os.Getenv("BITGET_API_PASS"),
	}

	if creds == nil {
		return _creds
	}

	if len(creds.Key()) > 0 {
		_creds.ApiKey = creds.Key()
	}

	if len(creds.Secret()) > 0 {
		_creds.SecretKey = creds.Secret()
	}

	if len(creds.Pass()) > 0 {
		_creds.PASSPHRASE = creds.Pass()
	}

	return _creds
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
