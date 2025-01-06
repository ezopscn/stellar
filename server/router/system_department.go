package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 系统部门路由，不需要权限校验
func SystemDepartmentAuthRoutes(rg *gin.RouterGroup) gin.IRoutes {
	return rg
}

// 系统部门路由，需要权限校验
func SystemDepartmentAuthAndPermissionRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/list", v1.SystemDepartmentListHandler)
	rg.GET("/detail", v1.SystemDepartmentDetailHandler)
	rg.POST("/add", v1.SystemDepartmentAddHandler)
	return rg
}
