package v1

import (
	"stellar/common"
	"stellar/pkg/response"
	"stellar/pkg/utils"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 用户列表接口
func UserListHandler(ctx *gin.Context) {
	// 获取当前用户的角色，如果是管理员以上的级别，则可以看到隐藏的手机号
	systemRoleKeyword, err := utils.ExtractStringResultFromContext(ctx, "systemRoleKeyword")
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 获取用户列表
	users, err := service.GetUserListService(ctx)
	if err != nil {
		response.FailedWithMessage("获取用户列表失败")
		return
	}

	// 判断角色是否在管理员列表，如果不是管理员，则隐藏设置了隐藏标识的手机号
	if !utils.IsStringInSlice(systemRoleKeyword, common.SystemRoleAdminList) {
		for i := range users {
			if *users[i].HidePhone == 1 {
				users[i].Phone = utils.HidePhoneNumber(users[i].Phone)
			}
		}
	}

	response.SuccessWithData(users)
}
