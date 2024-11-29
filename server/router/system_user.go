package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 用户路由
func SystemUserAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	return authRG
}
