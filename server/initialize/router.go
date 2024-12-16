package initialize

import (
	"stellar/common"
	"stellar/middleware"
	"stellar/router"

	"github.com/gin-gonic/gin"
)

// 路由初始化
func Router() *gin.Engine {
	// gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.Use(middleware.AccessLog)
	r.Use(middleware.Exception)
	r.Use(middleware.Cors)
	auth, err := middleware.JWTAuth()
	if err != nil {
		panic("JWT init error: " + err.Error())
	}
	// 公共路由，不需要认证
	{
		publicRG := r.Group(common.SystemApiPrefix)
		router.PublicRouter(publicRG, auth)
	}
	// 需要认证，但是无需权限校验的路由
	{
		authRG := r.Group(common.SystemApiPrefix)
		authRG.Use(auth.MiddlewareFunc())
		// 公共路由
		{
			router.PublicAuthRoutes(authRG, auth)                             // 公共路由
			router.CurrentSystemUserAuthRoutes(authRG.Group("/current/user")) // 当前用户路由
		}
		// 系统路由组
		{
			router.SystemUserAuthRoutes(authRG.Group("/system/user"))               // 用户路由
			router.SystemRoleAuthRoutes(authRG.Group("/system/role"))               // 角色路由
			router.SystemMenuAuthRoutes(authRG.Group("/system/menu"))               // 菜单路由
			router.SystemJobPositionAuthRoutes(authRG.Group("/system/jobPosition")) // 岗位路由
			router.SystemDepartmentAuthRoutes(authRG.Group("/system/department"))   // 部门路由
		}
	}
	// 需要认证，需要权限校验的路由
	{
		authPermissionRG := r.Group(common.SystemApiPrefix)
		authPermissionRG.Use(auth.MiddlewareFunc())
		authPermissionRG.Use(middleware.Casbin)
		// 系统路由组
		{
			router.SystemUserAuthAndPermissionRoutes(authPermissionRG.Group("/system/user"))               // 用户路由
			router.SystemRoleAuthAndPermissionRoutes(authPermissionRG.Group("/system/role"))               // 角色路由
			router.SystemMenuAuthAndPermissionRoutes(authPermissionRG.Group("/system/menu"))               // 菜单路由
			router.SystemJobPositionAuthAndPermissionRoutes(authPermissionRG.Group("/system/jobPosition")) // 岗位路由
			router.SystemDepartmentAuthAndPermissionRoutes(authPermissionRG.Group("/system/department"))   // 部门路由
		}
	}
	return r
}
