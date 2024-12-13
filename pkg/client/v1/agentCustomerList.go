package v1

import (
	"fmt"

	"github.com/yasseldg/bitget/internal/common"

	"github.com/yasseldg/bitget/pkg/model/v1/request"
	"github.com/yasseldg/bitget/pkg/model/v1/response"

	"github.com/yasseldg/go-simple/types/sJson"
)

type AgentCustomerList struct {
	c common.InterRestClient

	*request.CustomerList
}

func NewAgentCustomerList(c common.InterRestClient) (*AgentCustomerList, error) {
	return &AgentCustomerList{
		c:            c,
		CustomerList: request.NewCustomerList(),
	}, nil
}

func (p *AgentCustomerList) Do() (*response.CustomerList, error) {

	params, err := sJson.ToString(p.GetParams())
	if err != nil {
		return nil, fmt.Errorf("sJson.ToString(): %s", err)
	}

	uri := "/api/broker/v1/agent/customerList"

	resp, err := p.c.DoPost(uri, params)
	if err != nil {
		return nil, err
	}

	respObj := new(response.CustomerList)

	err = sJson.ToObj(resp, respObj)
	if err != nil {
		return nil, fmt.Errorf("sJson.ToObj(): %s", err)
	}

	return respObj, nil
}
