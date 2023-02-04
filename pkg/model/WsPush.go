package model

import "github.com/yasseldg/bitget/internal/model"

type BasePush struct {
	Action string             `json:"action"`
	Arg    model.SubscribeReq `json:"arg"`
	// Data   []interface{} `json:"data"`
}
