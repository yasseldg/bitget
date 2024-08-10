package wssV2

import (
	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/cmd/config"
	"github.com/yasseldg/bitget/internal"
	"github.com/yasseldg/bitget/pkg/model/wspush"
	"github.com/yasseldg/go-simple/logs/sLog"
)

func Private() {

	// private
	wss_priv := bitget.NewWsClient()

	wss_priv.Init(func(message string) {
		sLog.Info("message:" + message)
	}, func(message string) {
		sLog.Error(message)
	}, true, config.FullApi())

	uFunc_3 := wss_priv.SubscribeFuturesOrder(listOrders)
	defer uFunc_3()

	forever := make(chan struct{})
	<-forever
}

func listOrders(msg string) {
	var pushObj wspush.WsOrderPush

	internal.GetPushObj(msg, &pushObj)

	sLog.Info("Arg: %v  --  Action: %s", pushObj.Arg, pushObj.Action)

	for k, data := range pushObj.Data {
		sLog.Info("%d: %+v ", k, data)
	}
}
