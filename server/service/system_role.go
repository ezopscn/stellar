package service

import (
	"fmt"
	"stellar/common"
	"stellar/model"
	"stellar/pkg/gedis"
	"stellar/pkg/utils"
	"strings"
	"time"
)

// 获取角色列表
func GetSystemRoleListService() (roles []model.SystemRole, err error) {
	err = common.MySQLDB.Find(&roles).Error
	return
}

// 获取当前角色的API列表
func GetSystemRoleApiListService(roleKeyword string) (apiList []string, err error) {
	// 查询 Redis 中是否存在当前角色的 API 列表
	key := fmt.Sprintf("%s:%s", common.RKP.SystemRoleApis, roleKeyword)
	conn := gedis.NewRedisConnection()
	result := conn.GetString(key).Unwrap()
	if result == "" {
		var apis []model.SystemApi
		// 判断角色是不是管理员，管理员查询所有 API，其他角色查询当前角色的 API
		if utils.IsStringInSlice(roleKeyword, common.SystemRoleAdminList) {
			err = common.MySQLDB.Where("needPermission = ?", 1).Find(&apis).Error
		} else {
			var role model.SystemRole
			err = common.MySQLDB.Where("keyword = ?", roleKeyword).Preload("SystemApis").First(&role).Error
			apis = role.SystemApis
		}
		// 将查询到的 API 列表存放到 Redis 中，过期时间 1 分钟
		for _, api := range apis {
			apiList = append(apiList, api.Path)
		}
		apiListStr := strings.Join(apiList, ",")
		conn.Set(key, apiListStr, gedis.WithExpire(time.Second*60))
	} else {
		apiList = strings.Split(result, ",")
	}
	return
}
