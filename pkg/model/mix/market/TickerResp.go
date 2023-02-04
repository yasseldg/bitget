package market

import "github.com/yasseldg/bitget/pkg/model"

type TickersResp struct {
	model.BaseResp `json:",inline"`
	Data           []Ticker `json:"data"`
}

type TickerResp struct {
	model.BaseResp `json:",inline"`
	Data           Ticker `json:"data"`
}

type Ticker struct {
	Symbol             string `json:"symbol"`
	Last               string `json:"last"`
	BestAsk            string `json:"bestAsk"`
	BestBid            string `json:"bestBid"`
	BidSz              string `json:"bidSz"`
	AskSz              string `json:"askSz"`
	High24h            string `json:"high24h"`
	Low24h             string `json:"low24h"`
	Timestamp          string `json:"timestamp"`
	PriceChangePercent string `json:"priceChangePercent"`
	BaseVolume         string `json:"baseVolume"`
	UsdtVolume         string `json:"usdtVolume"`
	OpenUtc            string `json:"openUtc"`
	ChgUtc             string `json:"chgUtc"`
}

// {
// 	"code": "00000",
// 	"msg": "success",
// 	"requestTime": 0,
// 	"data": {
// 	  "symbol": "BTCUSDT_UMCBL",
// 	  "last": "23838.5",
// 	  "bestAsk": "23839",
// 	  "bestBid": "23838.5",
// 	  "bidSz": "15.783",
// 	  "askSz": "0.06",
// 	  "high24h": "24266",
// 	  "low24h": "22750.5",
// 	  "timestamp": "1675359101461",
// 	  "priceChangePercent": "0.00006",
// 	  "baseVolume": "225794.551",
// 	  "quoteVolume": "5333431960.3055",
// 	  "usdtVolume": "5333431960.3055",
// 	  "openUtc": "23723",
// 	  "chgUtc": "0.00487"
// 	}
// }
