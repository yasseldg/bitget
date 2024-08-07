package request

import "github.com/yasseldg/bitget/pkg/model/v2/errors"

type GetApiKeyList struct {
	subAccountUid string
}

func NewGetApiKeyList(subAccountUid string) (*GetApiKeyList, error) {
	if len(subAccountUid) == 0 {
		return nil, errors.ErrInvalidSubAccountUid
	}

	return &GetApiKeyList{
		subAccountUid: subAccountUid,
	}, nil
}

func (o *GetApiKeyList) GetParams() map[string]string {
	params := make(map[string]string)

	params["subAccountUid"] = o.subAccountUid

	return params
}
