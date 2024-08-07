package v2

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"
	"github.com/yasseldg/bitget/pkg/model/v2/request"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/types/sJson"
)

type GetContractConfig struct {
	c common.InterRestClient

	*request.GetContractConfig
}

func NewGetContractConfig(c common.InterRestClient, productType string) (*GetContractConfig, error) {

	r, err := request.NewGetContractConfig(productType)
	if err != nil {
		return nil, err
	}

	return &GetContractConfig{
		c:                 c,
		GetContractConfig: r,
	}, nil
}

func (p *GetContractConfig) Do() (*response.ContractResp, error) {

	uri := "/api/v2/mix/market/contracts"

	resp, err := p.c.DoGet(uri, p.GetParams())
	if err != nil {
		return nil, err
	}

	respObj := new(response.ContractResp)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
