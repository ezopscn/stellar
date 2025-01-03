package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 用户路由，不需要权限校验
func SystemUserAuthRoutes(rg *gin.RouterGroup) gin.IRoutes {
	return rg
}

// 用户路由，需要权限校验
func SystemUserAuthAndPermissionRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/list", v1.SystemUserListHandler)                              // 用户列表接口
	rg.POST("/add", v1.SystemUserAddHandler)                               // 添加用户接口
	rg.POST("/multi-add", v1.SystemUserMultiAddHandler)                    // 批量添加用户接口
	rg.POST("/modify-status", v1.SystemUserModifyStatusHandler)            // 修改用户状态
	rg.POST("/multi-modify-status", v1.SystemUserMultiModifyStatusHandler) // 批量修改用户状态
	return rg
}
