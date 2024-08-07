package response

import "github.com/yasseldg/bitget/pkg/model"

type MarginModeResp struct {
	model.BaseResp `json:",inline"`
	Data           MarginMode `json:"data"`
}

type MarginMode struct {
	Symbol        string `json:"symbol"`
	MarginCoin    string `json:"marginCoin"`
	LongLeverage  string `json:"longLeverage"`
	ShortLeverage string `json:"shortLeverage"`
	MarginMode    string `json:"marginMode"`
}

// {
//     "code": "00000",
//     "data": {
//         "symbol": "BTCUSDT",
//         "marginCoin": "USDT",
//         "longLeverage": "25",
//         "shortLeverage": "20",
//         "marginMode": "isolated"
//     },
//     "msg": "success",
//     "requestTime": 1627293445916
// }
