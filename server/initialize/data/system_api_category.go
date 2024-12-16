package data

import (
	"stellar/common"
	"stellar/model"
)

// 系统API分类数据
var SystemApiCategorys = []model.SystemApiCategory{
	{BaseModel: model.BaseModel{Id: 1}, Name: "部门管理", Description: "部门管理类型，用于部门管理相关的API"},
	{BaseModel: model.BaseModel{Id: 2}, Name: "岗位管理", Description: "岗位管理类型，用于岗位管理相关的API"},
	{BaseModel: model.BaseModel{Id: 3}, Name: "用户管理", Description: "用户管理类型，用于用户管理相关的API"},
	{BaseModel: model.BaseModel{Id: 4}, Name: "角色管理", Description: "角色管理类型，用于角色管理相关的API"},
	{BaseModel: model.BaseModel{Id: 5}, Name: "菜单管理", Description: "菜单管理类型，用于菜单管理相关的API"},
	{BaseModel: model.BaseModel{Id: 6}, Name: "API分类管理", Description: "API分类管理类型，用于API分类管理相关的API"},
	{BaseModel: model.BaseModel{Id: 7}, Name: "接口管理", Description: "接口管理类型，用于接口管理相关的API"},
}

// 初始化系统API分类数据
func InitSystemApiCategoryData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_api_category")
	common.MySQLDB.Create(&SystemApiCategorys)
}
