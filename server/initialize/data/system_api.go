package data

import (
	"stellar/common"
	"stellar/model"
)

// 系统API数据
var SystemApis = []model.SystemApi{
	{Name: "获取部门列表", Path: "/system/department/list", Method: "GET", Description: "获取部门列表数据 API 接口", CategoryId: 1, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建部门", Path: "/system/department/create", Method: "POST", Description: "创建部门数据 API 接口", CategoryId: 1},
	{Name: "更新部门", Path: "/system/department/update", Method: "PUT", Description: "更新部门数据 API 接口", CategoryId: 1},
	{Name: "删除部门", Path: "/system/department/delete", Method: "DELETE", Description: "删除部门数据 API 接口", CategoryId: 1},
	{Name: "获取职位列表", Path: "/system/jobPosition/list", Method: "GET", Description: "获取岗位列表数据 API 接口", CategoryId: 2, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建职位", Path: "/system/jobPosition/create", Method: "POST", Description: "创建岗位数据 API 接口", CategoryId: 2},
	{Name: "更新职位", Path: "/system/jobPosition/update", Method: "PUT", Description: "更新岗位数据 API 接口", CategoryId: 2},
	{Name: "删除职位", Path: "/system/jobPosition/delete", Method: "DELETE", Description: "删除岗位数据 API 接口", CategoryId: 2},
	{Name: "获取用户列表", Path: "/system/user/list", Method: "GET", Description: "获取用户列表数据 API 接口", CategoryId: 3, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建用户", Path: "/system/user/create", Method: "POST", Description: "创建用户数据 API 接口", CategoryId: 3},
	{Name: "更新用户", Path: "/system/user/update", Method: "PUT", Description: "更新用户数据 API 接口", CategoryId: 3},
	{Name: "删除用户", Path: "/system/user/delete", Method: "DELETE", Description: "删除用户数据 API 接口", CategoryId: 3},
	{Name: "获取角色列表", Path: "/system/role/list", Method: "GET", Description: "获取角色列表数据 API 接口", CategoryId: 4, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建角色", Path: "/system/role/create", Method: "POST", Description: "创建角色数据 API 接口", CategoryId: 4},
	{Name: "更新角色", Path: "/system/role/update", Method: "PUT", Description: "更新角色数据 API 接口", CategoryId: 4},
	{Name: "删除角色", Path: "/system/role/delete", Method: "DELETE", Description: "删除角色数据 API 接口", CategoryId: 4},
	{Name: "获取菜单列表", Path: "/system/menu/list", Method: "GET", Description: "获取菜单列表数据 API 接口", CategoryId: 5, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "获取菜单树", Path: "/system/menu/tree", Method: "GET", Description: "获取菜单树数据 API 接口", CategoryId: 5, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建菜单", Path: "/system/menu/create", Method: "POST", Description: "创建菜单数据 API 接口", CategoryId: 5},
	{Name: "更新菜单", Path: "/system/menu/update", Method: "PUT", Description: "更新菜单数据 API 接口", CategoryId: 5},
	{Name: "删除菜单", Path: "/system/menu/delete", Method: "DELETE", Description: "删除菜单数据 API 接口", CategoryId: 5},
	{Name: "获取API分类列表", Path: "/system/apiCategory/list", Method: "GET", Description: "获取API分类列表数据 API 接口", CategoryId: 6, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建API分类", Path: "/system/apiCategory/create", Method: "POST", Description: "创建API分类数据 API 接口", CategoryId: 6},
	{Name: "更新API分类", Path: "/system/apiCategory/update", Method: "PUT", Description: "更新API分类数据 API 接口", CategoryId: 6},
	{Name: "删除API分类", Path: "/system/apiCategory/delete", Method: "DELETE", Description: "删除API分类数据 API 接口", CategoryId: 6},
	{Name: "获取API列表", Path: "/system/api/list", Method: "GET", Description: "获取API列表数据 API 接口", CategoryId: 7, SystemRoles: []model.SystemRole{systemRoles[2]}},
	{Name: "创建API", Path: "/system/api/create", Method: "POST", Description: "创建API数据 API 接口", CategoryId: 7},
	{Name: "更新API", Path: "/system/api/update", Method: "PUT", Description: "更新API数据 API 接口", CategoryId: 7},
	{Name: "删除API", Path: "/system/api/delete", Method: "DELETE", Description: "删除API数据 API 接口", CategoryId: 7},
}

// 初始化系统API数据
func InitSystemApiData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_api")
	common.MySQLDB.Create(&SystemApis)
}
