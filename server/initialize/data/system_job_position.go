package data

import (
	"stellar/common"
	"stellar/model"
)

// 职位初始化
var systemJobPositions = []model.SystemJobPosition{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		Name:        "首席执行官（CEO）",
		Description: "首席执行官，负责公司主体的执行工作",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{
			systemUsers[0],
		},
	},
	{
		BaseModel: model.BaseModel{
			Id: 2,
		},
		Name:        "研发总监（CTO）",
		Description: "研发总监，负责公司主体的技术工作",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{
			systemUsers[0],
		},
	},
	{
		BaseModel: model.BaseModel{
			Id: 3,
		},
		Name:        "财务总监（CFO）",
		Description: "首席财务官，负责公司主体的财务工作",
		Creator:     defaultCreator,
	},
	{
		BaseModel: model.BaseModel{
			Id: 4,
		},
		Name:        "产品总监",
		Description: "产品总监，负责公司主体的产品工作",
		Creator:     defaultCreator,
	},
	{
		BaseModel: model.BaseModel{
			Id: 5,
		},
		Name:        "运维总监",
		Description: "运维总监，负责公司主体的运维工作",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{
			systemUsers[1],
		},
	},
	{
		BaseModel: model.BaseModel{
			Id: 6,
		},
		Name:        "高级运维工程师",
		Description: "高级运维工程师，负责公司主体的运维工作",
		Creator:     defaultCreator,
	},
	{
		BaseModel: model.BaseModel{
			Id: 7,
		},
		Name:        "运维开发工程师",
		Description: "运维开发工程师，负责公司主体的运维开发工作",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{
			systemUsers[2],
		},
	},
	{
		BaseModel: model.BaseModel{
			Id: 8,
		},
		Name:        "系统运维工程师",
		Description: "系统运维工程师，负责公司主体的系统运维工作",
		Creator:     defaultCreator,
	},
	{
		BaseModel: model.BaseModel{
			Id: 9,
		},
		Name:        "业务运维工程师",
		Description: "业务运维工程师，负责公司主体的业务运维工作",
		Creator:     defaultCreator,
	},
	{
		BaseModel: model.BaseModel{
			Id: 10,
		},
		Name:        "访客",
		Description: "访客，只读用户",
		Creator:     defaultCreator,
		SystemUsers: []model.SystemUser{
			systemUsers[3],
		},
	},
}

// 职位初始化
func InitJobPositionData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_job_position")
	common.MySQLDB.Create(&systemJobPositions)
}
