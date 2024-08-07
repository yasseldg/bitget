package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type PlaceOrder struct {
	c common.InterRestClient

	*request.PlaceOrder
}

func NewPlaceOrder(c common.InterRestClient, productType, symbol, side, size, marginMode, marginCoin, orderType string) (*PlaceOrder, error) {

	r, err := request.NewPlaceOrder(productType, symbol, side, size, marginMode, marginCoin, orderType)
	if err != nil {
		return nil, err
	}

	return &PlaceOrder{
		c:          c,
		PlaceOrder: r,
	}, nil
}

func (p *PlaceOrder) Do() (*response.OrderResp, error) {

	uri := "/api/v2/mix/order/place-order"

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	resp, err := p.c.DoPost(uri, params)
	if err != nil {
		return nil, fmt.Errorf("c.DoPost(): %s", err)
	}

	respObj := new(response.OrderResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
