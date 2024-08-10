package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type ModifyOrder struct {
	productType  string
	symbol       string
	newClientOid string

	orderId                   *string
	clientOid                 *string
	newSize                   *string
	newPrice                  *string
	newPresetStopSurplusPrice *string
	newPresetStopLossPrice    *string
}

func NewModifyOrder(productType, symbol, newClientOid string) (*ModifyOrder, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	if len(symbol) == 0 {
		return nil, errors.ErrInvalidSymbol
	}

	if len(newClientOid) == 0 {
		return nil, errors.ErrInvalidNewClientOid
	}

	return &ModifyOrder{
		productType:  productType,
		symbol:       symbol,
		newClientOid: newClientOid,
	}, nil
}

func (o *ModifyOrder) Log() {
	sLog.Info("ModifyOrder: %v ", o.GetParams())
}

func (o *ModifyOrder) Validate() bool {
	return o.orderId != nil || o.clientOid != nil
}

func (o *ModifyOrder) OrderId(orderId string) *ModifyOrder {
	o.orderId = &orderId
	return o
}

func (o *ModifyOrder) ClientOid(clientOid string) *ModifyOrder {
	o.clientOid = &clientOid
	return o
}

func (o *ModifyOrder) NewSize(newSize string) *ModifyOrder {
	o.newSize = &newSize
	return o
}

func (o *ModifyOrder) NewPrice(newPrice string) *ModifyOrder {
	o.newPrice = &newPrice
	return o
}

func (o *ModifyOrder) NewPresetStopSurplusPrice(newPresetStopSurplusPrice string) *ModifyOrder {
	o.newPresetStopSurplusPrice = &newPresetStopSurplusPrice
	return o
}

func (o *ModifyOrder) NewPresetStopLossPrice(newPresetStopLossPrice string) *ModifyOrder {
	o.newPresetStopLossPrice = &newPresetStopLossPrice
	return o
}

// Params returns the order parameters
func (o *ModifyOrder) GetParams() map[string]string {
	cp := make(map[string]string)

	cp["productType"] = o.productType
	cp["symbol"] = o.symbol
	cp["newClientOid"] = o.newClientOid

	if o.orderId != nil {
		cp["orderId"] = *o.orderId
	}

	if o.clientOid != nil {
		cp["clientOid"] = *o.clientOid
	}

	if o.newSize != nil {
		cp["newSize"] = *o.newSize
	}

	if o.newPrice != nil {
		cp["newPrice"] = *o.newPrice
	}

	if o.newPresetStopSurplusPrice != nil {
		cp["newPresetStopSurplusPrice"] = *o.newPresetStopSurplusPrice
	}

	if o.newPresetStopLossPrice != nil {
		cp["newPresetStopLossPrice"] = *o.newPresetStopLossPrice
	}

	return cp
}

// Parameter	Type	Required	Description
// orderId	String	No	Order ID
// Either orderId or clientOid is required. If both are entered, orderId prevails.
// clientOid	String	No	Customize order ID
// Either orderId or clientOid is required. If both are entered, orderId prevails.
// symbol	String	Yes	Trading pair, e.g. ETHUSDT
// productType	String	Yes	Product type
// USDT-FUTURES USDT professional futures
// COIN-FUTURES Mixed futures
// USDC-FUTURES USDC professional futures
// SUSDT-FUTURES USDT professional futures demo
// SCOIN-FUTURES Mixed futures demo
// SUSDC-FUTURES USDC professional futures demo
// newClientOid	String	Yes	New customized order ID after order modification
// newSize	String	No	Amount of the modified transaction
// The amount stays unchanged if the field if left blank.
// newPrice	String	No	Modified price for placing new orders.
// 1. When the existing order type is Limit, the original price will be maintained if the field is left empty.
// 2. When the existing order type is Limit market, the field should not be set.
// newPresetStopSurplusPrice	String	No	Modifying take-profit
// 1. If the original order has take-profit set and the field is empty, the original value will be kept.
// 2. If the original order has take-profit set and the field is filled in with a value, TP will be updated; if the original order has take-profit set and the field is not set, a new take-profit value will be added.
// If there was a TP value and a 0 is filled in the filled, the existing TP will be deleted.
// newPresetStopLossPrice	String	No	Modifying stop-loss
// 1. If the original order has stop-loss set and the field is empty, the original value will be kept.
// 2. If the original order has stop-loss set and the field is filled in with a value, TP will be updated; if the original order has stop-loss set and the field is not set, a new stop-loss value will be added.
// If there was a SL value and a 0 is filled in the filled, the existing SL will be deleted.
