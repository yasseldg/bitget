package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type GetAccountList struct {
	c common.InterRestClient

	*request.GetAccountList
}

func NewGetAccountList(c common.InterRestClient, productType string) (*GetAccountList, error) {

	r, err := request.NewGetAccountList(productType)
	if err != nil {
		return nil, err
	}

	return &GetAccountList{
		c:              c,
		GetAccountList: r,
	}, nil
}

func (p *GetAccountList) Do() (*response.AccountsResp, error) {

	uri := "/api/v2/mix/account/accounts"

	resp, err := p.c.DoGet(uri, p.GetParams())
	if err != nil {
		return nil, err
	}

	respObj := new(response.AccountsResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
