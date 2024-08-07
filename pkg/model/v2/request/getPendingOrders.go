package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"
)

type GetPendingOrders struct {
	productType string
	symbol      *string
	orderId     *string
	clientOid   *string
	status      *string
	idLessThan  *string
	startTime   *string
	endTime     *string
	limit       *string
}

func NewGetPendingOrders(productType string) (*GetPendingOrders, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	return &GetPendingOrders{
		productType: productType,
	}, nil
}

func (o *GetPendingOrders) Symbol(symbol string) *GetPendingOrders {
	o.symbol = &symbol
	return o
}

func (o *GetPendingOrders) OrderId(orderId string) *GetPendingOrders {
	o.orderId = &orderId
	return o
}

func (o *GetPendingOrders) ClientOid(clientOid string) *GetPendingOrders {
	o.clientOid = &clientOid
	return o
}

func (o *GetPendingOrders) Status(status string) *GetPendingOrders {
	o.status = &status
	return o
}

func (o *GetPendingOrders) IdLessThan(idLessThan string) *GetPendingOrders {
	o.idLessThan = &idLessThan
	return o
}

func (o *GetPendingOrders) StartTime(startTime string) *GetPendingOrders {
	o.startTime = &startTime
	return o
}

func (o *GetPendingOrders) EndTime(endTime string) *GetPendingOrders {
	o.endTime = &endTime
	return o
}

func (o *GetPendingOrders) Limit(limit string) *GetPendingOrders {
	o.limit = &limit
	return o
}

// Params returns the order parameters
func (o *GetPendingOrders) GetParams() map[string]string {
	cp := make(map[string]string)

	cp["productType"] = o.productType

	if o.symbol != nil {
		cp["symbol"] = *o.symbol
	}

	if o.orderId != nil {
		cp["orderId"] = *o.orderId
	}

	if o.clientOid != nil {
		cp["clientOid"] = *o.clientOid
	}

	if o.status != nil {
		cp["status"] = *o.status
	}

	if o.idLessThan != nil {
		cp["idLessThan"] = *o.idLessThan
	}

	if o.startTime != nil {
		cp["startTime"] = *o.startTime
	}

	if o.endTime != nil {
		cp["endTime"] = *o.endTime
	}

	if o.limit != nil {
		cp["limit"] = *o.limit
	}

	return cp
}
