package v1

import (
	"stellar/common"
	"stellar/dto"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/utils"
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

// 新增部门
func SystemDepartmentAddHandler(ctx *gin.Context) {
	// 获取当前用户信息
	creator := utils.GenerateCreator(ctx)
	if creator == "" {
		response.FailedWithMessage("生成创建者失败")
		return
	}

	// 获取请求参数
	var req dto.SystemDepartmentAddRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage("请求参数错误")
		return
	}

	// 校验参数
	if err := req.Validate(); err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 新增部门
	if err := common.MySQLDB.Create(&model.SystemDepartment{
		ParentId: *req.ParentId,
		Name:     *req.Name,
		Creator:  creator,
	}).Error; err != nil {
		response.FailedWithMessage("新增部门失败：" + err.Error())
		return
	}
	response.Success()
}

// 修改部门
func SystemDepartmentUpdateHandler(ctx *gin.Context) {
	// 获取请求参数
	var req dto.SystemDepartmentUpdateRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage("请求参数错误")
		return
	}

	// 校验参数合法性
}
