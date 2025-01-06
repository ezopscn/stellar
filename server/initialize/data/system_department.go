package data

import (
	"stellar/common"
	"stellar/model"
)

// 部门数据
var systemDepartments = []model.SystemDepartment{
	{BaseModel: model.BaseModel{Id: 1}, Name: "某科技公司", ParentId: 0, Creator: defaultCreator, SystemUsers: []model.SystemUser{systemUsers[0]}},
	{BaseModel: model.BaseModel{Id: 2}, Name: "未分配部门", ParentId: 1, Creator: defaultCreator},
	{BaseModel: model.BaseModel{Id: 3}, Name: "研发中心", ParentId: 1, Creator: defaultCreator, SystemUsers: []model.SystemUser{systemUsers[0]}, Children: []model.SystemDepartment{
		{BaseModel: model.BaseModel{Id: 4}, Name: "后台开发部", ParentId: 3, Creator: defaultCreator},
		{BaseModel: model.BaseModel{Id: 5}, Name: "前端开发部", ParentId: 3, Creator: defaultCreator},
		{BaseModel: model.BaseModel{Id: 6}, Name: "测试部", ParentId: 3, Creator: defaultCreator},
		{BaseModel: model.BaseModel{Id: 7}, Name: "运维部", ParentId: 3, Creator: defaultCreator, SystemUsers: []model.SystemUser{systemUsers[1], systemUsers[2]}},
	}},
	{BaseModel: model.BaseModel{Id: 8}, Name: "产品中心", ParentId: 1, Creator: defaultCreator, SystemUsers: []model.SystemUser{systemUsers[3]}},
}

// 递归插入数据方法
func insertDepartmentsData(departments []model.SystemDepartment) {
	for _, department := range departments {
		common.MySQLDB.Create(&department)
		if len(department.Children) > 0 {
			insertDepartmentsData(department.Children)
		}
	}
}

// 部门初始化
func InitDepartmentData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_department")
	insertDepartmentsData(systemDepartments)
}
