package v1

import (
	"stellar/common"
	"stellar/model"
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

	// API 列表
	var apis []model.SystemApi
	// 判断角色是不是管理员
	if utils.IsStringInSlice(roleKeyword, common.SystemRoleAdminList) {
		// 查询所有 API
		err = common.MySQLDB.Where("needPermission = ?", 1).Find(&apis).Error
	} else {
		// 查询当前角色的 API 列表，只需要查询鉴权的
		var role model.SystemRole
		err = common.MySQLDB.Where("keyword = ?", roleKeyword).Preload("SystemApis").First(&role).Error
		apis = role.SystemApis
	}
	if err != nil {
		response.FailedWithMessage("获取当前角色的API列表失败")
		return
	}
	// 只需要 API 列表即可
	var apiList []string
	for _, api := range apis {
		apiList = append(apiList, api.Path)
	}
	response.SuccessWithData(map[string]interface{}{
		"list": apiList,
	})
}
