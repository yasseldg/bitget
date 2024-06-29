package constants

const (
	/*
	  http headers
	*/
	ContentType        = "Content-Type"
	BgAccessKey        = "ACCESS-KEY"
	BgAccessSign       = "ACCESS-SIGN"
	BgAccessTimestamp  = "ACCESS-TIMESTAMP"
	BgAccessPassphrase = "ACCESS-PASSPHRASE"
	ApplicationJson    = "application/json"

	EN_US  = "en_US"
	ZH_CN  = "zh_CN"
	LOCALE = "locale="

	/*
	  http methods
	*/
	GET    = "GET"
	POST   = "POST"
	DELETE = "DELETE"

	/*
	 others
	*/
	ResultDataJsonString = "resultDataJsonString"
	ResultPageJsonString = "resultPageJsonString"

	/**
	 * SPOT URL
	 */
	SpotPublic  = "/api/spot/v1/public"
	SpotMarket  = "/api/spot/v1/market"
	SpotAccount = "/api/spot/v1/account"
	SpotTrade   = "/api/spot/v1/trade"

	/**
	 * MIX URL
	 */
	MixPlan     = "/api/mix/v1/plan"
	MixMarket   = "/api/mix/v1/market"
	MixAccount  = "/api/mix/v1/account"
	MixOrder    = "/api/mix/v1/order"
	MixPosition = "/api/mix/v1/position"
	MixTrace    = "/api/mix/v1/trace"

	/**
	 * V2 URL
	 */
	V2_User    = "/api/v2/user"
	V2_Account = "/api/v2/account"

	/**
	 * broker url
	 */
	BrokerAccount = "/api/broker/v1/account"

	/**
	websocket
	*/
	WsAuthMethod        = "GET"
	WsAuthPath          = "/user/verify"
	WsOpLogin           = "login"
	WsOpUnsubscribe     = "unsubscribe"
	WsOpSubscribe       = "subscribe"
	PingMessage         = "ping"
	PongMessage         = "pong"
	PingIntervalSecond  = 10
	TimerIntervalSecond = 2

	/**
	ProductType
	*/
	ProductType_UMCBL  = "umcbl"  // USDT perpetual contract
	ProductType_DMCBL  = "dmcbl"  // Universal margin perpetual contract
	ProductType_CMCBL  = "cmcbl"  // USDC perpetual contract
	ProductType_SUMCBL = "sumcbl" // USDT simulation perpetual contract
	ProductType_SDMCBL = "sdmcbl" // Universal margin simulation perpetual contract
	ProductType_SCMCBL = "scmcbl" // USDC simulation perpetual contract

	/**
	CandleInterval
	*/
	CandleInterval_1m    = "1m"    // 1m (1minute)
	CandleInterval_3m    = "3m"    // 3m (3minute)
	CandleInterval_5m    = "5m"    // 5m (5minute)
	CandleInterval_15m   = "15m"   // 15m (15minute)
	CandleInterval_30m   = "30m"   // 30m (30minute)
	CandleInterval_1H    = "1H"    // 1H (1hour)
	CandleInterval_2H    = "2H"    // 2H (2hour)
	CandleInterval_4H    = "4H"    // 4H (4hour)
	CandleInterval_1Dutc = "1Dutc" // 1Dutc (UTC0 1day)

	// 6H (6hour)
	// 12H (12hour)
	// 1D (1day)
	// 3D (3day)
	// 1W (1week)
	// 1M (1month)
	// 6Hutc (UTC0 6hour)
	// 12Hutc (UTC0 12hour)
	// 1Dutc (UTC0 1day)
	// 3Dutc (UTC0 3day)
	// 1Wutc (UTC0 1 week)
	// 1Mutc (UTC0 1 month)

	/**
	Symbol
	*/
	Symbol_BTCUSDT_UMCBL = "BTCUSDT_UMCBL"

	/**
	WsChannels
	*/
	WsChannel_candle1m  = "candle1m"
	WsChannel_candle5m  = "candle5m"
	WsChannel_candle15m = "candle15m"
	WsChannel_candle30m = "andle30m"
	WsChannel_candle1H  = "candle1H"
	WsChannel_candle4H  = "candle4H"
	WsChannel_candle12H = "candle12H"
	WsChannel_candle1D  = "candle1D"
	WsChannel_candle1W  = "candle1W"

	WsChannel_trade    = "trade"    // Retrieve the recent trades data.
	WsChannel_tradeNew = "tradeNew" // Retrieve the recent trades data. The first snapshot will push 50 trade records
	WsChannel_ticker   = "ticker"   // Retrieve the latest traded price, bid price, ask price and 24-hour trading volume of the instruments. Data will be pushed every 100 ms.

	/**
	InstrumentID
	*/
	InstrumentID_BTCUSDT = "BTCUSDT"
	InstrumentID_BTCUSD  = "BTCUSD"
)
