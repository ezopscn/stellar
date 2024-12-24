package data

import (
	"fmt"
	"stellar/pkg/utils"
)

var (
	defaultPassword = "p@ssw0rd"                   // 默认密码
	defaultAvatar   = "/images/avatar/default.png" // 默认头像
	defaultCreator  = "default,默认,Default,0"
)

// 随机生成男头像
func RandomMaleAvatar() string {
	idStr := utils.RandString(1, "123456")
	return fmt.Sprintf("/images/avatar/male_%s.svg", idStr)
}

// 随机生成女头像
func RandomFemaleAvatar() string {
	idStr := utils.RandString(1, "123456")
	return fmt.Sprintf("/images/avatar/female_%s.svg", idStr)
}
