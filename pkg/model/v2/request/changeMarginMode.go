package request

import (
	"github.com/yasseldg/bitget/pkg/model/v2/errors"
)

type ChangeMarginMode struct {
	productType string
	symbol      string
	marginCoin  string
	marginMode  string
}

func NewChangeMarginMode(productType, symbol, marginCoin, marginMode string) (*ChangeMarginMode, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	if len(symbol) == 0 {
		return nil, errors.ErrInvalidSymbol
	}

	if len(marginCoin) == 0 {
		return nil, errors.ErrInvalidMarginCoin
	}

	if len(marginMode) == 0 {
		return nil, errors.ErrInvalidMarginMode
	}

	return &ChangeMarginMode{
		productType: productType,
		symbol:      symbol,
		marginCoin:  marginCoin,
		marginMode:  marginMode,
	}, nil
}

func (o *ChangeMarginMode) GetParams() map[string]string {
	params := make(map[string]string)

	params["productType"] = o.productType
	params["symbol"] = o.symbol
	params["marginCoin"] = o.marginCoin
	params["marginMode"] = o.marginMode

	return params
}

// Parameter	Type	Required	Description
// symbol	String	Yes	Trading pair. e.g. BTCUSDT
// productType	String	Yes	Product type
// USDT-FUTURES USDT professional futures
// COIN-FUTURES Mixed futures
// USDC-FUTURES USDC professional futures
// SUSDT-FUTURES USDT professional futures demo
// SCOIN-FUTURES Mixed futures demo
// SUSDC-FUTURES USDC professional futures demo
// marginCoin	String	Yes	Margin coin, must be capitalized
// marginMode	String	Yes	Margin mode.
// isolated: isolated margin mode
// crossed: crossed margin mode
