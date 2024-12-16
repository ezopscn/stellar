package model

// 系统API模型
type SystemApi struct {
	BaseModel
	Name        string             `json:"name" gorm:"column:name;comment:API名称"`
	Path        string             `json:"path" gorm:"column:path;comment:API路径"`
	Method      string             `json:"method" gorm:"column:method;comment:访问方法"`
	Description string             `json:"description" gorm:"column:description;comment:描述"`
	CategoryId  uint               `json:"categoryId" gorm:"column:categoryId;comment:分类ID"`
	Category    *SystemApiCategory `json:"category" gorm:"foreignKey:CategoryId;comment:分类"`
	SystemRoles []SystemRole       `gorm:"many2many:system_role_api_relation" json:"SystemRoles,omitempty"`
}

// 表名
func (s *SystemApi) TableName() string {
	return "system_api"
}
