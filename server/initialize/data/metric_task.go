package data

import (
	"stellar/common"
	"stellar/model"
	"stellar/pkg/trans"
	"time"
)

// 指标任务数据
var metricTasks = []model.MetricTask{
	{
		BaseModel: model.BaseModel{
			Id: 1,
		},
		Name:           "测试指标任务1",
		Description:    "测试指标任务描述1",
		MetricName:     "test_metric_task1",
		MetricLabel:    "{name=\"DK\", age=\"18\"}",
		MetricType:     "gauge",
		MetricHelp:     "This is a test metric task1",
		CronExpression: "0/30 * * * * *",
		TaskContent:    "select 1",
		DatasourceId:   1,
		CreatorId:      1,
		UpdaterId:      1,
		Status:         trans.Uint(1),
		Timeout:        30,
		NextRunTime:    time.Now(),
	},
	{
		BaseModel: model.BaseModel{
			Id: 2,
		},
		Name:           "测试指标任务2",
		Description:    "测试指标任务描述2",
		MetricName:     "test_metric_task2",
		MetricLabel:    "{name=\"DK\", age=\"18\"}",
		MetricType:     "gauge",
		MetricHelp:     "This is a test metric task2",
		CronExpression: "0/30 * * * * *",
		TaskContent:    "select 1",
		DatasourceId:   1,
		CreatorId:      1,
		UpdaterId:      1,
		Status:         trans.Uint(1),
		Timeout:        30,
		NextRunTime:    time.Now(),
	},
}

// 指标任务初始化
func InitMetricTaskData() {
	common.MySQLDB.Exec("TRUNCATE TABLE metric_task")
	common.MySQLDB.Create(&metricTasks)
}
