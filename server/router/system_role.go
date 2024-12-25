package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 角色路由，不需要权限校验
func SystemRoleAuthRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/api/list", v1.SystemRoleApiListHandler)
	return rg
}

// 角色路由，需要权限校验
func SystemRoleAuthAndPermissionRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/list", v1.SystemRoleListHandler)
	return rg
}
