package market

import "github.com/yasseldg/bitget/pkg/model"

type FundingRateResp struct {
	model.BaseResp `json:",inline"`
	Data           FundingRate `json:"data"`
}

type FundingRate struct {
	Symbol      string `json:"symbol"`
	FundingRate string `json:"fundingRate"`
}

// {
// 	"code": "00000",
// 	"msg": "success",
// 	"requestTime": 0,
// 	"data": {
// 	  "symbol": "BTCUSDT_UMCBL",
// 	  "fundingRate": "0.000097"
// 	}
// }

type HistoryFundingRateResp struct {
	model.BaseResp `json:",inline"`
	Data           []HistoryFundingRate `json:"data"`
}

type HistoryFundingRate struct {
	FundingRate `json:",inline"`
	SettleTime  string `json:"settleTime"`
}

// {
// 	"code": "00000",
// 	"msg": "success",
// 	"requestTime": 0,
// 	"data": [
// 	  {
// 		"symbol": "BTCUSDT",
// 		"fundingRate": "0.000096",
// 		"settleTime": "1676300400000"
// 	  },

// 	  {
// 		"symbol": "BTCUSDT",
// 		"fundingRate": "0.000104",
// 		"settleTime": "1675983600000"
// 	  },
// 	  {
// 		"symbol": "BTCUSDT",
// 		"fundingRate": "0.000099",
// 		"settleTime": "1675810800000"
// 	  },
// 	  {
// 		"symbol": "BTCUSDT",
// 		"fundingRate": "0.0001",
// 		"settleTime": "1675782000000"
// 	  },
// 	  {
// 		"symbol": "BTCUSDT",
// 		"fundingRate": "0.000091",
// 		"settleTime": "1675753200000"
// 	  }
// 	]
// }
