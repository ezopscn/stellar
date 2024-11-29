package service

import (
	"context"
	"stellar/common"
	"stellar/model"
	"time"
)

// 检测任务并发布到消息队列
func CheckAndPublishMetricTask() {
	var metricTasks []model.MetricTask
	now := time.Now()
	err := common.MySQLDB.Where("nextRunTime <= ? AND status = ?", now, 1).Find(&metricTasks).Error
	if err != nil {
		common.SystemLog.Error("查询指标任务失败: ", err.Error())
		return
	}

	for _, metricTask := range metricTasks {
		common.SystemLog.Debug("检测到指标任务: ", metricTask.Name, ": ", metricTask.MetricName+metricTask.MetricLabel)

		// 发布到消息队列
		common.RedisCache.Publish(context.Background(), "METRIC_TASK", metricTask)
		// 更新任务下次执行时间
	}
}

// 订阅 Channel: METRIC_TASK
func SubscribeMetricTask() {
	ch := common.RedisCache.Subscribe(context.Background(), "METRIC_TASK")
	for msg := range ch.Channel() {
		common.SystemLog.Debug("获取到指标任务: ", msg.Payload)
	}
}
