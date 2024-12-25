package data

import (
	"stellar/common"
	"stellar/model"
)

// 用户相关规则
var systemUserCasbinRules = []model.SystemCasbinRule{
	// 用户列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/user/list", Method: "GET"},
	// // 用户状态修改
	// {PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/user/status-modify", Method: "POST"},
	// // 用户添加
	// {PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/user/add", Method: "POST"},
	// // 用户导入
	// {PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/user/import", Method: "POST"},
}

// 角色相关规则
var systemRoleCasbinRules = []model.SystemCasbinRule{
	// 角色列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/role/list", Method: "GET"},
	// 角色API列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/role/api/list", Method: "GET"},
	{PType: "p", Keyword: systemRoles[3].Keyword, Path: "/system/role/api/list", Method: "GET"},
}

// 菜单相关规则
var systemMenuCasbinRules = []model.SystemCasbinRule{
	// 菜单树
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/menu/tree", Method: "GET"},
}

// 部门相关规则
var systemDepartmentCasbinRules = []model.SystemCasbinRule{
	// 部门列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/department/list", Method: "GET"},
}

// 职位相关规则
var systemJobPositionCasbinRules = []model.SystemCasbinRule{
	// 职位列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/job-position/list", Method: "GET"},
}

// API相关规则
var systemApiCasbinRules = []model.SystemCasbinRule{
	// API列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/api/list", Method: "GET"},
}

// API分类相关规则
var systemApiCategoryCasbinRules = []model.SystemCasbinRule{
	// API分类列表
	{PType: "p", Keyword: systemRoles[2].Keyword, Path: "/system/api-category/list", Method: "GET"},
}

// 初始化系统CASBIN规则数据
func InitSystemCasbinRuleData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_casbin_rule")
	common.MySQLDB.Create(&systemUserCasbinRules)
	common.MySQLDB.Create(&systemRoleCasbinRules)
	common.MySQLDB.Create(&systemMenuCasbinRules)
	common.MySQLDB.Create(&systemDepartmentCasbinRules)
	common.MySQLDB.Create(&systemJobPositionCasbinRules)
	common.MySQLDB.Create(&systemApiCasbinRules)
	common.MySQLDB.Create(&systemApiCategoryCasbinRules)
}
