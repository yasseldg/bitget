package v2

import (
	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
)

type GetApiKeyList struct {
	c common.InterRestClient

	*request.GetApiKeyList
}

func NewGetApiKeyList(c common.InterRestClient, subAccountUid string) (*GetApiKeyList, error) {

	r, err := request.NewGetApiKeyList(subAccountUid)
	if err != nil {
		return nil, err
	}

	return &GetApiKeyList{
		c:             c,
		GetApiKeyList: r,
	}, nil
}

func (p *GetApiKeyList) Do() (string, error) {

	uri := "/api/v2/user/virtual-subaccount-apikey-list"

	resp, err := p.c.DoGet(uri, p.GetParams())
	if err != nil {
		return "", err
	}

	return resp, nil
}
