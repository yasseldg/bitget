package restV2

import (
	"github.com/yasseldg/bitget"
	"github.com/yasseldg/bitget/cmd/config"
)

func Orders() {
	println()
	rest := bitget.NewRest(config.FullApi())

	getPendingOrders(rest)

	println("\n")

	rest1 := bitget.NewRest(config.PereApi())

	getPendingOrders(rest1)
}
