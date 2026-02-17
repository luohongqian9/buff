package initialize

import (
	"github.com/gin-gonic/gin"
	"server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	Router := gin.Default()

	var store = cookie.NewStore([]byte(global.Config.System.SessionsSecret))
	Router.Use(sessions.Sessions("session",store))

	routerGroup := router.RouterGroupApp
	
	return Router
}
