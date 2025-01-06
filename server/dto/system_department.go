package dto

// 新增部门请求
type SystemDepartmentAddRequest struct {
	ParentId *uint   `form:"parentId" json:"parentId"`
	Name     *string `form:"name" json:"name"`
}
