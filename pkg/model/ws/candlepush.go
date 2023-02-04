package wspush

import (
	"github.com/yasseldg/bitget/pkg/model"
	"github.com/yasseldg/bitget/pkg/model/mix/market"
)

type WsCandlePush struct {
	model.BasePush `json:",inline"`
	Data           market.CandlesResp `json:"data"`
}

// type WsCandle struct {
// 	UnixTs string
// 	Open   string
// 	High   string
// 	Low    string
// 	Close  string
// 	Volume string
// }

// > ts	String	Data generation time, Unix timestamp format in milliseconds
// > o	String	Open price
// > h	String	highest price
// > l	String	Lowest price
// > c	String	Close price
// > baseVol	String	Trading volume, with a unit of base coin.
