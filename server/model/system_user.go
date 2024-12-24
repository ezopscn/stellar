package model

import (
	"github.com/dromara/carbon/v2"
)

// 用户模型
type SystemUser struct {
	BaseModel
	Username               string              `gorm:"column:username;uniqueIndex:uidx_username;comment:用户名" json:"username"`
	Password               string              `gorm:"column:password;comment:密码" json:"-"`
	CNName                 string              `gorm:"column:cnName;comment:中文名" json:"cnName"`
	ENName                 string              `gorm:"column:enName;comment:英文名" json:"enName"`
	Email                  string              `gorm:"column:email;uniqueIndex:uidx_email;comment:邮箱" json:"email"`
	Phone                  string              `gorm:"column:phone;uniqueIndex:uidx_phone;comment:手机号" json:"phone"`
	HidePhone              *uint               `gorm:"column:hidePhone;default:1;comment:是否隐藏手机号(0=不隐藏,1=隐藏)" json:"hidePhone"`
	Gender                 *uint               `gorm:"column:gender;default:3;comment:性别(1=男,2=女,3=未知)" json:"gender"`
	Avatar                 string              `gorm:"column:avatar;comment:头像" json:"avatar"`
	SystemDepartments      []SystemDepartment  `gorm:"many2many:system_user_department_relation" json:"systemDepartments,omitempty"`
	SystemJobPositions     []SystemJobPosition `gorm:"many2many:system_user_job_position_relation" json:"systemJobPositions,omitempty"`
	Description            string              `gorm:"column:description;comment:描述" json:"description"`
	LastLoginIP            string              `gorm:"column:lastLoginIP;comment:最后一次登录IP" json:"lastLoginIP"`
	LastLoginTime          carbon.Carbon       `gorm:"column:lastLoginTime;comment:最后一次登录时间" json:"lastLoginTime"`
	LastChangePasswordTime carbon.Carbon       `gorm:"column:lastChangePasswordTime;comment:最后一次修改密码时间" json:"lastChangePasswordTime"`
	Status                 *uint               `gorm:"column:status;type:tinyint(1);default:1;comment:用户状态(0=禁用,1=正常)" json:"status"`
	Creator                string              `gorm:"column:creator;comment:创建者（由用户名，中文名，英文名，ID组成，格式：username,cnName,enName,id）" json:"creator"`
	SystemRoleId           uint                `gorm:"column:systemRoleId;comment:角色id" json:"systemRoleId"`
	SystemRole             *SystemRole         `gorm:"foreignKey:SystemRoleId;" json:"systemRole,omitempty"`
}

// 表名设置
func (SystemUser) TableName() string {
	return "system_user"
}
