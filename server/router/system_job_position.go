package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 岗位路由，不需要权限校验
func SystemJobPositionAuthRoutes(rg *gin.RouterGroup) gin.IRoutes {
	return rg
}

// 岗位路由，需要权限校验
func SystemJobPositionAuthAndPermissionRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/list", v1.SystemJobPositionListHandler)
	return rg
}
