package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"

	"github.com/yasseldg/bitget/pkg/model/v2/errors"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"

	"github.com/yasseldg/go-simple/types/sJson"
)

type DetailOrder struct {
	c common.InterRestClient

	*request.DetailOrder
}

func NewDetailOrder(c common.InterRestClient, productType, symbol string) (*DetailOrder, error) {

	r, err := request.NewDetailOrder(productType, symbol)
	if err != nil {
		return nil, err
	}

	return &DetailOrder{
		c:           c,
		DetailOrder: r,
	}, nil
}

func (p *DetailOrder) Do() (*response.OrderResp, error) {

	if !p.Validate() {
		return nil, errors.ErrMissingOrderIdOrClientOid
	}

	uri := "/api/v2/mix/order/detail"

	resp, err := p.c.DoGet(uri, p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("c.DoGet(): %s", err)
	}

	respObj := new(response.OrderResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
