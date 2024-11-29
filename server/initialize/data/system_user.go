package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
	"stellar/pkg/utils"
)

const (
	defaultPassword = "p@ssw0rd"                                                            // 默认密码
	defaultAvatar   = "https://gw.alipayobjects.com/zos/rmsportal/BiazfanxmamNRoxxVxka.png" // 默认头像
)

// 用户初始化数据
var systemUsers = []model.SystemUser{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		Username:     "admin",
		Password:     utils.CryptoPassword(defaultPassword),
		CNName:       "超管",
		ENName:       "Administrator",
		Email:        "admin@onething.net",
		Phone:        "18888888888",
		HidePhone:    trans.Uint(0),
		Avatar:       defaultAvatar,
		Department:   "系统工程部",
		JobPosition:  "高级运维工程师",
		Description:  "系统超级管理员",
		Status:       trans.Uint(1),
		CreatorId:    1,
		SystemRoleId: 1,
	},
}

// 用户初始化
func InitUserData() {
	for _, user := range systemUsers {
		common.MySQLDB.Exec("DELETE FROM system_user WHERE id = ?", user.Id)
		common.MySQLDB.Create(&user)
	}
}
