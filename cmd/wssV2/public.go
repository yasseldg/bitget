package wssV2

import (
	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/constants"
	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/pkg/model/wspush"
	"github.com/yasseldg/go-simple/logs/sLog"
)

func Public() {
	// public
	wss := bitget.NewWsClient()

	wss.Init(func(message string) {
		sLog.Info("message:" + message)
	}, func(message string) {
		sLog.Error(message)
	}, false, nil)

	// uFunc = wss.SubscribeFutures(listTrade, constants.WsChannel_tradeNew, constants.InstrumentID_BTCUSDT)
	// defer uFunc()

	// uFunc_1 := wss.SubscribeFutures(listTicker, constants.WsChannel_ticker, constants.InstrumentID_BTCUSDT)
	// defer uFunc_1()

	uFunc_2 := wss.SubscribeFutures(listCandle, constants.WsChannel_candle1m, constants.InstrumentID_BTCUSDT)
	defer uFunc_2()

	forever := make(chan struct{})
	<-forever
}

func listCandle(msg string) {
	var pushObj wspush.WsCandlePush

	internal.GetPushObj(msg, &pushObj)

	sLog.Debug("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Info("%d: %v ", k, data)
	}
}

func listTrade(msg string) {
	var pushObj wspush.WsTradePush

	internal.GetPushObj(msg, &pushObj)

	sLog.Debug("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Info("%d: %v ", k, data)
	}
}

func listTicker(msg string) {
	var pushObj wspush.WsTickerPush

	internal.GetPushObj(msg, &pushObj)

	sLog.Debug("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Info("%d: %#v ", k, data)
	}
}
