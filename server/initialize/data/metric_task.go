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
		Name:           "测试指标任务",
		Description:    "测试指标任务描述",
		MetricName:     "test_metric_task",
		MetricLabel:    "{name=\"DK\", age=\"18\"}",
		MetricType:     "gauge",
		MetricHelp:     "This is a test metric task",
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
	for _, metricTask := range metricTasks {
		common.MySQLDB.Exec("DELETE FROM metric_task WHERE id = ?", metricTask.Id)
		common.MySQLDB.Create(&metricTask)
	}
}
