package router

import (
	v1 "stellar/api/v1"

	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 开放路由
func PublicRouter(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/health", v1.SystemHealthCheckHandler)      // 健康检查
	rg.GET("/information", v1.SystemInformationHandler) // 系统信息
	rg.GET("/version", v1.SystemVersionHandler)         // 系统版本
	rg.POST("/login", auth.LoginHandler)                // 用户登录
	return rg
}

// 登录路由组
func PublicAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	rg.GET("/node/information", v1.NodeInformationHandler)     // 节点信息
	rg.GET("/token/verification", v1.TokenVerificationHandler) // Token 校验
	rg.GET("/logout", auth.LogoutHandler)                      // 用户注销
	return rg
}
