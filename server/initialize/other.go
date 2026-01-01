package initialize

import (
	"os"
	"server/global"
	"server/utils"

	"github.com/songzhibin97/gkit/cache/local_cache"
	"go.uber.org/zap"
)

func InitOther() {
	refreshTokenExpireTime, err := utils.ParseDuration(global.Config.Jwt.RefreshTokenExpiryTime)
	if err != nil {
		global.Log.Error("解析刷新token过期时间失败", zap.Error(err))
		os.Exit(1)
	}
	_, err = utils.ParseDuration(global.Config.Jwt.AccessTokenExpiryTime)
	if err != nil {
		global.Log.Error("解析访问token过期时间失败", zap.Error(err))
		os.Exit(1)
	}

	global.BlackCache = local_cache.NewCache(local_cache.SetDefaultExpire(refreshTokenExpireTime))
}
