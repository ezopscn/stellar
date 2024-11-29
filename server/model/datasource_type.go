package model

// 数据源类型
type DatasourceType struct {
	Id          uint   `gorm:"column:id;primaryKey;comment:自增编号" json:"id"`
	Name        string `gorm:"column:name;uniqueIndex:uidx_name;comment:数据源类型名称" json:"name"`
	Description string `gorm:"column:description;comment:数据源类型描述" json:"description"`
	Logo        string `gorm:"column:logo;comment:数据源类型logo" json:"logo"`
}

// 表名设置
func (DatasourceType) TableName() string {
	return "datasource_type"
}
