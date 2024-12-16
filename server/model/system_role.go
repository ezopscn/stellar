package model

// 角色模型
type SystemRole struct {
	Id          uint         `gorm:"column:id;primaryKey;comment:自增编号" json:"id"`
	Name        string       `gorm:"column:name;uniqueIndex:uidx_name;comment:角色名称" json:"name"`
	Keyword     string       `gorm:"column:keyword;uniqueIndex:uidx_keyword;comment:角色关键字" json:"keyword"`
	Description string       `gorm:"column:description;not null;comment:角色描述" json:"description"`
	SystemUsers []uint       `gorm:"-" json:"systemUsers,omitempty"`
	SystemMenus []SystemMenu `gorm:"many2many:system_role_menu_relation" json:"systemMenus,omitempty"`
	SystemApis  []SystemApi  `gorm:"many2many:system_role_api_relation" json:"systemApis,omitempty"`
}

// 表名设置
func (SystemRole) TableName() string {
	return "system_role"
}
