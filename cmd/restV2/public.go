package restV2

import (
	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/go-simple/logs/sLog"
)

func Public() {
	rest := bitget.NewRest(nil)

	contractConfig(rest)
}

func contractConfig(rest bitget.InterRest) {

	client, err := rest.GetContractConfig(constants.ProductType_USDT_FUTURES)
	if err != nil {
		sLog.Error("%s", err)
		return
	}
	client.Symbol("BTCUSDT")

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d \n\n", resp.Code, resp.Msg, resp.RequestTime)

	for k, data := range resp.Data {
		sLog.Info("%d: %+v ", k, data)
	}
}
