package model

// 系统API分类模型
type SystemApiCategory struct {
	BaseModel
	Name        string `gorm:"column:name;comment:名称" json:"name"`
	Description string `gorm:"column:description;comment:描述" json:"description"`
	SystemApis  []uint `gorm:"-" json:"systemApis,omitempty"`
}

// 表名
func (s *SystemApiCategory) TableName() string {
	return "system_api_category"
}
