package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 用户路由
func SystemUserAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	authRG.GET("/list", v1.SystemUserListHandler) // 用户列表接口
	authRG.POST("/add", v1.SystemUserAddHandler)  // 添加用户接口
	return authRG
}
