package market

type CandlesResp [][]string

// [
//   [
//     "1673036160000",		Timestamp in milliseconds		0
//     "16995",				Opening price					1
//     "17034.5",			Highest price					2
//     "16995",				Lowest price					3
//     "17010.5",			Closing price   				4  ( value of the latest candle stick might change, please try subscribe the websocket candlestick channel for the updates )
//     "1523.183",			Base currency trading volume	5
//     "25920426.944"		Quote currency trading volume	6
//   ],
//   [
//     "1673036220000",
//     "17010.5",
//     "17022.5",
//     "17006.5",
//     "17008",
//     "556.984",
//     "9475143.812"
//   ],
//   [
//     "1673036280000",
//     "17008",
//     "17008.5",
//     "17000",
//     "17003",
//     "465.066",
//     "7907436.9925"
//   ]
// ]
