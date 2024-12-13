package response

import "github.com/yasseldg/bitget/pkg/model"

type CustomerList struct {
	model.BaseResp `json:",inline"`
	Data           []Customer `json:"data"`
}

type Customer struct {
	Uid          string `json:"uid"`
	RegisterTime string `json:"registerTime"`
}

// {
//  "code":"00000",
//  "msg":"success",
//  "requestTime":163123213132,
//  "data":[
//    {
//      "uid":"435435345",
//      "registerTime":"1679991960110"
//    },
//    {
//      "uid":"435435345",
//      "registerTime":"1679991960110"
//    }
//  ]
// }
