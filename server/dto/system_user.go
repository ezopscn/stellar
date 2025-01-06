package dto

import (
	"regexp"
	"stellar/pkg/utils"
)

// 用户筛选请求
type SystemUserFilterRequest struct {
	Username          *string `form:"username" json:"username"`
	Name              *string `form:"name" json:"name"`
	Email             *string `form:"email" json:"email"`
	Phone             *string `form:"phone" json:"phone"`
	Status            *uint   `form:"status" json:"status"`
	Gender            *uint   `form:"gender" json:"gender"`
	SystemDepartment  *uint   `form:"systemDepartment" json:"systemDepartment"`
	SystemJobPosition *uint   `form:"systemJobPosition" json:"systemJobPosition"`
	SystemRole        *uint   `form:"systemRole" json:"systemRole"`
	PageNumber        *uint   `form:"pageNumber" json:"pageNumber"`
	PageSize          *uint   `form:"pageSize" json:"pageSize"`
	IsPagination      *bool   `form:"isPagination" json:"isPagination"`
}

// 添加用户请求
type SystemUserAddRequest struct {
	Username           *string `form:"username" json:"username"`
	Password           *string `form:"password" json:"password"`
	CNName             *string `form:"cnName" json:"cnName"`
	ENName             *string `form:"enName" json:"enName"`
	Email              *string `form:"email" json:"email"`
	Phone              *string `form:"phone" json:"phone"`
	HidePhone          *uint   `form:"hidePhone" json:"hidePhone"`
	Gender             *uint   `form:"gender" json:"gender"`
	SystemDepartments  []uint  `form:"systemDepartments" json:"systemDepartments"`
	SystemJobPositions []uint  `form:"systemJobPositions" json:"systemJobPositions"`
	SystemRole         *uint   `form:"systemRole" json:"systemRole"`
	Description        *string `form:"description" json:"description"`
}

// 添加用户参数校验
func (req *SystemUserAddRequest) Validate() (errList []string) {
	// 用户名校验
	if !utils.IsUsername(*req.Username) {
		errList = append(errList, "用户名格式错误")
	}
	// 密码校验
	if !utils.IsPassword(*req.Password) {
		errList = append(errList, "密码格式错误")
	}
	// 中文名校验
	if !utils.IsCNName(*req.CNName) {
		errList = append(errList, "中文名格式错误")
	}
	// 英文名校验
	if !utils.IsENName(*req.ENName) {
		errList = append(errList, "英文名格式错误")
	}
	// 邮箱校验
	if !utils.IsEmail(*req.Email) {
		errList = append(errList, "邮箱格式错误")
	}
	// 手机号校验
	if !utils.IsPhoneNumber(*req.Phone) {
		errList = append(errList, "手机号格式错误")
	}
	// 隐藏手机号校验
	if *req.HidePhone != 0 && *req.HidePhone != 1 {
		errList = append(errList, "隐藏手机号格式错误，只能是 0 或 1")
	}
	// 性别校验
	if *req.Gender != 1 && *req.Gender != 2 && *req.Gender != 3 {
		errList = append(errList, "性别格式错误，只能是 1 或 2 或 3")
	}
	// 部门校验
	if len(req.SystemDepartments) == 0 {
		errList = append(errList, "部门不能为空")
	}
	// 职位校验
	if len(req.SystemJobPositions) == 0 {
		errList = append(errList, "职位不能为空")
	}
	// 角色校验
	if req.SystemRole == nil || *req.SystemRole == 0 {
		errList = append(errList, "角色不能为空")
	}
	return
}

// 批量添加用户请求
type SystemUserMultiAddRequest struct {
	Username           *string `form:"username" json:"username"`
	Password           *string `form:"password" json:"password"`
	CNName             *string `form:"cnName" json:"cnName"`
	ENName             *string `form:"enName" json:"enName"`
	Email              *string `form:"email" json:"email"`
	Phone              *string `form:"phone" json:"phone"`
	HidePhone          *string `form:"hidePhone" json:"hidePhone"`
	Gender             *string `form:"gender" json:"gender"`
	SystemDepartments  *string `form:"systemDepartments" json:"systemDepartments"`
	SystemJobPositions *string `form:"systemJobPositions" json:"systemJobPositions"`
	SystemRole         *string `form:"systemRole" json:"systemRole"`
	Description        *string `form:"description" json:"description"`
}

// 批量添加用户参数校验
func (req *SystemUserMultiAddRequest) Validate() (errList []string) {
	if !utils.IsUsername(*req.Username) {
		errList = append(errList, "用户名格式错误")
	}
	if !utils.IsPassword(*req.Password) {
		errList = append(errList, "密码格式错误")
	}
	if !utils.IsCNName(*req.CNName) {
		errList = append(errList, "中文名格式错误")
	}
	if !utils.IsENName(*req.ENName) {
		errList = append(errList, "英文名格式错误")
	}
	if !utils.IsEmail(*req.Email) {
		errList = append(errList, "邮箱格式错误")
	}
	if !utils.IsPhoneNumber(*req.Phone) {
		errList = append(errList, "手机号格式错误")
	}
	if *req.HidePhone != "0" && *req.HidePhone != "1" {
		errList = append(errList, "隐藏手机号格式错误，只能是 0 或 1")
	}
	if *req.Gender != "1" && *req.Gender != "2" && *req.Gender != "3" {
		errList = append(errList, "性别格式错误，只能是 1 或 2 或 3")
	}
	if req.SystemDepartments == nil || *req.SystemDepartments == "" || !regexp.MustCompile(`^(\d+)(,\d+)*$`).MatchString(*req.SystemDepartments) {
		errList = append(errList, "部门格式错误")
	}
	if req.SystemJobPositions == nil || *req.SystemJobPositions == "" || !regexp.MustCompile(`^(\d+)(,\d+)*$`).MatchString(*req.SystemJobPositions) {
		errList = append(errList, "职位格式错误")
	}
	if !regexp.MustCompile(`^(\d+)$`).MatchString(*req.SystemRole) {
		errList = append(errList, "角色格式错误")
	}
	return
}

// 修改用户状态请求
type SystemUserModifyStatusRequest struct {
	Id      uint   `form:"id" json:"id"`
	Operate string `form:"operate" json:"operate"`
}

// 批量修改用户状态请求
type SystemUserMultiModifyStatusRequest struct {
	Ids     []uint `form:"ids" json:"ids"`
	Operate string `form:"operate" json:"operate"`
}
