package request

import "github.com/yasseldg/bitget/pkg/model/v2/errors"

type GetAccountList struct {
	productType string
}

func NewGetAccountList(productType string) (*GetAccountList, error) {
	if len(productType) == 0 {
		return nil, errors.ErrInvalidProductType
	}

	return &GetAccountList{
		productType: productType,
	}, nil
}

func (o *GetAccountList) GetParams() map[string]string {
	params := make(map[string]string)

	params["productType"] = o.productType

	return params
}
