package v1

import (
	"stellar/pkg/response"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 系统部门列表
func SystemDepartmentListHandler(ctx *gin.Context) {
	departments, err := service.GetSystemDepartmentListService()
	if err != nil {
		response.FailedWithMessage("获取部门列表失败")
		return
	}
	response.SuccessWithData(departments)
}
