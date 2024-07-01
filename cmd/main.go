package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/config"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/internal/model"
	"github.com/yasseldg/bitget/pkg/client/ws"
	"github.com/yasseldg/bitget/pkg/model/mix/account"
	"github.com/yasseldg/bitget/pkg/model/mix/market"
	wspush "github.com/yasseldg/bitget/pkg/model/ws"

	"github.com/yasseldg/simplego/sLog"
)

func main() {
	t := time.Now()
	sLog.Info("Cmd Main \n\n")

	// rest
	rest()

	// ws
	// wss()

	fmt.Println()
	sLog.Info("Cmd Main: time: %d Segundos \n\n", time.Since(t)/time.Second)
}

func wss() {

	wss := bitget.NewWsClient()

	wss.Init(func(message string) {
		sLog.Debug("message:" + message)
	}, func(message string) {
		sLog.Error(message)
	}, false)

	uFunc := wss.SubscribeFutures(listTicker, constants.WsChannel_ticker, constants.InstrumentID_BTCUSDT)
	defer uFunc()

	// uFunc = wss.SubscribeFutures(listCandle, constants.WsChannel_candle1m, constants.InstrumentID_BTCUSDT)
	// defer uFunc()

	// uFunc = wss.SubscribeFutures(listTrade, constants.WsChannel_tradeNew, constants.InstrumentID_BTCUSDT)
	// defer uFunc()

	forever := make(chan struct{})
	<-forever
}

func listCandle(msg string) {
	var pushObj wspush.WsCandlePush

	internal.GetPushObj(msg, &pushObj)

	sLog.Debug("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Debug("%d: %v ", k, data)
	}
}

func listTrade(msg string) {
	var pushObj wspush.WsTradePush

	internal.GetPushObj(msg, &pushObj)

	sLog.Debug("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Debug("%d: %v ", k, data)
	}
}

func listTicker(msg string) {
	var pushObj wspush.WsTickerPush

	internal.GetPushObj(msg, &pushObj)

	sLog.Debug("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Debug("%d: %#v ", k, data)
	}
}

func fullApi() *config.ApiCreds {
	return &config.ApiCreds{
		ApiKey:     "bg_a76e95cc3da174f8741b498de7e88594",
		SecretKey:  "f49682195d9dbfbb40fa287f7919dbe928aa91a948d3bf54d29a161d677e83fe",
		PASSPHRASE: "GiXmpxVeX7CboxjgZjpYtTQtkKCtevrH",
	}
}

func pereApi() *config.ApiCreds {
	return &config.ApiCreds{
		ApiKey:     "bg_0bb706efb2350a11055269d07e45439e",
		SecretKey:  "aec2187f4016e03aa7d086b1cb54d54ea5356a2282fc343d627d7cde94f3a5ad",
		PASSPHRASE: "MmCvf5FLxeQgdi3t28B9emgVVZw7FsRY",
	}
}

func rest() {
	// os.Setenv("BITGET_API_KEY", "bg_0bb706efb2350a11055269d07e45439e")
	// os.Setenv("BITGET_API_SECRET", "aec2187f4016e03aa7d086b1cb54d54ea5356a2282fc343d627d7cde94f3a5ad")
	// os.Setenv("BITGET_API_PASS", "MmCvf5FLxeQgdi3t28B9emgVVZw7FsRY")
	// c := bitget.NewClient()

	c := bitget.NewClientWithCreds(pereApi())
	accountInfo(c)
	apiKeyList(c)
	// allAccountBalance(c)

	// historyFundRate(c)
	// currentFundRate(c)
	// openInterest(c)
	// candles(c)
	// tickers(c)
	// ticker(c)
	// depth(c)
	// contracts(c)
}

func accountInfo(c *bitget.Client) {

	resp, err := c.GetAccountService().AccountInfo()
	if err != nil {
		sLog.Debug("[Err] %s", err)
		return
	}

	sLog.Debug("resp: %v ", resp)

	var respObj account.AccountInfoResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", respObj.Code, respObj.Msg, respObj.RequestTime)

	sLog.Info("%#v \n", respObj.Data)
}

func allAccountBalance(c *bitget.Client) {

	resp, err := c.GetAccountService().AllAccountBalance()
	if err != nil {
		sLog.Debug("[Err] %s", err)
		return
	}

	sLog.Debug("resp: %v ", resp)
}

