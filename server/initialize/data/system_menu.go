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
		Id:       1200,
		Label:    "即时查询",
		Icon:     "ConsoleSqlOutlined",
		Key:      "/query",
		Sort:     1200,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
	},
	{
		Id:       1300,
		Label:    "告警管理",
		Icon:     "AlertOutlined",
		Key:      "/alert",
		Sort:     1300,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1310,
				Label:    "活跃告警",
				Key:      "/alert/active",
				Sort:     1310,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1320,
				Label:    "告警规则",
				Key:      "/alert/rule",
				Sort:     1320,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1330,
				Label:    "告警订阅",
				Key:      "/alert/subscription",
				Sort:     1330,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1340,
				Label:    "告警屏蔽",
				Key:      "/alert/block",
				Sort:     1340,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1350,
				Label:    "告警历史",
				Key:      "/alert/history",
				Sort:     1350,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1360,
				Label:    "告警回调",
				Key:      "/alert/callback",
				Sort:     1360,
				ParentId: 1300,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
		},
	},
	{
		Id:       1400,
		Label:    "告警通知",
		Icon:     "MailOutlined",
		Key:      "/alert-notification",
		Sort:     1400,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				Id:       1410,
				Label:    "通知媒介",
				Key:      "/alert-notification/media",
				Sort:     1410,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1420,
				Label:    "通知模板",
				Key:      "/alert-notification/template",
				Sort:     1420,
				ParentId: 1400,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		Id:       1500,
		Label:    "告警分组",
		Icon:     "ProjectOutlined",
		Key:      "/alert-group",
		Sort:     1500,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
		Children: []model.SystemMenu{
			{
				Id:       1510,
				Label:    "项目管理",
				Key:      "/alert-group/project",
				Sort:     1510,
				ParentId: 1500,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1520,
				Label:    "团队管理",
				Key:      "/alert-group/team",
				Sort:     1520,
				ParentId: 1500,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1530,
				Label:    "人员排班",
				Key:      "/alert-group/schedule",
				Sort:     1530,
				ParentId: 1500,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		Id:       1600,
		Label:    "数据来源",
		Icon:     "ApiOutlined",
		Key:      "/datasource",
		Sort:     1600,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
		},
	},
	{
		Id:       1700,
		Label:    "系统设置",
		Icon:     "ClusterOutlined",
		Key:      "/system",
		Sort:     1700,
		ParentId: 0,
		SystemRoles: []model.SystemRole{
			systemRoles[0],
			systemRoles[1],
			systemRoles[2],
		},
		Children: []model.SystemMenu{
			{
				Id:       1710,
				Label:    "部门管理",
				Key:      "/system/department",
				Sort:     1710,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1720,
				Label:    "职位管理",
				Key:      "/system/job-position",
				Sort:     1720,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1730,
				Label:    "用户管理",
				Key:      "/system/user",
				Sort:     1730,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
					systemRoles[2],
				},
			},
			{
				Id:       1740,
				Label:    "角色管理",
				Key:      "/system/role",
				Sort:     1740,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
			{
				Id:       1750,
				Label:    "菜单管理",
				Key:      "/system/menu",
				Sort:     1750,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1760,
				Label:    "接口管理",
				Key:      "/system/api",
				Sort:     1760,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
				},
			},
			{
				Id:       1770,
				Label:    "权限管理",
				Key:      "/system/permission",
				Sort:     1770,
				ParentId: 1700,
				SystemRoles: []model.SystemRole{
					systemRoles[0],
					systemRoles[1],
				},
			},
		},
	},
	{
		Id:       1800,
		Label:    "个人中心",
		Icon:     "UserOutlined",
		Key:      "/me",
		Sort:     1800,
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
	insertMenusData(systemMenus)
}
