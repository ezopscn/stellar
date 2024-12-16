package data

import (
	"stellar/common"
	"stellar/model"
)

// 系统CASBIN规则数据
var systemCasbinRules = []model.SystemCasbinRule{
	{PType: "p", Keyword: "DevOps", Path: "/system/department/list", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/jobPosition/list", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/user/list", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/role/list", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/menu/list", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/menu/tree", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/apiCategory/list", Method: "GET"},
	{PType: "p", Keyword: "DevOps", Path: "/system/api/list", Method: "GET"},
}

// 初始化系统CASBIN规则数据
func InitSystemCasbinRuleData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_casbin_rule")
	common.MySQLDB.Create(&systemCasbinRules)
}
