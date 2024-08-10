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

type InterRestClient interface {
	DoPost(uri string, params string) (string, error)
	DoGet(uri string, params map[string]string) (string, error)
}

type BitgetRestClient struct {
	HttpClient *http.Client
	Signer     *Signer

	creds config.InterApiCreds
}

func (p *BitgetRestClient) Init(creds config.InterApiCreds) *BitgetRestClient {
	p.creds = config.GetDefaultCredentials(creds)

	p.Signer = new(Signer).Init(p.creds.Secret())
	p.HttpClient = &http.Client{
		Timeout: time.Duration(config.TimeoutSecond) * time.Second,
	}

	return p
}

func (p *BitgetRestClient) DoPost(uri string, params string) (string, error) {
	timesStamp := internal.TimesStamp()

	sign := p.Signer.Sign(constants.POST, uri, params, timesStamp)
	requestUrl := config.BaseUrl + uri

	buffer := strings.NewReader(params)
	request, err := http.NewRequest(constants.POST, requestUrl, buffer)

	internal.Headers(request, p.creds.Key(), timesStamp, sign, p.creds.Pass())
	if err != nil {
		return "", err
	}

	return p.do(request)
}

func (p *BitgetRestClient) DoGet(uri string, params map[string]string) (string, error) {
	timesStamp := internal.TimesStamp()
	body := internal.BuildGetParams(params)

	sign := p.Signer.Sign(constants.GET, uri, body, timesStamp)

	requestUrl := config.BaseUrl + uri + body

	request, err := http.NewRequest(constants.GET, requestUrl, nil)
	if err != nil {
		return "", err
	}
	internal.Headers(request, p.creds.Key(), timesStamp, sign, p.creds.Pass())

	return p.do(request)
}

func (p *BitgetRestClient) do(request *http.Request) (string, error) {

	response, err := p.HttpClient.Do(request)

	if err != nil {
		return "", err
	}

	defer response.Body.Close()

	bodyStr, err := io.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	return string(bodyStr), err
}
