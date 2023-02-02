package common

import (
	"bitget/config"
	"bitget/constants"
	"bitget/internal"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
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
func (p *BitgetRestClient) InitWithCreds(creds *config.ApiCreds) *BitgetRestClient {
	p.ApiKey = creds.ApiKey
	p.ApiSecretKey = creds.SecretKey
	p.Passphrase = creds.PASSPHRASE

	p.BaseUrl = config.BaseUrl
	p.Signer = new(Signer).Init(creds.SecretKey)
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

	bodyStr, err := ioutil.ReadAll(response.Body)
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

	bodyStr, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	responseBodyString := string(bodyStr)
	return responseBodyString, err
}
