package service

import (
	"stellar/common"
	"stellar/model"
)

// 获取部门列表
func GetDepartmentListService() (departments []model.SystemDepartment, err error) {
	err = common.MySQLDB.Find(&departments).Error
	return
}
