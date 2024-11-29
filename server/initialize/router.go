package initialize

import (
	"stellar/common"
	"stellar/middleware"
	"stellar/router"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func Router() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.AccessLog)
	r.Use(middleware.Exception)
	r.Use(middleware.Cors)
	auth, err := middleware.JWTAuth()
	if err != nil {
		panic("JWT init error: " + err.Error())
	}
	{
		router.PublicRouter(r.Group(common.SystemApiPrefix), auth)
		router.PublicAuthRoutes(r.Group(common.SystemApiPrefix), auth)
		router.CurrentUserAuthRoutes(r.Group(common.SystemApiPrefix+"/current/user"), auth)
		router.SystemUserAuthRoutes(r.Group(common.SystemApiPrefix+"/user"), auth)
		router.SystemRoleAuthRoutes(r.Group(common.SystemApiPrefix+"/role"), auth)
		router.SystemMenuAuthRoutes(r.Group(common.SystemApiPrefix+"/menu"), auth)
	}
	return r
}
