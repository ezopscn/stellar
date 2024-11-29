package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 当前用户的路由组
func CurrentUserAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	authRG.GET("/menu/tree", v1.GetCurrentUserSystemMenuTreeHandler) // 获取当前用户的菜单列表接口
	return authRG
}
