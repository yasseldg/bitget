package wspush

import (
	"github.com/yasseldg/bitget/pkg/model"
)

type WsTickerPush struct {
	model.BasePush `json:",inline"`
	Data           []WsTicker `json:"data"`
}

type WsTicker struct {
	SystemTime  int64  `json:"systemTime"`
	InstId      string `json:"instId"`
	CapitalRate string `json:"capitalRate"`
	Holding     string `json:"holding"`

	// Last               string `json:"last"`
	// High24h            string `json:"high24h"`
	// Low24h             string `json:"low24h"`
	// PriceChange        string `json:"priceChange"`
	// PriceChangePercent string `json:"priceChangePercent"`
	// MarketPrice        string `json:"marketPrice"`
	// IndexPrice         string `json:"indexPrice"`
	// BaseVolume         string `json:"baseVolume"`
	// QuoteVolume        string `json:"quoteVolume"`
	// OpenUtc            string `json:"openUtc"`
	// ChgUtc             string `json:"chgUtc"`
	// SymbolType         int64  `json:"symbolType"`
	// SymbolId           string `json:"symbolId"`
	// DeliveryPrice      string `json:"deliveryPrice"`
	// BidSz              string `json:"bidSz"`
	// AskSz              string `json:"askSz"`
	// BestBid            string `json:"bestBid"`
	// BestAsk            string `json:"bestAsk"`
	// NextSettleTime     int64  `json:"nextSettleTime"`
}

// > instId 	String 	Instrument Name
// > last 	String 	Last traded price
// > bestAsk 	String 	Best ask price
// > bestBid 	String 	Best bid price
// > high24h 	String 	Highest price in the past 24 hours
// > low24h 	String 	Lowest price in the past 24 hours
// > priceChangePercent 	String 	Price change int the past 24 hours
// > capitalRate 	String 	Funding rate
// > nextSettleTime 	String 	The next fund rate settlement time timestamp milliseconds
// > systemTime 	String 	system time
// > marketPrice 	String 	Market price
// > indexPrice 	String 	Index price
// > holding 	String 	Open interest
// > baseVolume 	String 	24h trading volume, with a unit of base.
// > quoteVolume 	String 	24h trading volume, with a unit of quote
// > openUtc 	String 	Open price at UTC 00:00
// > chgUTC 	String 	Price change since UTC 00:00
// > symbolType 	String 	SymbolType: delivery: Settled Futures; perpetual: Perpetual Futures
// > symbolId 	String 	Symbol Id
// > deliveryPrice 	String 	Delivery price - 0 when SymbolType=perpetual
// > bidSz 	String 	Best bid size
// > askSz 	String 	Best ask size
