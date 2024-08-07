package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type GetAccountInfo struct {
	c common.InterRestClient
}

func NewGetAccountInfo(c common.InterRestClient) (*GetAccountInfo, error) {
	return &GetAccountInfo{
		c: c,
	}, nil
}

func (p *GetAccountInfo) Do() (*response.AccountInfoResp, error) {

	uri := "/api/v2/spot/account/info"

	resp, err := p.c.DoGet(uri, internal.NewParams())
	if err != nil {
		return nil, err
	}

	respObj := new(response.AccountInfoResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
