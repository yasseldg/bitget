package olds

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/pkg/model/mix/market"
	"github.com/yasseldg/bitget/pkg/model/v2/response"
	"github.com/yasseldg/go-simple/logs/sLog"
)

func mixAccount(c *bitget.Client) {

	resp, err := c.GetMixAccountService().Accounts(constants.ProductType_UMCBL)
	if err != nil {
		sLog.Debug("[Err] %s", err)
		return
	}

	sLog.Debug("resp: %v ", resp)

	var respObj response.AccountsResp

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

	var respObj response.ContractResp

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
