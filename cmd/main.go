package main

import (
	"github.com/yasseldg/bitget/cmd/restV2"

	"github.com/yasseldg/go-simple/logs/sLog"
	"github.com/yasseldg/go-simple/types/sTime"
)

func main() {
	clean := sLog.SetByName(sLog.Zap, sLog.LevelInfo, "")
	defer clean()

	sLog.Info("Starting...")

	sTime.TimeControl(restV2.Private, "Start")
}
