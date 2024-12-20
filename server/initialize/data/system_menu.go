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
		Id:       1700,
		Label:    "指标配置",
		Icon:     "ConsoleSqlOutlined",
		Key:      "/metrics",
		Sort:     1700,
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
		Key:      "/datasources",
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
				Key:      "/system/departments",
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
				Key:      "/system/jobpositions",
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
				Key:      "/system/users",
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
				Key:      "/system/roles",
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
				Key:      "/system/menus",
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
				Key:      "/system/apis",
				Sort:     1960,
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
		Key:      "/securityaudit",
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
				Key:      "/securityaudit/login",
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
				Key:      "/securityaudit/operation",
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
