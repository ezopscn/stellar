package v1

import (
	"stellar/pkg/response"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 部门列表
func DepartmentListHandler(ctx *gin.Context) {
	departments, err := service.GetDepartmentListService()
	if err != nil {
		response.FailedWithMessage("获取部门列表失败")
		return
	}
	response.SuccessWithData(departments)
}
