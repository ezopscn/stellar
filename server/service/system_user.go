package service

import (
	"errors"
	"fmt"
	"stellar/common"
	"stellar/dto"
	"stellar/model"
	"stellar/pkg/gedis"
	"stellar/pkg/trans"
	"stellar/pkg/utils"

	"github.com/gin-gonic/gin"
)

// 获取用户列表
func GetSystemUserListService(ctx *gin.Context) (users []model.SystemUser, pagination dto.Pagination, err error) {
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
	if filter.SystemDepartment != nil {
		query = query.Joins("JOIN system_user_department_relation ON system_user_department_relation.system_user_id = system_user.id").
			Where("system_user_department_relation.system_department_id = ?", *filter.SystemDepartment)
	}

	// 岗位
	if filter.SystemJobPosition != nil {
		query = query.Joins("JOIN system_user_job_position_relation ON system_user_job_position_relation.system_user_id = system_user.id").
			Where("system_user_job_position_relation.system_job_position_id = ?", *filter.SystemJobPosition)
	}

	// 角色
	if filter.SystemRole != nil {
		query = query.Where("systemRoleId = ?", *filter.SystemRole)
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

// 修改用户状态
func ModifySystemUserStatusService(ctx *gin.Context, ids []uint, operate string) error {
	// 默认 ID 为 1 的用户不能修改状态，属于系统预设用户
	if utils.IsUintInSlice(1, ids) {
		return errors.New("系统预设用户不能修改状态")
	}

	// 判断操作类型
	if operate == "disable" {
		// 禁用用户，先查询所有用户的用户名，清除用户登录的 token，然后更新用户状态
		var usernames []string
		err := common.MySQLDB.Model(&model.SystemUser{}).Select("username").Where("id IN (?)", ids).Find(&usernames).Error
		if err != nil {
			return errors.New("查询需要禁用的用户信息失败")
		}
		conn := gedis.NewRedisConnection()
		for _, username := range usernames {
			key := fmt.Sprintf("%s:%s", common.RKP.LoginToken, username)
			conn.Del(key)
		}
		return common.MySQLDB.Model(&model.SystemUser{}).Where("id IN (?)", ids).Update("status", trans.Uint(0)).Error
	} else if operate == "enable" {
		return common.MySQLDB.Model(&model.SystemUser{}).Where("id IN (?)", ids).Update("status", trans.Uint(1)).Error
	} else {
		return errors.New("不支持的操作类型")
	}
}
