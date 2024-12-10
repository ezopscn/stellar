package service

import (
	"errors"
	"fmt"
	"stellar/common"
	"stellar/dto"
	"stellar/model"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func GetUserListService(ctx *gin.Context) (users []model.SystemUser, pagination dto.Pagination, err error) {
	// 获取筛选条件
	filter := dto.SystemUserFilterRequest{}
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		return nil, pagination, errors.New("获取用户筛选条件失败")
	}

	// 初始化查询条件
	query := common.MySQLDB.Model(&model.SystemUser{})

	// 用户名
	if filter.Username != nil {
		query = query.Where("username LIKE ?", "%"+*filter.Username+"%")
	}

	// 姓名
	if filter.Name != nil {
		query = query.Where("cnName LIKE ? OR enName LIKE ?", "%"+*filter.Name+"%", "%"+*filter.Name+"%")
	}

	// 邮箱
	if filter.Email != nil {
		query = query.Where("email LIKE ?", "%"+*filter.Email+"%")
	}

	// 手机号
	if filter.Phone != nil {
		query = query.Where("phone LIKE ?", "%"+*filter.Phone+"%")
	}

	// 状态
	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	// 性别
	if filter.Gender != nil {
		query = query.Where("gender = ?", *filter.Gender)
	}

	// 部门
	if filter.Department != nil {
		query = query.Joins("JOIN system_user_department_relation ON system_user_department_relation.system_user_id = system_user.id").
			Where("system_user_department_relation.system_department_id = ?", *filter.Department)
	}

	// 岗位
	if filter.JobPosition != nil {
		query = query.Joins("JOIN system_user_job_position_relation ON system_user_job_position_relation.system_user_id = system_user.id").
			Where("system_user_job_position_relation.system_job_position_id = ?", *filter.JobPosition)
	}

	// 角色
	if filter.Role != nil {
		query = query.Where("systemRoleId = ?", *filter.Role)
	}

	// 统计记录数量
	var total int64
	query.Count(&total)
	pagination.Total = total

	// 不传递分页，则默认要分页
	if filter.IsPagination != nil && *filter.IsPagination {
		pagination.IsPagination = true
		if filter.PageNumber != nil {
			pagination.PageNumber = *filter.PageNumber
		}
		if filter.PageSize != nil {
			pagination.PageSize = *filter.PageSize
		}
		fmt.Println("pagination: ", pagination)
		limit, offset := pagination.GetPaginationLimitAndOffset()
		query = query.Limit(limit).Offset(offset)
	}

	// 加入查询条件
	err = query.Preload("SystemRole").Preload("SystemDepartments").Preload("SystemJobPositions").Find(&users).Error
	return
}
