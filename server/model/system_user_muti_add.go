package model

// 用户批量导入任务模型
type SystemUserMutiAddTask struct {
	BaseModel
	CreatorId     uint        `gorm:"column:creatorId;comment:创建者id" json:"creatorId"`
	Creator       *SystemUser `gorm:"foreignKey:CreatorId" json:"creator,omitempty"`
	UserNumber    uint        `gorm:"column:userNumber;comment:用户数量" json:"userNumber"`
	SuccessNumber uint        `gorm:"column:successNumber;comment:成功数量" json:"successNumber"`
	FailNumber    uint        `gorm:"column:failNumber;comment:失败数量" json:"failNumber"`
	Status        *uint       `gorm:"column:status;type:tinyint(1);default:1;comment:状态(1=进行中,2=已完成)" json:"status"`
}

// 表名设置
func (SystemUserMutiAddTask) TableName() string {
	return "system_user_muti_add_task"
}

// 用户批量导入任务记录模型
type SystemUserMutiAddDetail struct {
	BaseModel
	TaskId       uint                   `gorm:"column:taskId;comment:任务id" json:"taskId"`
	Task         *SystemUserMutiAddTask `gorm:"foreignKey:TaskId" json:"task,omitempty"`
	Username     string                 `gorm:"column:username;comment:用户名" json:"username"`
	CNName       string                 `gorm:"column:cnName;comment:中文名" json:"cnName"`
	ENName       string                 `gorm:"column:enName;comment:英文名" json:"enName"`
	Email        string                 `gorm:"column:email;comment:邮箱" json:"email"`
	Phone        string                 `gorm:"column:phone;comment:手机号" json:"phone"`
	HidePhone    string                 `gorm:"column:hidePhone;comment:是否隐藏手机号(0=否,1=是)" json:"hidePhone"`
	Gender       string                 `gorm:"column:gender;comment:性别(1=男,2=女,3=未知)" json:"gender"`
	Departments  string                 `gorm:"column:departments;comment:部门" json:"departments"`
	JobPositions string                 `gorm:"column:jobPositions;comment:职位" json:"jobPositions"`
	Role         string                 `gorm:"column:role;comment:角色" json:"role"`
	Description  string                 `gorm:"column:description;comment:描述" json:"description"`
	Status       *uint                  `gorm:"column:status;type:tinyint(1);default:1;comment:状态(1=进行中,2=已完成,3=失败)" json:"status"`
	Result       string                 `gorm:"column:result;comment:结果" json:"result"`
}

// 表名设置
func (SystemUserMutiAddDetail) TableName() string {
	return "system_user_muti_add_detail"
}
