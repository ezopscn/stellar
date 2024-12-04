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
		router.PublicRouter(r.Group(common.SystemApiPrefix), auth)                          // 公共路由（不需要认证）
		router.PublicAuthRoutes(r.Group(common.SystemApiPrefix), auth)                      // 公共路由（需要认证）
		router.CurrentUserAuthRoutes(r.Group(common.SystemApiPrefix+"/current/user"), auth) // 当前用户路由（需要认证）
		// 系统路由组
		{
			router.SystemUserAuthRoutes(r.Group(common.SystemApiPrefix+"/system/user"), auth) // 用户路由（需要认证）
			router.SystemRoleAuthRoutes(r.Group(common.SystemApiPrefix+"/system/role"), auth) // 角色路由（需要认证）
			router.SystemMenuAuthRoutes(r.Group(common.SystemApiPrefix+"/system/menu"), auth) // 菜单路由（需要认证）
		}
	}
	return r
}
