package dto

import "errors"

// 新增部门请求
type SystemDepartmentAddRequest struct {
	ParentId *uint   `form:"parentId" json:"parentId"`
	Name     *string `form:"name" json:"name"`
}

// 新增部门参数校验
func (req *SystemDepartmentAddRequest) Validate() error {
	// 检查名称是否为空或长度是否在3到30个字符之间
	if req.Name == nil || *req.Name == "" || len(*req.Name) < 3 || len(*req.Name) > 30 {
		return errors.New("名称不能为空，且长度必须在3到30个字符之间")
	}
	// 检查父级部门是否为空或是否为预留的未分配部门
	if req.ParentId == nil || *req.ParentId == 0 || *req.ParentId == 2 {
		return errors.New("父级部门不能为空，且不能为未分配部门")
	}
	return nil
}

// 修改部门请求
type SystemDepartmentUpdateRequest struct {
	Id       *uint   `form:"id" json:"id"`
	ParentId *uint   `form:"parentId" json:"parentId"`
	Name     *string `form:"name" json:"name"`
}

// 修改部门参数校验
func (req *SystemDepartmentUpdateRequest) Validate() error {
	// 检查ID是否为空
	if req.Id == nil || *req.Id == 0 {
		return errors.New("部门id不能为空")
	}
	// 检查名称是否为空或长度是否在3到30个字符之间
	if req.Name == nil || *req.Name == "" || len(*req.Name) < 3 || len(*req.Name) > 30 {
		return errors.New("名称不能为空，且长度必须在3到30个字符之间")
	}
	// 检查父级部门是否为空或是否为预留的未分配部门
	if req.ParentId == nil || *req.ParentId == 0 || *req.ParentId == 2 {
		return errors.New("父级部门不能为空，且不能为未分配部门")
	}
	return nil
}
