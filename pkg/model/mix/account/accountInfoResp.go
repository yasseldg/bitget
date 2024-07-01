package account

import "github.com/yasseldg/bitget/pkg/model"

type AccountInfoResp struct {
	model.BaseResp `json:",inline"`
	Data           AccountInfo `json:"data"`
}

type AccountInfo struct {
	UserId      string   `json:"userId"`
	InviterId   string   `json:"inviterId"`
	Ips         string   `json:"ips"`
	Authorities []string `json:"authorities"`
	ParentId    int      `json:"parentId"`
	TraderType  string   `json:"traderType"`
	ChannelCode string   `json:"channelCode"`
	Channel     string   `json:"channel"`
	RegisTime   string   `json:"regisTime"`
}

// {
//     "code": "00000",
//     "msg": "success",
//     "requestTime": 1695808949356,
//     "data": {
//         "userId": "**********",
//         "inviterId": "**********",
//         "ips": "127.0.0.1",
//         "authorities": [
//             "trade",
//             "readonly"
//         ],
//         "parentId": 1,
//         "traderType": "trader",
//         "channelCode": "XXX",
//         "channel": "YYY",
//         "regisTime":"1246566789345"
//     }
// }
