package v1

import (
	"stellar/pkg/response"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 获取角色列表
func RoleListHandler(ctx *gin.Context) {
	roles, err := service.GetRoleListService()
	if err != nil {
		response.FailedWithMessage("获取角色列表失败")
		return
	}
	response.SuccessWithData(roles)
}
