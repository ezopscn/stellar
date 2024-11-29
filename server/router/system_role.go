package router

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"github.com/gin-gonic/gin"
)

// 角色路由
func SystemRoleAuthRoutes(rg *gin.RouterGroup, auth *jwt.GinJWTMiddleware) gin.IRoutes {
	authRG := rg.Use(auth.MiddlewareFunc())
	return authRG
}
