package market

import "github.com/yasseldg/bitget/pkg/model"

type OpenInterestResp struct {
	model.BaseResp `json:",inline"`
	Data           OpenInterest `json:"data"`
}

type OpenInterest struct {
	Symbol    string `json:"symbol"`
	Amount    string `json:"amount"`
	Timestamp string `json:"timestamp"`
}

// {
// 	"code": "00000",
// 	"msg": "success",
// 	"requestTime": 0,
// 	"data": {
// 	  "symbol": "BTCUSDT_UMCBL",
// 	  "amount": "92361.353",
// 	  "timestamp": "1676317051816"
// 	}
// }
