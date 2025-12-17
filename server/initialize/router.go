package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()
	return Router
}
