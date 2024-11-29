package model

import (
	"time"
)

// 指标任务模型
type MetricTask struct {
	BaseModel
	Name           string      `gorm:"column:name;uniqueIndex:uidx_name;comment:任务名称" json:"name"`
	Description    string      `gorm:"column:description;comment:任务描述" json:"description"`
	MetricName     string      `gorm:"column:metricName;uniqueIndex:uidx_metricName_metricLabel;comment:指标名称" json:"metricName"`
	MetricLabel    string      `gorm:"column:metricLabel;uniqueIndex:uidx_metricName_metricLabel;comment:指标标签" json:"metricLabel"`
	MetricType     string      `gorm:"column:metricType;comment:指标类型" json:"metricType"`
	MetricHelp     string      `gorm:"column:metricHelp;comment:指标帮助" json:"metricHelp"`
	DatasourceId   uint        `gorm:"column:datasourceId;comment:数据源id" json:"datasourceId"`
	Datasource     *Datasource `gorm:"foreignKey:DatasourceId;" json:"datasource,omitempty"`
	CronExpression string      `gorm:"column:cronExpression;not null;comment:Cron表达式" json:"cronExpression"`
	TaskContent    string      `gorm:"column:taskContent;not null;comment:任务内容" json:"taskContent"`
	NextRunTime    time.Time   `gorm:"column:nextRunTime;comment:下次执行时间" json:"nextRunTime"`
	Timeout        uint        `gorm:"column:timeout;type:int(11);default:0;comment:任务超时时间(秒)" json:"timeout"`
	Status         *uint       `gorm:"column:status;type:tinyint(1);default:1;comment:任务状态(0=禁用,1=启用)" json:"status"`
	CreatorId      uint        `gorm:"column:creatorId;comment:创建人id" json:"creatorId"`
	Creator        *SystemUser `gorm:"foreignKey:CreatorId;" json:"creator,omitempty"`
	UpdaterId      uint        `gorm:"column:updaterId;comment:更新人id" json:"updaterId"`
	Updater        *SystemUser `gorm:"foreignKey:UpdaterId;" json:"updater,omitempty"`
}

// 表名设置
func (MetricTask) TableName() string {
	return "metric_task"
}
