package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type DetailOrder struct {
	productType string
	symbol      string

	orderId   *string
	clientOid *string
}

func NewDetailOrder(productType, symbol string) (*DetailOrder, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	if len(symbol) == 0 {
		return nil, errors.ErrInvalidSymbol
	}

	return &DetailOrder{
		productType: productType,
		symbol:      symbol,
	}, nil
}

func (o *DetailOrder) Log() {
	sLog.Info("DetailOrder: %v ", o.GetParams())
}

func (o *DetailOrder) Validate() bool {
	return o.orderId != nil || o.clientOid != nil
}

func (o *DetailOrder) OrderId(orderId string) *DetailOrder {
	o.orderId = &orderId
	return o
}

func (o *DetailOrder) ClientOid(clientOid string) *DetailOrder {
	o.clientOid = &clientOid
	return o
}

// Params returns the order parameters
func (o *DetailOrder) GetParams() map[string]string {
	cp := make(map[string]string)

	cp["productType"] = o.productType
	cp["symbol"] = o.symbol

	if o.orderId != nil {
		cp["orderId"] = *o.orderId
	}

	if o.clientOid != nil {
		cp["clientOid"] = *o.clientOid
	}

	return cp
}
