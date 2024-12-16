package service

import (
	"stellar/common"
	"stellar/model"
)

// 获取角色列表
func GetSystemRoleListService() (roles []model.SystemRole, err error) {
	err = common.MySQLDB.Find(&roles).Error
	return
}
