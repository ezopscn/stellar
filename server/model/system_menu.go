package model

// 菜单模型
type SystemMenu struct {
	Id          uint         `gorm:"column:id;primaryKey;comment:自增编号" json:"id"`
	Label       string       `gorm:"column:label;uniqueIndex:uidx_label;comment:菜单名称" json:"label"`
	Icon        string       `gorm:"column:icon;comment:菜单图标" json:"icon"`
	Key         string       `gorm:"column:key;uniqueIndex:uidx_key;comment:菜单路径" json:"key"`
	Sort        uint         `gorm:"column:sort;comment:排序" json:"sort"`
	ParentId    uint         `gorm:"column:parentId;comment:父id" json:"parentId"`
	Children    []SystemMenu `gorm:"-" json:"children"`
	SystemRoles []SystemRole `gorm:"many2many:system_role_menu_relation" json:"SystemRoles,omitempty"`
}

// 表名设置
func (SystemMenu) TableName() string {
	return "system_menu"
}