func apiKeyList(c *bitget.Client) {

	resp, err := c.GetUserService().ApiKeyList("8562705684")
	if err != nil {
		sLog.Debug("[Err] %s", err)
		return
	}

	sLog.Debug("resp: %v ", resp)
}

func mixAccount(c *bitget.Client) {

	resp, err := c.GetMixAccountService().Accounts(constants.ProductType_UMCBL)
	if err != nil {
		sLog.Debug("[Err] %s", err)
		return
	}

	sLog.Debug("resp: %v ", resp)

	var respObj account.AccountsResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", respObj.Code, respObj.Msg, respObj.RequestTime)

	sLog.Info("%#v \n", respObj.Data)
}

// Get History Funding Rate
func historyFundRate(c *bitget.Client) {

	resp, err := c.GetMixMarketService().HistoryFundRate(constants.Symbol_BTCUSDT_UMCBL, "", "", "false")
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.HistoryFundingRateResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	arrObjs := []market.HistoryFundingRate{}
	for k, data := range respObj.Data {
		sLog.Info("%d: %#v ", k, data)

		arrObjs = append([]market.HistoryFundingRate{data}, arrObjs...)
	}
	fmt.Println()

	for k, data := range arrObjs {
		sLog.Info("%d: %#v ", k, data)
	}
	fmt.Println()
}

// Get Current Funding Rate
func currentFundRate(c *bitget.Client) {

	resp, err := c.GetMixMarketService().CurrentFundRate(constants.Symbol_BTCUSDT_UMCBL)
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.FundingRateResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", respObj.Code, respObj.Msg, respObj.RequestTime)

	sLog.Info("%#v \n", respObj.Data)
}

// Get Open Interest
func openInterest(c *bitget.Client) {

	resp, err := c.GetMixMarketService().OpenInterest(constants.Symbol_BTCUSDT_UMCBL)
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.OpenInterestResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d ", respObj.Code, respObj.Msg, respObj.RequestTime)

	sLog.Info("%#v \n", respObj.Data)
}

// Get Candle Data
func candles(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Candles(constants.Symbol_BTCUSDT_UMCBL, constants.CandleInterval_5m, "1673696121000", "1673796121000")
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.CandlesResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	for k, data := range respObj {
		sLog.Info("%d: %v ", k, data)
	}
}

// Get All Symbol Ticker
func tickers(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Tickers(constants.ProductType_UMCBL)
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.TickersResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)

	for k, data := range respObj.Data {
		sLog.Info("%d: %v ", k, data)
	}
}

// Get Single Symbol Ticker
func ticker(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Ticker(constants.Symbol_BTCUSDT_UMCBL)
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.TickerResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)
	sLog.Info("data: %v ", respObj.Data)
}

// Get Depth, Asks and Bids
func depth(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Depth(constants.Symbol_BTCUSDT_UMCBL, "15")
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.DepthtResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)
	sLog.Info("data: %v ", respObj.Data)
}

// Get All Symbols
func contracts(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Contracts(constants.ProductType_UMCBL)
	if err != nil {
		sLog.Error("%s", err)
		return
	}

	var respObj market.ContractResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		sLog.Error("json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
		return
	}

	sLog.Info("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)

	patterns := []string{"BTCUSDT", "ETHUSDT", "ADAUSDT", "BNB", "DOGE", "LTC", "MATIC"}

	for k, data := range respObj.Data {
		if FindPatterns(data.Symbol, patterns...) {
			sLog.Info("%d: %v ", k, data)
		} else {
			sLog.Debug("%d: %v ", k, data)
		}
	}
}

func FindPatterns(path string, patterns ...string) bool {

	if len(patterns) == 0 {
		return true
	}

	for _, p := range patterns {
		if strings.Contains(path, p) {
			return true
		}
	}

	return false
}

func mainOLD() {
	client := new(ws.BitgetWsClient).Init(true, func(message string) {
		fmt.Println("default error:" + message)
	}, func(message string) {
		fmt.Println("default error:" + message)
	})

	var channelsDef []model.SubscribeReq
	subReqDef1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channelsDef = append(channelsDef, subReqDef1)
	client.SubscribeDef(channelsDef)

	var channels []model.SubscribeReq
	subReq1 := model.SubscribeReq{
		InstType: "UMCBL",
		Channel:  "account",
		InstId:   "default",
	}
	channels = append(channels, subReq1)
	client.Subscribe(channels, func(message string) {
		fmt.Println("appoint:" + message)
	})
	fmt.Println("Press ENTER to unsubscribe and stop...")
	fmt.Scanln()
}
