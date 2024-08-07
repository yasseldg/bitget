package wspush

import (
	"github.com/yasseldg/bitget/pkg/model"
	"github.com/yasseldg/bitget/pkg/model/common"
)

type WsOrderPush struct {
	model.BasePush `json:",inline"`
	Data           []WsOrder `json:"data"`
}

type WsOrder struct {
	common.Order `json:",inline"`

	AccBaseVolume string `json:"accBaseVolume"`
	FeeDetail     []struct {
		FeeCoin string `json:"feeCoin"`
		Fee     string `json:"fee"`
	} `json:"feeDetail"`
	FillFee          string `json:"fillFee"`
	FillFeeCoin      string `json:"fillFeeCoin"`
	FillNotionalUsd  string `json:"fillNotionalUsd"`
	FillPrice        string `json:"fillPrice"`
	BaseVolume       string `json:"baseVolume"`
	FillTime         string `json:"fillTime"`
	Force            string `json:"force"`
	Leverage         string `json:"leverage"`
	NotionalUsd      string `json:"notionalUsd"`
	Pnl              string `json:"pnl"`
	StpMode          string `json:"stpMode"`
	EnterPointSource string `json:"enterPointSource"`
	OrderScope       string `json:"orderScope"`
	OrderSide        string `json:"orderSide"`
}

// {
//     "action": "snapshot",
//     "arg": {
//         "instType": "USDT-FUTURES",
//         "channel": "orders",
//         "instId": "default"
//     },
//     "data": [
//         {
//             "accBaseVolume": "0.01",
//             "cTime": "1695718781129",
//             "clientOId": "1",
//             "feeDetail": [
//                 {
//                     "feeCoin": "USDT",
//                     "fee": "-0.162003"
//                 }
//             ],
//             "fillFee": "-0.162003",
//             "fillFeeCoin": "USDT",
//             "fillNotionalUsd": "270.005",
//             "fillPrice": "27000.5",
//             "baseVolume": "0.01",
//             "fillTime": "1695718781146",
//             "force": "gtc",
//             "instId": "BTCUSDT",
//             "leverage": "20",
//             "marginCoin": "USDT",
//             "marginMode": "crossed",
//             "notionalUsd": "270",
//             "orderId": "1",
//             "orderType": "market",
//             "pnl": "0",
//             "posMode": "hedge_mode",
//             "posSide": "long",
//             "price": "0",
//             "priceAvg": "27000.5",
//             "reduceOnly": "no",
//             "stpMode": "cancel_taker",
//             "side": "buy",
//             "size": "0.01",
//             "enterPointSource": "WEB",
//             "status": "filled",
//             "orderScope": "T",
//             "orderId": "1111111111",
//             "orderSide": "open",
//             "uTime": "1695718781146"
//         }
//     ],
//     "ts": 1695718781206
// }
