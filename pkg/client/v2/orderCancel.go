package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"

	"github.com/yasseldg/bitget/pkg/model/v2/errors"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"

	"github.com/yasseldg/go-simple/types/sJson"
)

type CancelOrder struct {
	c common.InterRestClient

	*request.CancelOrder
}

func NewCancelOrder(c common.InterRestClient, productType, symbol string) (*CancelOrder, error) {

	r, err := request.NewCancelOrder(productType, symbol)
	if err != nil {
		return nil, err
	}

	return &CancelOrder{
		c:           c,
		CancelOrder: r,
	}, nil
}

func (p *CancelOrder) Do() (*response.OrderResp, error) {

	if !p.Validate() {
		return nil, errors.ErrMissingOrderIdOrClientOid
	}

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	uri := "/api/v2/mix/order/cancel-order"

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
