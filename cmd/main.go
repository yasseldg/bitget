package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/yasseldg/bitget"
	sLog "github.com/yasseldg/bitget/as/log"

	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/internal/model"
	"github.com/yasseldg/bitget/pkg/client/ws"
	"github.com/yasseldg/bitget/pkg/model/mix/market"
	wspush "github.com/yasseldg/bitget/pkg/model/ws"
)

func main() {
	t := time.Now()
	log.Print("Test Main \n\n")

	// rest
	// rest()

	// ws
	wss()

	fmt.Println()
	log.Printf("Test Main: time: %d Segundos \n\n", time.Since(t)/time.Second)
}

func wss() {

	wss := bitget.NewWsClient()

	wss.Init(func(message string) {
		sLog.Debug("message:" + message)
	}, func(message string) {
		sLog.Error(message)
	}, false)

	uFunc := wss.SubscribeFutures(listCandle, constants.WsChannel_candle1m, constants.InstrumentID_BTCUSDT)
	defer uFunc()

	uFunc = wss.SubscribeFutures(listTrade, constants.WsChannel_tradeNew, constants.InstrumentID_BTCUSDT)
	defer uFunc()

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

func rest() {
	c := bitget.NewClient()
	first(c)

	candles(c)
	tickers(c)
	ticker(c)
	depth(c)
	contracts(c)
}

func first(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Candles(constants.Symbol_BTCUSDT_UMCBL, constants.CandleInterval_1m, "1673036121000", "1673039121000")
	if err != nil {
		log.Printf("[Err] %s", err)
		return
	}

	log.Printf("resp: %v ", resp)
}

func candles(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Candles(constants.Symbol_BTCUSDT_UMCBL, constants.CandleInterval_5m, "1673036121000", "1673039121000")
	if err != nil {
		log.Printf("[Err] %s", err)
		return
	}

	var respObj market.CandlesResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		log.Printf("[Err] json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	for k, data := range respObj {
		log.Printf("%d: %v ", k, data)
	}
}

func tickers(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Tickers(constants.ProductType_UMCBL)
	if err != nil {
		log.Printf("[Err] %s", err)
		return
	}

	var respObj market.TickersResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		log.Printf("[Err] json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	log.Printf("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)

	for k, data := range respObj.Data {
		log.Printf("%d: %v ", k, data)
	}
}

func ticker(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Ticker(constants.Symbol_BTCUSDT_UMCBL)
	if err != nil {
		log.Printf("[Err] %s", err)
		return
	}

	var respObj market.TickerResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		log.Printf("[Err] json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
	}

	log.Printf("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)
	log.Printf("data: %v ", respObj.Data)
}

func depth(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Depth(constants.Symbol_BTCUSDT_UMCBL, "15")
	if err != nil {
		log.Printf("[Err] %s", err)
		return
	}

	var respObj market.DepthtResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		log.Printf("[Err] json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
		return
	}

	log.Printf("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)
	log.Printf("data: %v ", respObj.Data)
}

func contracts(c *bitget.Client) {

	resp, err := c.GetMixMarketService().Contracts(constants.ProductType_UMCBL)
	if err != nil {
		log.Printf("[Err] %s", err)
		return
	}

	var respObj market.ContractResp

	err = json.Unmarshal([]byte(resp), &respObj)
	if err != nil {
		log.Printf("[Err] json.Unmarshal([]byte(resp), respObj)  --  error: %s", err)
		return
	}

	log.Printf("code: %s  --  msg: %s  --  reqTime: %d \n\n", respObj.Code, respObj.Msg, respObj.RequestTime)

	for k, data := range respObj.Data {
		log.Printf("%d: %v ", k, data)
	}
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
