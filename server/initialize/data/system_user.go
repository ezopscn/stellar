package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
	"stellar/pkg/utils"
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
		Avatar:       RandomMaleAvatar(),
		Description:  "系统超级管理员",
		Status:       trans.Uint(1),
		Creator:      defaultCreator,
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
		Avatar:       RandomMaleAvatar(),
		Description:  "系统管理员",
		Status:       trans.Uint(1),
		Creator:      defaultCreator,
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
		Avatar:       RandomFemaleAvatar(),
		Description:  "系统运维工程师",
		Status:       trans.Uint(0),
		Creator:      defaultCreator,
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
		Description:  "系统访客，只读用户",
		Status:       trans.Uint(0),
		Creator:      defaultCreator,
		SystemRoleId: 4,
	},
}

// 用户初始化
func InitUserData() {
	common.MySQLDB.Exec("TRUNCATE TABLE system_user")
	common.MySQLDB.Create(&systemUsers)
}
