package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"

	"github.com/yasseldg/bitget/pkg/model/v2/errors"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"

	"github.com/yasseldg/go-simple/types/sJson"
)

type ModifyOrder struct {
	c common.InterRestClient

	*request.ModifyOrder
}

func NewModifyOrder(c common.InterRestClient, productType, symbol, newClientOid string) (*ModifyOrder, error) {

	r, err := request.NewModifyOrder(productType, symbol, newClientOid)
	if err != nil {
		return nil, err
	}

	return &ModifyOrder{
		c:           c,
		ModifyOrder: r,
	}, nil
}

func (p *ModifyOrder) Do() (*response.OrderResp, error) {

	if !p.Validate() {
		return nil, errors.ErrMissingOrderIdOrClientOid
	}

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	uri := "/api/v2/mix/order/modify-order"

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
