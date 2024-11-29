package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 菜单路由
func SystemMenuAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	authRG.GET("/tree", v1.GetSystemMenuTreeHandler) // 系统菜单树接口
	return authRG
}
