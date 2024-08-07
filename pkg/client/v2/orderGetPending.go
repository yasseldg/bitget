package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type GetPendingOrders struct {
	c common.InterRestClient

	*request.GetPendingOrders
}

func NewGetPendingOrders(c common.InterRestClient, productType string) (*GetPendingOrders, error) {

	r, err := request.NewGetPendingOrders(productType)
	if err != nil {
		return nil, err
	}

	return &GetPendingOrders{
		c:                c,
		GetPendingOrders: r,
	}, nil
}

func (p *GetPendingOrders) Do() (*response.OrdersResp, error) {

	uri := "/api/v2/mix/order/orders-pending"

	resp, err := p.c.DoGet(uri, p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("c.DoGet(): %s", err)
	}

	respObj := new(response.OrdersResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
