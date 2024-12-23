package data

import (
	"stellar/common"
	"stellar/model"
)

// 菜单初始化
var systemMenus = []model.SystemMenu{
	{
		Id:       1100,
		Label:    "工作空间",
		Icon:     "DesktopOutlined",
		Key:      "/dashboard",
		Sort:     1100,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
			systemRoles[3],
		},
	},
	{
		Id:       1800,
		Label:    "数据源",
		Icon:     "ApiOutlined",
		Key:      "/datasource",
		Sort:     1800,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
	},
	{
		Id:       1900,
		Label:    "系统设置",
		Icon:     "ClusterOutlined",
		Key:      "/system",
		Sort:     1900,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1910,
				Label:    "部门管理",
				Key:      "/system/department",
				Sort:     1910,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1920,
				Label:    "职位管理",
				Key:      "/system/job-position",
				Sort:     1920,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1930,
				Label:    "用户管理",
				Key:      "/system/user",
				Sort:     1930,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1940,
				Label:    "角色管理",
				Key:      "/system/role",
				Sort:     1940,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1950,
				Label:    "菜单管理",
				Key:      "/system/menu",
				Sort:     1950,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1960,
				Label:    "接口管理",
				Key:      "/system/api",
				Sort:     1960,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1970,
				Label:    "权限管理",
				Key:      "/system/permission",
				Sort:     1970,
				ParentId: 1900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		Id:       2000,
		Label:    "个人中心",
		Icon:     "UserOutlined",
		Key:      "/me",
		Sort:     2000,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
			systemRoles[3],
		},
	},
	{
		Id:       9900,
		Label:    "安全审计",
		Icon:     "FileProtectOutlined",
		Key:      "/security-audit",
		Sort:     9900,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
			systemRoles[3],
		},
		Children: []model.SystemMenu{
			{
				Id:       9910,
				Label:    "登录日志",
				Key:      "/security-audit/login",
				Sort:     9910,
				ParentId: 9900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
					systemRoles[3],
				},
			},
			{
				Id:       9920,
				Label:    "操作日志",
				Key:      "/security-audit/operation",
				Sort:     9920,
				ParentId: 9900,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
					systemRoles[3],
				},
			},
		},
	},
}

// 递归插入数据方法
func insertMenusData(menus []model.SystemMenu) {
	for _, menu := range menus {
		common.MySQLDB.Create(&menu)
		if len(menu.Children) > 0 {
			insertMenusData(menu.Children)
		}
	}
}

// 菜单初始化
func InitMenuData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_menu")
	common.MySQLDB.Exec("TRUNCATE TABLE system_role_menu_relation")
	insertMenusData(systemMenus)
}
