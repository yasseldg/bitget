package request

import "github.com/yasseldg/bitget/pkg/model/v2/errors"

type ChangePositionMode struct {
	productType string
	posMode     string
}

func NewChangePositionMode(productType, posMode string) (*ChangePositionMode, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	if len(posMode) == 0 {
		return nil, errors.ErrInvalidPosMode
	}

	return &ChangePositionMode{
		productType: productType,
		posMode:     posMode,
	}, nil
}

func (o *ChangePositionMode) GetParams() map[string]string {
	params := make(map[string]string)

	params["productType"] = o.productType
	params["posMode"] = o.posMode

	return params
}
