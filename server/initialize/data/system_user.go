package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
	"stellar/pkg/utils"
)

const (
	defaultPassword     = "p@ssw0rd"                                                            // 默认密码
	defaultAvatar       = "https://gw.alipayobjects.com/zos/rmsportal/jZUIxmJycoymBprLOUbT.png" // 默认头像
	defaultMaleAvatar   = "https://gw.alipayobjects.com/zos/rmsportal/ubnKSIfAJTxIgXOKlciN.png" // 默认头像(男)
	defaultFemaleAvatar = "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png" // 默认头像(女)
)

// 用户初始化数据
var systemUsers = []model.SystemUser{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		Username:     "super",
		Password:     utils.CryptoPassword(defaultPassword),
		CNName:       "超管",
		ENName:       "Super",
		Email:        "super@ezops.cn",
		Phone:        "19999999999",
		HidePhone:    trans.Uint(1),
		Gender:       trans.Uint(1),
		Avatar:       defaultMaleAvatar,
		Department:   "系统工程部",
		JobPosition:  "高级运维工程师",
		Description:  "系统超级管理员",
		Status:       trans.Uint(1),
		CreatorId:    1,
		SystemRoleId: 1,
	},
	{
		BaseModel: model.BaseModel{
			Id: 2,
		},
		Username:     "admin",
		Password:     utils.CryptoPassword(defaultPassword),
		CNName:       "管理员",
		ENName:       "Administrator",
		Email:        "admin@ezops.cn",
		Phone:        "18888888888",
		HidePhone:    trans.Uint(1),
		Gender:       trans.Uint(1),
		Avatar:       defaultMaleAvatar,
		Department:   "系统工程部",
		JobPosition:  "高级运维工程师",
		Description:  "系统超级管理员",
		Status:       trans.Uint(1),
		CreatorId:    1,
		SystemRoleId: 2,
	},
	{
		BaseModel: model.BaseModel{
			Id: 3,
		},
		Username:     "devops",
		Password:     utils.CryptoPassword(defaultPassword),
		CNName:       "运维",
		ENName:       "DevOps",
		Email:        "devops@ezops.cn",
		Phone:        "17777777777",
		HidePhone:    trans.Uint(0),
		Gender:       trans.Uint(2),
		Avatar:       defaultFemaleAvatar,
		Department:   "系统工程部",
		JobPosition:  "高级运维工程师",
		Description:  "系统高级运维工程师",
		Status:       trans.Uint(0),
		CreatorId:    1,
		SystemRoleId: 3,
	},
	{
		BaseModel: model.BaseModel{
			Id: 4,
		},
		Username:     "guest",
		Password:     utils.CryptoPassword(defaultPassword),
		CNName:       "访客",
		ENName:       "Guest",
		Email:        "guest@ezops.cn",
		Phone:        "16666666666",
		HidePhone:    trans.Uint(0),
		Gender:       trans.Uint(0),
		Avatar:       defaultAvatar,
		Department:   "系统工程部",
		JobPosition:  "访客",
		Description:  "系统访客",
		Status:       trans.Uint(0),
		CreatorId:    1,
		SystemRoleId: 4,
	},
}

// 用户初始化
func InitUserData() {
	for _, user := range systemUsers {
		common.MySQLDB.Exec("DELETE FROM system_user WHERE id = ?", user.Id)
		common.MySQLDB.Create(&user)
	}
}
