package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type ChangePositionMode struct {
	c common.InterRestClient

	*request.ChangePositionMode
}

func NewChangePositionMode(c common.InterRestClient, productType, posMode string) (*ChangePositionMode, error) {

	r, err := request.NewChangePositionMode(productType, posMode)
	if err != nil {
		return nil, err
	}

	return &ChangePositionMode{
		c:                  c,
		ChangePositionMode: r,
	}, nil
}

func (p *ChangePositionMode) Do() (*response.PositionModeResp, error) {

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	uri := "/api/v2/mix/account/set-position-mode"

	resp, err := p.c.DoPost(uri, params)
	if err != nil {
		return nil, err
	}

	respObj := new(response.PositionModeResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
