package response

import "github.com/yasseldg/bitget/pkg/model"

type ContractResp struct {
	model.BaseResp `json:",inline"`
	Data           []Contract `json:"data"`
}

type Contract struct {
	Symbol              string   `json:"symbol"`
	BaseCoin            string   `json:"baseCoin"`
	QuoteCoin           string   `json:"quoteCoin"`
	BuyLimitPriceRatio  string   `json:"buyLimitPriceRatio"`
	SellLimitPriceRatio string   `json:"sellLimitPriceRatio"`
	FeeRateUpRatio      string   `json:"feeRateUpRatio"`
	MakerFeeRate        string   `json:"makerFeeRate"`
	TakerFeeRate        string   `json:"takerFeeRate"`
	OpenCostUpRatio     string   `json:"openCostUpRatio"`
	SupportMarginCoins  []string `json:"supportMarginCoins"`
	MinTradeUSDT        string   `json:"minTradeUSDT"`
	MinTradeNum         string   `json:"minTradeNum"`
	PriceEndStep        string   `json:"priceEndStep"`
	VolumePlace         string   `json:"volumePlace"`
	PricePlace          string   `json:"pricePlace"`
	SizeMultiplier      string   `json:"sizeMultiplier"`
	SymbolType          string   `json:"symbolType"`
	MaxSymbolOrderNum   string   `json:"maxSymbolOrderNum"`
	MaxProductOrderNum  string   `json:"maxProductOrderNum"`
	MaxPositionNum      string   `json:"maxPositionNum"`
	SymbolStatus        string   `json:"symbolStatus"`
	OffTime             string   `json:"offTime"`
	LimitOpenTime       string   `json:"limitOpenTime"`
	DeliveryTime        string   `json:"deliveryTime"`
	DeliveryStartTime   string   `json:"deliveryStartTime"`
	LaunchTime          string   `json:"launchTime"`
	FundInterval        string   `json:"fundInterval"`
	MinLever            string   `json:"minLever"`
	MaxLever            string   `json:"maxLever"`
	PosLimit            string   `json:"posLimit"`
	MaintainTime        string   `json:"maintainTime"`
}

// {
//     "data": [
//         {
//             "symbol": "BTCUSDT",
//             "baseCoin": "BTC",
//             "quoteCoin": "USDT",
//             "buyLimitPriceRatio": "0.9",
//             "sellLimitPriceRatio": "0.9",
//             "feeRateUpRatio": "0.1",
//             "makerFeeRate": "0.0004",
//             "takerFeeRate": "0.0006",
//             "openCostUpRatio": "0.1",
//             "supportMarginCoins": [
//                 "USDT"
//             ],
//             "minTradeNum": "0.01",
//             "priceEndStep": "1",
//             "volumePlace": "2",
//             "pricePlace": "1",
//             "sizeMultiplier": "0.01",
//             "symbolType": "perpetual",
//             "minTradeUSDT": "5",
//             "maxSymbolOrderNum": "999999",
//             "maxProductOrderNum": "999999",
//             "maxPositionNum": "150",
//             "symbolStatus": "normal",
//             "offTime": "-1",
//             "limitOpenTime": "-1",
//             "deliveryTime": "",
//             "deliveryStartTime": "",
//             "launchTime": "",
//             "fundInterval": "8",
//             "minLever": "1",
//             "maxLever": "125",
//             "posLimit": "0.05",
//             "maintainTime": "1680165535278"
//         }
//     ]
// }
