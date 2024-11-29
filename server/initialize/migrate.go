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
	common.MySQLDB.AutoMigrate(new(model.SystemUser))
	common.MySQLDB.AutoMigrate(new(model.SystemRole))
	common.MySQLDB.AutoMigrate(new(model.SystemMenu))
	common.MySQLDB.AutoMigrate(new(model.DatasourceType))
	common.MySQLDB.AutoMigrate(new(model.Datasource))
	common.MySQLDB.AutoMigrate(new(model.MetricTask))
	common.MySQLDB.AutoMigrate(new(model.MetricTaskLog))
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Table migrate completed")
}

// 迁移数据
func MigrateData() {
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Start migrate data")
	data.InitUserData()
	data.InitRoleData()
	data.InitMenuData()
	data.InitDatasourceType()
	data.InitMetricTaskData()
	fmt.Println(time.Now().Format(common.TimeMillisecondFormat), "Data migrate completed")
}
