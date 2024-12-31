package model

// 部门模型
type SystemDepartment struct {
	BaseModel
	Name        string             `gorm:"column:name;comment:部门名称" json:"name"`
	Description string             `gorm:"column:description;comment:部门描述" json:"description"`
	LeaderId    uint               `gorm:"column:leaderId;comment:部门领导id" json:"leaderId"`
	Leader      *SystemUser        `gorm:"foreignKey:LeaderId" json:"leader,omitempty"`
	Creator     string             `gorm:"column:creator;comment:创建者（由用户名，中文名，英文名，ID组成，格式：username,cnName,enName,id）" json:"creator"`
	ParentId    uint               `gorm:"column:parentId;comment:父id" json:"parentId"`
	Children    []SystemDepartment `gorm:"-" json:"children"`
	SystemUsers []SystemUser       `gorm:"many2many:system_user_department_relation" json:"SystemUsers,omitempty"`
}

// 表名设置
func (SystemDepartment) TableName() string {
	return "system_department"
}
