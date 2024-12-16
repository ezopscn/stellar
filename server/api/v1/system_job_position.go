package v1

import (
	"stellar/pkg/response"
	"stellar/service"

	"github.com/gin-gonic/gin"
)

// 系统岗位列表
func SystemJobPositionListHandler(ctx *gin.Context) {
	jobPositions, err := service.GetSystemJobPositionListService()
	if err != nil {
		response.FailedWithMessage("获取岗位列表失败")
		return
	}
	response.SuccessWithData(jobPositions)
}
