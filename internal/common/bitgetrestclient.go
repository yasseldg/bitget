package common

import (
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/yasseldg/bitget/config"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/internal"
)

type BitgetRestClient struct {
	ApiKey       string
	ApiSecretKey string
	Passphrase   string
	BaseUrl      string
	HttpClient   *http.Client
	Signer       *Signer
}

func (p *BitgetRestClient) Init() *BitgetRestClient {
	creds := config.GetDefaultCredentials()
	return p.InitWithCreds(creds)
}
func (p *BitgetRestClient) InitWithCreds(creds config.InterApiCreds) *BitgetRestClient {
	p.ApiKey = creds.Key()
	p.ApiSecretKey = creds.Secret()
	p.Passphrase = creds.Pass()

	p.BaseUrl = config.BaseUrl
	p.Signer = new(Signer).Init(creds.Secret())
	p.HttpClient = &http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}

	return p
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := internal.TimesStamp()
	//body, _ := internal.BuildJsonParams(params)

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	requestUrl := config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	internal.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)
	if err != nil {
		return "", err
	}
	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := internal.TimesStamp()
	body := internal.BuildGetParams(params)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := p.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	internal.Headers(request, p.ApiKey, timesStamp, sign, p.Passphrase)

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
