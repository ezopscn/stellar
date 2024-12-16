package router

import (
	v1 "stellar/api/v1"

	"github.com/gin-gonic/gin"
)

// 当前用户的路由组
func CurrentSystemUserAuthRoutes(rg *gin.RouterGroup) gin.IRoutes {
	rg.GET("/menu/tree", v1.GetCurrentSystemUserSystemMenuTreeHandler) // 获取当前用户的菜单列表接口
	return rg
}
