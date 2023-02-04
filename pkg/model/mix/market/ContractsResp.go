package market

import "github.com/yasseldg/bitget/pkg/model"

type ContractResp struct {
	model.BaseResp `json:",inline"`
	Data           []Contract `json:"data"`
}

type Contract struct {
	Symbol     string `json:"symbol"`
	BaseCoin   string `json:"baseCoin"`
	QuoteCoin  string `json:"quoteCoin"`
	SymbolType string `json:"symbolType"`
}

// {
// 	"code": "00000",
// 	"msg": "success",
// 	"requestTime": 0,
// 	"data": [
// 		{
// 			"symbol": "BTCUSDT_UMCBL",
// 			"makerFeeRate": "0.0002",
// 			"takerFeeRate": "0.0006",
// 			"feeRateUpRatio": "0.005",
// 			"openCostUpRatio": "0.01",
// 			"quoteCoin": "USDT",
// 			"baseCoin": "BTC",
// 			"buyLimitPriceRatio": "0.01",
// 			"sellLimitPriceRatio": "0.01",
// 			"supportMarginCoins": [
// 				"USDT"
// 			],
// 			"minTradeNum": "0.001",
// 			"priceEndStep": "5",
// 			"volumePlace": "3",
// 			"pricePlace": "1",
// 			"sizeMultiplier": "0.001",
// 			"symbolType": "perpetual",
// 			"symbolStatus": "normal",
// 			"offTime": "-1",
// 			"limitOpenTime": "-1"
// 		}
// 	]
// }
