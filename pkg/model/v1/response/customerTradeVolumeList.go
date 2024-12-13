package response

import "github.com/yasseldg/bitget/pkg/model"

type CustomerTradeVolumeList struct {
	model.BaseResp `json:",inline"`
	Data           []TradeVolume `json:"data"`
}

type TradeVolume struct {
	Uid    string `json:"uid"`
	Volumn string `json:"volumn"`
	Time   string `json:"time"`
}

// {
//  "code":"00000",
//  "msg":"success",
//  "requestTime":163123213132,
//  "data":[
//    {
//      "uid":"435435345",
//      "volumn": "34234.4",
//      "time":"1679991960110"
//    },
//    {
//      "uid":"435435345",
//      "volumn": "34234.4",
//      "time":"1679991960110"
//    }
//  ]
// }
