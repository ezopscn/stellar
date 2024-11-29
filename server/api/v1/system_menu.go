package v1

import (
	"stellar/pkg/response"
	"stellar/pkg/utils"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 菜单树列表接口
func GetSystemMenuTreeHandler(ctx *gin.Context) {
	tree, err := service.GenerateSystemMenuTreeByRoleIdService(0)
	if err != nil {
		response.FailedWithMessage("获取菜单树数据失败")
		return
	}
	response.SuccessWithData(tree)
}

// 获取当前用户的菜单树接口
func GetCurrentUserSystemMenuTreeHandler(ctx *gin.Context) {
	roleId, err := utils.ExtractUintResultFromContext(ctx, "systemRoleId")
	if err != nil {
		response.FailedWithMessage("获取当前用户角色ID失败")
		return
	}
	tree, err := service.GenerateSystemMenuTreeByRoleIdService(roleId)
	if err != nil {
		response.FailedWithMessage("获取当前用户的菜单列表失败")
		return
	}
	response.SuccessWithData(tree)
}
