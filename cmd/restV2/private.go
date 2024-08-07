package restV2

import (
	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/cmd/config"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/go-simple/logs/sLog"
)

func Private() {
	println()
	rest := bitget.NewRest(config.FullApi())
	// getApiKeyList(rest)
	// accountInfo(rest)
	// accountList(rest)
	// allAccountBalance(rest)
	// getPendingOrders(rest)
	changeMarginMode(rest)
	changePositionMode(rest)

	println("\n")

	rest1 := bitget.NewRest(config.PereApi())
	// getApiKeyList(rest1)
	// accountInfo(rest1)
	// accountList(rest1)
	// allAccountBalance(rest1)
	// getPendingOrders(rest1)
	changeMarginMode(rest1)
	changePositionMode(rest1)
}

func accountList(rest bitget.InterRest) {

	client, err := rest.GetAccountList(constants.ProductType_USDT_FUTURES)
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", resp.Code, resp.Msg, resp.RequestTime)

	sLog.Info("%#v \n", resp.Data)
}

func accountInfo(rest bitget.InterRest) {

	client, err := rest.GetAccountInfo()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", resp.Code, resp.Msg, resp.RequestTime)

	sLog.Info("%#v \n", resp.Data)
}

func allAccountBalance(rest bitget.InterRest) {

	client, err := rest.GetAllAccountBalance()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", resp.Code, resp.Msg, resp.RequestTime)

	sLog.Info("%#v \n", resp.Data)
}

func getApiKeyList(rest bitget.InterRest) {

	client, err := rest.GetApiKeyList("8562705684")
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("%#v \n", resp)
}

func getPendingOrders(rest bitget.InterRest) {

	client, err := rest.GetPendingOrders(constants.ProductType_USDT_FUTURES)
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

	sLog.Info("resp: %+v ", resp)
}

func changeMarginMode(rest bitget.InterRest) {

	client, err := rest.ChangeMarginMode(constants.ProductType_USDT_FUTURES, "BTCUSDT", "USDT", "isolated")
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", resp.Code, resp.Msg, resp.RequestTime)

	sLog.Info("%#v \n", resp.Data)
}

func changePositionMode(rest bitget.InterRest) {

	client, err := rest.ChangePositionMode(constants.ProductType_USDT_FUTURES, "hedge_mode")
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	resp, err := client.Do()
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", resp.Code, resp.Msg, resp.RequestTime)

	sLog.Info("%#v \n", resp.Data)
}
