package spot

import (
	"fmt"
	"testing"

	"github.com/yasseldg/bitget/pkg/model/spot/account"
)

func TestSpotAccountClient_Assets(t *testing.T) {
	client := new(SpotAccountClient).Init()

	resp, err := client.Assets()

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestSpotAccountClient_Bills(t *testing.T) {
	client := new(SpotAccountClient).Init()

	req := account.BillsReq{CoinId: "1", Before: "777031099461570560"}

	resp, err := client.Bills(req)

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}

func TestSpotAccountClient_TransferRecords(t *testing.T) {
	client := new(SpotAccountClient).Init()

	resp, err := client.TransferRecords("2", "USDT_MIX", "10", "", "")

	if err != nil {
		println(err.Error())
	}
	fmt.Println(resp)
}
