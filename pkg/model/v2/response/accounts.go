package response

import "github.com/yasseldg/bitget/pkg/model"

type AccountsResp struct {
	model.BaseResp `json:",inline"`
	Data           []Account `json:"data"`
}

type Account struct {
	MarginCoin        string `json:"marginCoin"`
	Locked            string `json:"locked"`
	Available         string `json:"available"`
	CrossMaxAvailable string `json:"crossMaxAvailable"`
	FixedMaxAvailable string `json:"fixedMaxAvailable"`
	MaxTransferOut    string `json:"maxTransferOut"`
	Equity            string `json:"equity"`
	UsdtEquity        string `json:"usdtEquity"`
	BtcEquity         string `json:"btcEquity"`
	CrossRiskRate     string `json:"crossRiskRate"`
	UnrealizedPL      string `json:"unrealizedPL"`
	Bonus             string `json:"bonus"`
}

//   "data":[
//     {
//       "marginCoin":"USDT",
//       "locked":"0.31876482",
//       "available":"10575.26735771",
//       "crossMaxAvailable":"10580.56434289",
//       "fixedMaxAvailable":"10580.56434289",
//       "maxTransferOut":"10572.92904289",
//       "equity":"10582.90265771",
//       "usdtEquity":"10582.902657719473",
//       "btcEquity":"0.204885807029",
//       "crossRiskRate": "0",
//       "unrealizedPL": null,
//       "bonus": "0"
//     }
//   ],
