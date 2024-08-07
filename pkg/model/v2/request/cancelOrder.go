package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type CancelOrder struct {
	productType string
	symbol      string

	orderId    *string
	clientOid  *string
	marginCoin *string
}

func NewCancelOrder(productType, symbol string) (*CancelOrder, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	if len(symbol) == 0 {
		return nil, errors.ErrInvalidSymbol
	}

	return &CancelOrder{
		productType: productType,
		symbol:      symbol,
	}, nil
}

func (o *CancelOrder) Log() {
	sLog.Info("CancelOrder: %v ", o.GetParams())
}

func (o *CancelOrder) Validate() bool {
	return o.orderId != nil || o.clientOid != nil
}

func (o *CancelOrder) OrderId(orderId string) *CancelOrder {
	o.orderId = &orderId
	return o
}

func (o *CancelOrder) ClientOid(clientOid string) *CancelOrder {
	o.clientOid = &clientOid
	return o
}

func (o *CancelOrder) MarginCoin(marginCoin string) *CancelOrder {
	o.marginCoin = &marginCoin
	return o
}

// Params returns the order parameters
func (o *CancelOrder) GetParams() map[string]string {
	cp := make(map[string]string)

	cp["productType"] = o.productType
	cp["symbol"] = o.symbol

	if o.orderId != nil {
		cp["orderId"] = *o.orderId
	}

	if o.clientOid != nil {
		cp["clientOid"] = *o.clientOid
	}

	if o.marginCoin != nil {
		cp["marginCoin"] = *o.marginCoin
	}

	return cp
}

// Parameter	Type	Required	Description
// symbol	String	Yes	Trading pair
// productType	String	Yes	Product type
// USDT-FUTURES USDT professional futures
// COIN-FUTURES Mixed futures
// USDC-FUTURES USDC professional futures
// SUSDT-FUTURES USDT professional futures demo
// SCOIN-FUTURES Mixed futures demo
// SUSDC-FUTURES USDC professional futures demo
// marginCoin	String	No	Margin coin must be capitalized
// orderId	String	No	Order ID
// Either orderId or clientOid is required.
// If both are present, orderId prevails.
// clientOid	String	No	Customize order ID
// Either orderId or clientOid is required.
// If both are present, orderId prevails.
