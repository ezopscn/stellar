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

// 系统部门详情
func SystemDepartmentDetailHandler(ctx *gin.Context) {
	id := ctx.Query("id")
	department, err := service.GetSystemDepartmentDetailService(id)
	if err != nil {
		response.FailedWithMessage("获取部门详情失败")
		return
	}
	response.SuccessWithData(department)
}
