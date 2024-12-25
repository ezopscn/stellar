package model

// 系统API模型
type SystemApi struct {
	BaseModel
	Name           string             `gorm:"column:name;comment:API名称" json:"name"`
	Path           string             `gorm:"column:path;comment:API路径" json:"path"`
	Method         string             `gorm:"column:method;comment:访问方法" json:"method"`
	NeedPermission *uint              `gorm:"column:needPermission;type:tinyint(1);default:0;comment:是否需要权限(0=不需要,1=需要)" json:"needPermission"`
	Description    string             `gorm:"column:description;comment:描述" json:"description"`
	CategoryId     uint               `gorm:"column:categoryId;comment:分类ID" json:"categoryId"`
	Category       *SystemApiCategory `gorm:"foreignKey:categoryId;comment:分类" json:"category"`
	SystemRoles    []SystemRole       `gorm:"many2many:system_role_api_relation" json:"systemRoles,omitempty"`
}

// 表名
func (s *SystemApi) TableName() string {
	return "system_api"
}
