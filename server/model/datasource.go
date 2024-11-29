package model

// 数据源模型
type Datasource struct {
	BaseModel
	DatasourceTypeId uint            `gorm:"column:datasourceTypeId;comment:数据源类型id" json:"datasourceTypeId"`
	DatasourceType   *DatasourceType `gorm:"foreignKey:DatasourceTypeId;" json:"datasourceType,omitempty"`
	Name             string          `gorm:"column:name;uniqueIndex:uidx_name;comment:数据源名称" json:"name"`
	Description      string          `gorm:"column:description;comment:数据源描述" json:"description"`
	Url              string          `gorm:"column:url;comment:数据源地址" json:"url"`
	Username         string          `gorm:"column:username;comment:数据源用户名" json:"username"`
	Password         string          `gorm:"column:password;comment:数据源密码" json:"password"`
	Status           *uint           `gorm:"column:status;type:tinyint(1);default:1;comment:数据源状态(0=禁用,1=启用)" json:"status"`
	CreatorId        uint            `gorm:"column:creatorId;comment:创建人id" json:"creatorId"`
	Creator          *SystemUser     `gorm:"foreignKey:CreatorId;" json:"creator,omitempty"`
	UpdaterId        uint            `gorm:"column:updaterId;comment:更新人id" json:"updaterId"`
	Updater          *SystemUser     `gorm:"foreignKey:UpdaterId;" json:"updater,omitempty"`
}

// 表名设置
func (Datasource) TableName() string {
	return "datasource"
}
