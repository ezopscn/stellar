package model

import "github.com/dromara/carbon/v2"

// 指标任务日志模型
type MetricTaskLog struct {
	BaseModel
	MetricTaskId uint          `gorm:"column:metricTaskId;comment:指标任务id" json:"metricTaskId"`
	MetricTask   *MetricTask   `gorm:"foreignKey:MetricTaskId;" json:"metricTask,omitempty"`
	StartTime    carbon.Carbon `gorm:"column:startTime;comment:开始时间" json:"startTime"`
	EndTime      carbon.Carbon `gorm:"column:endTime;comment:结束时间" json:"endTime"`
	ClientId     string        `gorm:"column:clientId;comment:执行客户端id" json:"clientId"`
	Status       *uint         `gorm:"column:status;type:tinyint(1);default:1;comment:任务状态(0=失败,1=成功)" json:"status"`
	Result       string        `gorm:"column:result;comment:执行结果" json:"result"`
}

// 表名设置
func (MetricTaskLog) TableName() string {
	return "metric_task_log"
}
