package v1

import (
	"stellar/pkg/response"
	"stellar/pkg/utils"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 获取角色列表
func SystemRoleListHandler(ctx *gin.Context) {
	roles, err := service.GetSystemRoleListService()
	if err != nil {
		response.FailedWithMessage("获取角色列表失败")
		return
	}
	response.SuccessWithData(roles)
}

// 获取当前角色的API列表
func SystemRoleApiListHandler(ctx *gin.Context) {
	// 获取当前角色
	roleKeyword, err := utils.ExtractStringResultFromContext(ctx, "systemRoleKeyword")
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 获取当前角色的API列表
	apiList, err := service.GetSystemRoleApiListService(roleKeyword)
	if err != nil {
		response.FailedWithMessage("获取当前角色的API列表失败")
		return
	}
	response.SuccessWithData(map[string]interface{}{
		"list": apiList,
	})
}
