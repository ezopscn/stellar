package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 菜单路由，不需要权限校验
func SystemMenuAuthRoutes(rg *gin.RouterGroup) gin.IRoutes {
	return rg
}

// 菜单路由，需要权限校验
func SystemMenuAuthAndPermissionRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/tree", v1.GetSystemMenuTreeHandler) // 系统菜单树接口
	return rg
}
