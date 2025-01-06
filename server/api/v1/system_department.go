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

	// 校验父部门是否为空
	if req.ParentId == nil {
		response.FailedWithMessage("父部门不能为空")
		return
	}

	// 不允许 ParentId 为 2，因为 2 是预留的未分配部门
	if *req.ParentId == 2 {
		response.FailedWithMessage("父部门不能为未分配部门")
		return
	}

	// 校验名称是否合法
	if req.Name == nil || *req.Name == "" || len(*req.Name) > 30 || len(*req.Name) < 3 {
		response.FailedWithMessage("部门名称长度不合法")
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
