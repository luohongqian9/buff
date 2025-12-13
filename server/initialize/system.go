package initialize

import (
	"fmt"
	"server/config"
	"server/global"
	"server/model/appTypes"
	"strings"
)

func InitSystem() {
	sys := config.Config.System

	global.Addr = initialize.BuildAddr(sys)
	global.StorageType = initialize.ParseStorageType(sys)
}

func BuildAddr(s config.System) string {
	return fmt.Sprintf("%s:%d", s.Host, s.Port)
}

func ParseStorageType(s config.System) appTypes.Storage {
	switch strings.ToLower(s.OssType) {
	case "local":
		return appTypes.Local
	case "qiniu":
		return appTypes.Qiniu
	default:
		return appTypes.Local
	}

}
