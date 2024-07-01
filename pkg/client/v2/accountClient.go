package v2

import (
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/internal/common"
)

type AccountClient struct {
	BitgetRestClient *common.BitgetRestClient
}

func (p *AccountClient) Init() *AccountClient {
	p.BitgetRestClient = new(common.BitgetRestClient).Init()
	return p
}

func (p *AccountClient) AllAccountBalance() (string, error) {

	//  "/api/v2/account/all-account-balance"

	params := internal.NewParams()

	uri := constants.V2_Account + "/all-account-balance"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}

func (p *AccountClient) AccountInfo() (string, error) {

	//  "/api/v2/spot/account/info"

	params := internal.NewParams()

	uri := constants.V2_SpotAccount + "/info"

	resp, err := p.BitgetRestClient.DoGet(uri, params)

	return resp, err
}
