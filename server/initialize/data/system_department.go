package data

import (
	"stellar/common"
	"stellar/model"
)

// 部门数据
var systemDepartments = []model.SystemDepartment{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		Name:        "某科技公司",
		ParentId:    0,
		LeaderId:    1,
		Description: "公司主体，最顶层架构，其它部门都依赖于该架构",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{systemUsers[0]},
	},
	{
		BaseModel: model.BaseModel{
			Id: 2,
		},
		Name:        "研发中心",
		ParentId:    1,
		LeaderId:    1,
		Description: "研发中心，负责公司主体的研发工作",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{systemUsers[0]},
		Children: []model.SystemDepartment{
			{
				BaseModel: model.BaseModel{
					Id: 3,
				},
				Name:        "后台开发部",
				ParentId:    2,
				LeaderId:    1,
				Description: "后台开发部，负责公司主体的后台开发工作",
				Creator:     defaultCreator,
			},
			{
				BaseModel: model.BaseModel{
					Id: 4,
				},
				Name:        "前端开发部",
				ParentId:    2,
				LeaderId:    1,
				Description: "前端开发部，负责公司主体的前端开发工作",
				Creator:     defaultCreator,
			},
			{
				BaseModel: model.BaseModel{
					Id: 5,
				},
				Name:        "测试部",
				ParentId:    2,
				LeaderId:    1,
				Description: "测试部，负责公司主体的测试工作",
				Creator:     defaultCreator,
			},
			{
				BaseModel: model.BaseModel{
					Id: 6,
				},
				Name:        "运维部",
				ParentId:    2,
				LeaderId:    2,
				Description: "运维部，负责公司主体的运维工作",
				Creator:     defaultCreator,
				SystemUsers: []model.SystemUser{systemUsers[1], systemUsers[2]},
			},
		},
	},
	{
		BaseModel: model.BaseModel{
			Id: 7,
		},
		Name:        "产品中心",
		ParentId:    1,
		LeaderId:    4,
		Description: "产品中心，负责公司主体的产品规划和设计",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{systemUsers[3]},
	},
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
	common.MySQLDB.Exec("TRUNCATE TABLE system_user_department_relation")
	insertDepartmentsData(systemDepartments)
}
