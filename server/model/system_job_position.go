package model

// 岗位模型
type SystemJobPosition struct {
	BaseModel
	Name        string       `gorm:"column:name;uniqueIndex:uidx_name;comment:岗位名称" json:"name"`
	Description string       `gorm:"column:description;comment:岗位描述" json:"description"`
	Creator     string       `gorm:"column:creator;comment:创建者（由中文名，英文名，ID组成，格式：张三,ZhangSan,1）" json:"creator"`
	SystemUsers []SystemUser `gorm:"many2many:system_user_job_position_relation" json:"systemUsers,omitempty"`
}

// 表名设置
func (SystemJobPosition) TableName() string {
	return "system_job_position"
}
