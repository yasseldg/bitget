package response

import "github.com/yasseldg/bitget/pkg/model"

type PositionModeResp struct {
	model.BaseResp `json:",inline"`
	Data           PositionMode `json:"data"`
}

type PositionMode struct {
	PosMode string `json:"posMode"`
}

// {
//     "code": "00000",
//     "msg": "success",
//     "data": {
//         "posMode": "one_way_mode"
//     },
//     "requestTime": 1627293445916
// }
