package request

import "github.com/yasseldg/bitget/pkg/model/v2/errors"

type GetContractConfig struct {
	productType string
	symbol      *string
}

func NewGetContractConfig(productType string) (*GetContractConfig, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	return &GetContractConfig{
		productType: productType,
	}, nil
}

func (o *GetContractConfig) Symbol(symbol string) *GetContractConfig {
	o.symbol = &symbol
	return o
}

func (o *GetContractConfig) GetParams() map[string]string {
	params := make(map[string]string)

	params["productType"] = o.productType

	if o.symbol != nil {
		params["symbol"] = *o.symbol
	}

	return params
}
