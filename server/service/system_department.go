package service

import (
	"stellar/common"
	"stellar/model"
)

// 获取部门列表
func GetSystemDepartmentListService() (departments []model.SystemDepartment, err error) {
	err = common.MySQLDB.Find(&departments).Error
	return
}

// 获取部门详情
func GetSystemDepartmentDetailService(id string) (department model.SystemDepartment, err error) {
	err = common.MySQLDB.Where("id = ?", id).First(&department).Error
	return
}
