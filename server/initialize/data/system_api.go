package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
)

// 开放 API，无需登录
var publicApis = []model.SystemApi{
	{Name: "登录", Path: "/login", Method: "POST", Description: "登录接口", NeedPermission: trans.Uint(0), CategoryId: 1},
	{Name: "健康检查", Path: "/health", Method: "GET", Description: "健康检查接口", NeedPermission: trans.Uint(0), CategoryId: 1},
	{Name: "系统信息", Path: "/info", Method: "GET", Description: "系统信息接口", NeedPermission: trans.Uint(0), CategoryId: 1},
	{Name: "系统版本", Path: "/version", Method: "GET", Description: "系统版本接口", NeedPermission: trans.Uint(0), CategoryId: 1},
}

// 开放 API，需要登录
var publicAuthApis = []model.SystemApi{
	{Name: "Token校验", Path: "/token/verification", Method: "GET", Description: "Token校验接口", NeedPermission: trans.Uint(0), CategoryId: 1},
	{Name: "用户注销", Path: "/logout", Method: "GET", Description: "用户注销接口", NeedPermission: trans.Uint(0), CategoryId: 1},
}

// 当前用户 API，无需鉴权
var currentUserAuthApis = []model.SystemApi{
	{Name: "当前用户菜单树", Path: "/current/user/menu/tree", Method: "GET", Description: "当前用户菜单树接口", NeedPermission: trans.Uint(0), CategoryId: 2},
}

// 用户 API，无需鉴权
var systemUserAuthApis = []model.SystemApi{}

// 用户 API，需要鉴权
var systemUserAuthAndPermissionApis = []model.SystemApi{
	{Name: "用户列表", Path: "/system/user/list", Method: "GET", Description: "用户列表接口", NeedPermission: trans.Uint(1), CategoryId: 3, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "用户添加", Path: "/system/user/add", Method: "POST", Description: "用户添加接口", NeedPermission: trans.Uint(1), CategoryId: 3},
	{Name: "用户批量添加", Path: "/system/user/muti-add", Method: "POST", Description: "用户批量添加接口", NeedPermission: trans.Uint(1), CategoryId: 3},
	{Name: "用户状态修改", Path: "/system/user/status-modify", Method: "POST", Description: "用户状态修改接口", NeedPermission: trans.Uint(1), CategoryId: 3},
	{Name: "用户批量状态修改", Path: "/system/user/status-muti-modify", Method: "POST", Description: "用户批量状态修改接口", NeedPermission: trans.Uint(1), CategoryId: 3},
}

// 角色 API，无需鉴权
var systemRoleAuthApis = []model.SystemApi{}

// 角色 API，需要鉴权
var systemRoleAuthAndPermissionApis = []model.SystemApi{
	{Name: "角色列表", Path: "/system/role/list", Method: "GET", Description: "角色列表接口", NeedPermission: trans.Uint(1), CategoryId: 4, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "角色API列表", Path: "/system/role/api/list", Method: "GET", Description: "角色API列表接口", NeedPermission: trans.Uint(1), CategoryId: 4, SystemRoles: []model.SystemRole{systemRoles[2], systemRoles[3]}},
}

// 菜单 API，无需鉴权
var systemMenuAuthApis = []model.SystemApi{}

// 菜单 API，需要鉴权
var systemMenuAuthAndPermissionApis = []model.SystemApi{}

// 部门 API，无需鉴权
var systemDepartmentAuthApis = []model.SystemApi{}

// 部门 API，需要鉴权
var systemDepartmentAuthAndPermissionApis = []model.SystemApi{
	{Name: "部门列表", Path: "/system/department/list", Method: "GET", Description: "部门列表接口", NeedPermission: trans.Uint(1), CategoryId: 6, SystemRoles: []model.SystemRole{systemRoles[2]}},
}

// 职位 API，无需鉴权
var systemJobPositionAuthApis = []model.SystemApi{}

// 职位 API，需要鉴权
var systemJobPositionAuthAndPermissionApis = []model.SystemApi{
	{Name: "职位列表", Path: "/system/job-position/list", Method: "GET", Description: "职位列表接口", NeedPermission: trans.Uint(1), CategoryId: 7, SystemRoles: []model.SystemRole{systemRoles[2]}},
}

// 初始化系统API数据
func InitSystemApiData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_api")
	common.MySQLDB.Create(&publicApis)
	common.MySQLDB.Create(&publicAuthApis)
	common.MySQLDB.Create(&currentUserAuthApis)
	common.MySQLDB.Create(&systemUserAuthApis)
	common.MySQLDB.Create(&systemUserAuthAndPermissionApis)
	common.MySQLDB.Create(&systemRoleAuthApis)
	common.MySQLDB.Create(&systemRoleAuthAndPermissionApis)
	common.MySQLDB.Create(&systemMenuAuthApis)
	common.MySQLDB.Create(&systemMenuAuthAndPermissionApis)
	common.MySQLDB.Create(&systemDepartmentAuthApis)
	common.MySQLDB.Create(&systemDepartmentAuthAndPermissionApis)
	common.MySQLDB.Create(&systemJobPositionAuthApis)
	common.MySQLDB.Create(&systemJobPositionAuthAndPermissionApis)
}
