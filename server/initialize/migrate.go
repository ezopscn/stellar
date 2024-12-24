package initialize

import (
	"fmt"
	"stellar/common"
	"stellar/initialize/data"
	"stellar/model"
	"time"
)

// 迁移数据表
func MigrateTable() {
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Start migrate table")
	common.MySQLDB.AutoMigrate(new(model.SystemCasbinRuleTable))   // Casbin 规则
	common.MySQLDB.AutoMigrate(new(model.SystemDepartment))        // 部门
	common.MySQLDB.AutoMigrate(new(model.SystemJobPosition))       // 岗位
	common.MySQLDB.AutoMigrate(new(model.SystemUser))              // 用户
	common.MySQLDB.AutoMigrate(new(model.SystemUserMutiAddTask))   // 用户批量导入任务
	common.MySQLDB.AutoMigrate(new(model.SystemUserMutiAddDetail)) // 用户批量导入任务详情
	common.MySQLDB.AutoMigrate(new(model.SystemRole))              // 角色
	common.MySQLDB.AutoMigrate(new(model.SystemMenu))              // 菜单
	common.MySQLDB.AutoMigrate(new(model.SystemApiCategory))       // API分类
	common.MySQLDB.AutoMigrate(new(model.SystemApi))               // API
	common.MySQLDB.AutoMigrate(new(model.DatasourceType))          // 数据源类型
	common.MySQLDB.AutoMigrate(new(model.Datasource))              // 数据源
	common.MySQLDB.AutoMigrate(new(model.MetricTask))              // 指标任务
	common.MySQLDB.AutoMigrate(new(model.MetricTaskLog))           // 指标任务日志
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Table migrate completed")
}

// 迁移数据
func MigrateData() {
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Start migrate data")
	data.InitDepartmentData()
	data.InitJobPositionData()
	data.InitUserData()
	data.InitRoleData()
	data.InitMenuData()
	data.InitSystemApiCategoryData()
	data.InitSystemApiData()
	data.InitSystemCasbinRuleData()
	data.InitDatasourceType()
	data.InitMetricTaskData()
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Data migrate completed")
}
