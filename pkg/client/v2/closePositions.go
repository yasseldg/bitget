package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"

	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"

	"github.com/yasseldg/go-simple/types/sJson"
)

type ClosePositions struct {
	c common.InterRestClient

	*request.ClosePositions
}

func NewClosePositions(c common.InterRestClient, productType string) (*ClosePositions, error) {

	r, err := request.NewClosePositions(productType)
	if err != nil {
		return nil, err
	}

	return &ClosePositions{
		c:              c,
		ClosePositions: r,
	}, nil
}

func (p *ClosePositions) Do() (*response.PositionResp, error) {

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	uri := "/api/v2/mix/order/close-positions"

	resp, err := p.c.DoPost(uri, params)
	if err != nil {
		return nil, fmt.Errorf("c.DoPost(): %s", err)
	}

	respObj := new(response.PositionResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
