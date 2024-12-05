package service

import (
	"stellar/common"
	"stellar/model"
)

// 获取用户列表
func GetUserListService() (users []model.SystemUser, err error) {
	// 加入查询条件
	err = common.MySQLDB.Preload("SystemRole").Preload("SystemDepartments").Preload("SystemJobPositions").Find(&users).Error
	return
}
