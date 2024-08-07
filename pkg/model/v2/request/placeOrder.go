package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type PlaceOrder struct {
	productType            string
	symbol                 string
	side                   string
	size                   string
	marginMode             string
	marginCoin             string
	orderType              string
	price                  *string
	tradeSide              *string
	force                  *string
	clientOid              *string
	reduceOnly             *string
	presetStopSurplusPrice *string
	presetStopLossPrice    *string
	stpMode                *string
}

func NewPlaceOrder(productType, symbol, side, size, marginMode, marginCoin, orderType string) (*PlaceOrder, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	if len(symbol) == 0 {
		return nil, errors.ErrInvalidSymbol
	}

	if len(side) == 0 {
		return nil, errors.ErrInvalidSide
	}

	if len(size) == 0 {
		return nil, errors.ErrInvalidSize
	}

	if len(marginMode) == 0 {
		return nil, errors.ErrInvalidMarginMode
	}

	if len(marginCoin) == 0 {
		return nil, errors.ErrInvalidMarginCoin
	}

	if len(orderType) == 0 {
		return nil, errors.ErrInvalidOrderType
	}

	return &PlaceOrder{
		productType: productType,
		symbol:      symbol,
		side:        side,
		size:        size,
		marginMode:  marginMode,
		marginCoin:  marginCoin,
		orderType:   orderType,
	}, nil
}

func (o *PlaceOrder) Log() {
	sLog.Info("PlaceOrder: %v ", o.GetParams())
}

func (o *PlaceOrder) Price(price string) *PlaceOrder {
	o.price = &price
	return o
}

func (o *PlaceOrder) TradeSide(tradeSide string) *PlaceOrder {
	o.tradeSide = &tradeSide
	return o
}

func (o *PlaceOrder) Force(force string) *PlaceOrder {
	o.force = &force
	return o
}

func (o *PlaceOrder) ClientOid(clientOid string) *PlaceOrder {
	if len(clientOid) > 0 {
		o.clientOid = &clientOid
	}
	return o
}

func (o *PlaceOrder) ReduceOnly(reduceOnly string) *PlaceOrder {
	o.reduceOnly = &reduceOnly
	return o
}

func (o *PlaceOrder) PresetStopSurplusPrice(presetStopSurplusPrice string) *PlaceOrder {
	o.presetStopSurplusPrice = &presetStopSurplusPrice
	return o
}

func (o *PlaceOrder) PresetStopLossPrice(presetStopLossPrice string) *PlaceOrder {
	o.presetStopLossPrice = &presetStopLossPrice
	return o
}

func (o *PlaceOrder) StpMode(stpMode string) *PlaceOrder {
	o.stpMode = &stpMode
	return o
}

// Params returns the order parameters
func (o *PlaceOrder) GetParams() map[string]string {
	cp := make(map[string]string)

	cp["productType"] = o.productType
	cp["symbol"] = o.symbol
	cp["side"] = o.side
	cp["size"] = o.size
	cp["marginMode"] = o.marginMode
	cp["marginCoin"] = o.marginCoin
	cp["orderType"] = o.orderType

	if o.price != nil {
		cp["price"] = *o.price
	}

	if o.tradeSide != nil {
		cp["tradeSide"] = *o.tradeSide
	}

	if o.force != nil {
		cp["force"] = *o.force
	}

	if o.clientOid != nil {
		cp["clientOid"] = *o.clientOid
	}

	if o.reduceOnly != nil {
		cp["reduceOnly"] = *o.reduceOnly
	}

	if o.presetStopSurplusPrice != nil {
		cp["presetStopSurplusPrice"] = *o.presetStopSurplusPrice
	}

	if o.presetStopLossPrice != nil {
		cp["presetStopLossPrice"] = *o.presetStopLossPrice
	}

	if o.stpMode != nil {
		cp["stpMode"] = *o.stpMode
	}

	return cp
}
