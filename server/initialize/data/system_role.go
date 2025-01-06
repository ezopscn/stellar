package data

import (
	"stellar/common"
	"stellar/model"
)

// 角色初始化
var systemRoles = []model.SystemRole{
	{
		Id:          1,
		Name:        "超级管理员",
		Keyword:     "SuperAdministrator",
		Description: "系统最高权限管理角色，不推荐用户直接加入该角色",
	},
	{
		Id:          2,
		Name:        "管理员",
		Keyword:     "Administrator",
		Description: "系统管理角色，拥有用户的管理权限，但是不具备系统的管理权限",
	},
	{
		Id:          3,
		Name:        "运维",
		Keyword:     "DevOps",
		Description: "监控配置角色，拥有监控指标配置的权限",
	},
	{
		Id:          4,
		Name:        "访客",
		Keyword:     "Guest",
		Description: "只读普通用户",
	},
}

// 角色初始化
func InitRoleData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_role")
	common.MySQLDB.Exec("TRUNCATE TABLE system_role_api_relation")
	common.MySQLDB.Exec("TRUNCATE TABLE system_role_menu_relation")
	common.MySQLDB.Create(&systemRoles)
}
