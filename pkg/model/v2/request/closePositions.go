package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"

	"github.com/yasseldg/go-simple/logs/sLog"
)

type ClosePositions struct {
	productType string

	symbol   *string
	holdSide *string
}

func NewClosePositions(productType string) (*ClosePositions, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	return &ClosePositions{
		productType: productType,
	}, nil
}

func (o *ClosePositions) Log() {
	sLog.Info("ClosePositions: %v ", o.GetParams())
}

func (o *ClosePositions) Symbol(symbol string) *ClosePositions {
	o.symbol = &symbol
	return o
}

func (o *ClosePositions) HoldSide(holdSide string) *ClosePositions {
	o.holdSide = &holdSide
	return o
}

// Params returns the order parameters
func (o *ClosePositions) GetParams() map[string]string {
	cp := make(map[string]string)

	cp["productType"] = o.productType

	if o.symbol != nil {
		cp["symbol"] = *o.symbol
	}

	if o.holdSide != nil {
		cp["holdSide"] = *o.holdSide
	}

	return cp
}

// Parameter	Type	Required	Description
// symbol	String	No	Trading pair
// holdSide	String	Optional	Position direction
// 1. In one-way position mode(buy or sell): This field should be left blank. Will be ignored if filled in.
// 2. In hedge-mode position(open or close): All positions will be closed if the field is left blank; Positions of the specified direction will be closed is the field is filled in.
// long: Long position; short: Short position
// productType	String	Yes	Product type
// USDT-FUTURES USDT professional futures
// COIN-FUTURES Mixed futures
// USDC-FUTURES USDC professional futures
// SUSDT-FUTURES USDT professional futures demo
// SCOIN-FUTURES Mixed futures demo
// SUSDC-FUTURES USDC professional futures demo
