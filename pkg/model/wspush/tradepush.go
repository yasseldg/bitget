package wspush

import (
	"github.com/yasseldg/bitget/pkg/model"
)

type WsTradePush struct {
	model.BasePush `json:",inline"`
	Data           []WsTrade `json:"data"`
}

type WsTrade struct {
	UnixTs  int64  `json:"ts"`
	Price   string `json:"p"`
	Size    string `json:"c"`
	Side    string `json:"ty"`
	TradeId string `json:"ti"`
}

// > ts	String	Filled time, Unix timestamp format in milliseconds
// > p	String	Trade price
// > c	String	Trade size
// > ty	String	Trade direction, buy, sell
// > ti	String	Trade ID
