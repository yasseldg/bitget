package common

type OrderIds struct {
	OrderId   string `json:"orderId"`
	ClientOId string `json:"clientOId"`
}

type Order struct {
	OrderIds `json:",inline"`

	InstId string `json:"instId"`
	Symbol string `json:"symbol"`

	Side        string `json:"side"`
	Price       string `json:"price"`
	PriceAvg    string `json:"priceAvg"`
	Size        string `json:"size"`
	Status      string `json:"status"`
	OrderType   string `json:"orderType"`
	OrderSource string `json:"orderSource"`

	PosMode    string `json:"posMode"`
	PosSide    string `json:"posSide"`
	TradeSide  string `json:"tradeSide"`
	MarginCoin string `json:"marginCoin"`
	MarginMode string `json:"marginMode"`
	CTime      string `json:"cTime"`
	UTime      string `json:"uTime"`

	ReduceOnly string `json:"reduceOnly"`

	PresetStopSurplusPrice string `json:"presetStopSurplusPrice"`
	PresetStopLossPrice    string `json:"presetStopLossPrice"`

	EnterPointSource string `json:"enterPointSource"`

	Force    string `json:"force"`
	Leverage string `json:"leverage"`
}
