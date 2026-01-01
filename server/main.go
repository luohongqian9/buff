package main

import (
	"server/core"
	"server/flags"
	"server/global"
	"server/initialize"
)

func main() {
	global.Config = core.InitConfig()
	global.Log = core.InitLogger()
	global.DB = initialize.InitGorm()
	global.Redis = initialize.ConnectRedis()
	global.ESClient = initialize.ConnectEs()
	initialize.InitOther()
	defer global.Redis.Close()

	flags.InitFlag()

	initialize.InitCron()

	core.RunServer()
}
