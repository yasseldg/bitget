package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type ChangeMarginMode struct {
	c common.InterRestClient

	*request.ChangeMarginMode
}

func NewChangeMarginMode(c common.InterRestClient, productType, symbol, marginCoin, marginMode string) (*ChangeMarginMode, error) {

	r, err := request.NewChangeMarginMode(productType, symbol, marginCoin, marginMode)
	if err != nil {
		return nil, err
	}

	return &ChangeMarginMode{
		c:                c,
		ChangeMarginMode: r,
	}, nil
}

func (p *ChangeMarginMode) Do() (*response.MarginModeResp, error) {

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	uri := "/api/v2/mix/account/set-margin-mode"

	resp, err := p.c.DoPost(uri, params)
	if err != nil {
		return nil, err
	}

	respObj := new(response.MarginModeResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
