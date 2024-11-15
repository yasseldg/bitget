package response

import (
	"github.com/yasseldg/bitget/pkg/model"
)

type PositionResp struct {
	model.BaseResp `json:",inline"`
	Data           Position `json:"data"`
}

type PositionsResp struct {
	model.BaseResp `json:",inline"`
	Data           PositionsData `json:"data"`
}

type PositionsData struct {
	SuccessList []success `json:"successList"`
	FailureList []failure `json:"failureList"`
}

type Position struct {
	// common.Position `json:",inline"`

	BaseVolume   string `json:"baseVolume"`
	Fee          string `json:"fee"`
	TotalProfits string `json:"totalProfits"`
	QuoteVolume  string `json:"quoteVolume"`
}

type success struct {
	OrderId   string `json:"orderId"`
	ClientOId string `json:"clientOid"`
}

type failure struct {
	success   `json:",inline"`
	ErrorMsg  string `json:"errorMsg"`
	ErrorCode string `json:"errorCode"`
}

// {
//     "code": "00000",
//     "data": {
//         "successList": [
//             {
//                 "orderId": "123",
//                 "clientOid": "xxxxx"
//             }
//         ],
//         "failureList": [
//             {
//                 "orderId": "1234",
//                 "clientOid": "321",
//                 "errorMsg": "xxx",
//                 "errorCode": "xxxx"
//             }
//         ]
//     },
//     "msg": "success",
//     "requestTime": 1627293504612
// }
