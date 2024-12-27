package v1

import (
	"fmt"
	"regexp"
	"stellar/common"
	"stellar/dto"
	"stellar/initialize/data"
	"stellar/model"
	"stellar/pkg/response"
	"stellar/pkg/trans"
	"stellar/pkg/utils"
	"stellar/service"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// 用户列表接口
func SystemUserListHandler(ctx *gin.Context) {
	// 获取当前用户的角色，如果是管理员以上的级别，则可以看到隐藏的手机号
	systemRoleKeyword, err := utils.ExtractStringResultFromContext(ctx, "systemRoleKeyword")
	if err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 获取用户列表
	users, pagination, err := service.GetSystemUserListService(ctx)
	if err != nil {
		response.FailedWithMessage("获取用户列表失败")
		return
	}

	// 判断角色是否在管理员列表，如果不是管理员，则隐藏设置了隐藏标识的手机号
	if !utils.IsStringInSlice(systemRoleKeyword, common.SystemRoleAdminList) {
		for i := range users {
			if *users[i].HidePhone == 1 {
				users[i].Phone = utils.HidePhoneNumber(users[i].Phone)
			}
		}
	}

	response.SuccessWithData(dto.PaginationResponse{
		List:       users,
		Pagination: pagination,
	})
}

// 添加用户接口
func SystemUserAddHandler(ctx *gin.Context) {
	// 获取 post 参数
	req := dto.SystemUserAddRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 校验提交数据
	errList := []string{}
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
	if *req.HidePhone != 0 && *req.HidePhone != 1 {
		errList = append(errList, "隐藏手机号格式错误，只能是 0 或 1")
	}
	if *req.Gender != 1 && *req.Gender != 2 && *req.Gender != 3 {
		errList = append(errList, "性别格式错误，只能是 1 或 2 或 3")
	}
	if len(req.SystemDepartments) == 0 {
		errList = append(errList, "部门不能为空")
	}
	if len(req.SystemJobPositions) == 0 {
		errList = append(errList, "职位不能为空")
	}
	if req.SystemRole == nil || *req.SystemRole == 0 {
		errList = append(errList, "角色不能为空")
	}
	if req.Description != nil {
		if len(*req.Description) > 200 {
			errList = append(errList, "描述长度不能超过 200 个字符")
		}
	} else {
		req.Description = trans.String("")
	}

	// 如果错误列表不为空，则返回错误
	if len(errList) > 0 {
		response.FailedWithMessage(strings.Join(errList, ","))
		return
	}
	fmt.Println(1)

	// 用户模型
	user := model.SystemUser{
		Username:  *req.Username,
		Password:  utils.CryptoPassword(*req.Password),
		CNName:    *req.CNName,
		ENName:    *req.ENName,
		Email:     *req.Email,
		Phone:     *req.Phone,
		HidePhone: trans.Uint(*req.HidePhone),
		Gender:    trans.Uint(*req.Gender),
		Avatar: func() string {
			if *req.Gender == 1 {
				return data.RandomMaleAvatar()
			}
			return data.RandomFemaleAvatar()
		}(),
		SystemDepartments: func() []model.SystemDepartment {
			var result []model.SystemDepartment
			for _, deptId := range req.SystemDepartments {
				result = append(result, model.SystemDepartment{
					BaseModel: model.BaseModel{Id: deptId},
				})
			}
			return result
		}(),
		SystemJobPositions: func() []model.SystemJobPosition {
			var result []model.SystemJobPosition
			for _, posId := range req.SystemJobPositions {
				result = append(result, model.SystemJobPosition{
					BaseModel: model.BaseModel{Id: posId},
				})
			}
			return result
		}(),
		SystemRoleId: uint(*req.SystemRole),
		Description:  *req.Description,
		Creator: func() string {
			username, _ := utils.ExtractStringResultFromContext(ctx, "username")
			cnName, _ := utils.ExtractStringResultFromContext(ctx, "cnName")
			enName, _ := utils.ExtractStringResultFromContext(ctx, "enName")
			userId, _ := utils.ExtractUintResultFromContext(ctx, "userId")
			return fmt.Sprintf("%s,%s,%s,%d", username, cnName, enName, userId)
		}(),
	}

	// 创建用户
	if err := common.MySQLDB.Create(&user).Error; err != nil {
		response.FailedWithMessage("创建用户失败，" + err.Error())
		return
	}
	response.Success()
}

// 批量添加用户接口
func SystemUserMultiAddHandler(ctx *gin.Context) {
	// 获取 post 参数
	req := []dto.SystemUserMultiAddRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 判断是否为空
	if len(req) == 0 {
		response.FailedWithMessage("未获取到导入数据，请检测文件格式是否正确")
		return
	}

	// 创建任务
	task := model.SystemUserMultiAddTask{
		CreatorId: func() uint {
			userId, _ := utils.ExtractStringResultFromContext(ctx, "userId")
			return utils.StringToUint(userId)
		}(),
		UserNumber: uint(len(req)),
	}
	if err := common.MySQLDB.Create(&task).Error; err != nil {
		response.FailedWithMessage("创建批量导入任务失败")
		return
	}

	// 遍历数据进行校验
	for idx, v := range req {
		go func(i int) {
			// 任务详情
			taskDetail := model.SystemUserMultiAddDetail{
				TaskId:             task.Id,
				Username:           *v.Username,
				CNName:             *v.CNName,
				ENName:             *v.ENName,
				Email:              *v.Email,
				Phone:              *v.Phone,
				HidePhone:          *v.HidePhone,
				Gender:             *v.Gender,
				SystemDepartments:  *v.SystemDepartments,
				SystemJobPositions: *v.SystemJobPositions,
				SystemRole:         *v.SystemRole,
				Description:        *v.Description,
			}

			// 设置状态为进行中
			taskDetail.Status = trans.Uint(1)
			taskDetail.Result = "用户创建中"
			common.MySQLDB.Save(&taskDetail)

			// 校验提交数据
			errList := []string{}
			if !utils.IsUsername(*v.Username) {
				errList = append(errList, "用户名格式错误")
			}
			if !utils.IsPassword(*v.Password) {
				errList = append(errList, "密码格式错误")
			}
			if !utils.IsCNName(*v.CNName) {
				errList = append(errList, "中文名格式错误")
			}
			if !utils.IsENName(*v.ENName) {
				errList = append(errList, "英文名格式错误")
			}
			if !utils.IsEmail(*v.Email) {
				errList = append(errList, "邮箱格式错误")
			}
			if !utils.IsPhoneNumber(*v.Phone) {
				errList = append(errList, "手机号格式错误")
			}
			if *v.HidePhone != "0" && *v.HidePhone != "1" {
				errList = append(errList, "隐藏手机号格式错误，只能是 0 或 1")
			}
			if *v.Gender != "1" && *v.Gender != "2" && *v.Gender != "3" {
				errList = append(errList, "性别格式错误，只能是 1 或 2 或 3")
			}
			if v.SystemDepartments == nil || *v.SystemDepartments == "" || !regexp.MustCompile(`^(\d+)(,\d+)*$`).MatchString(*v.SystemDepartments) {
				errList = append(errList, "部门格式错误")
			}
			if v.SystemJobPositions == nil || *v.SystemJobPositions == "" || !regexp.MustCompile(`^(\d+)(,\d+)*$`).MatchString(*v.SystemJobPositions) {
				errList = append(errList, "职位格式错误")
			}
			if !regexp.MustCompile(`^(\d+)$`).MatchString(*v.SystemRole) {
				errList = append(errList, "角色格式错误")
			}
			if v.Description != nil {
				if len(*v.Description) > 200 {
					errList = append(errList, "描述长度不能超过 200 个字符")
				}
			} else {
				v.Description = trans.String("")
			}

			// 如果错误列表不为空，则返回错误
			if len(errList) > 0 {
				// 更新状态和原因
				common.MySQLDB.Model(&taskDetail).Where("id = ?", taskDetail.Id).Updates(map[string]interface{}{
					"status": trans.Uint(3),
					"result": strings.Join(errList, ","),
				})
			} else {
				// 转换成用户模型
				user := model.SystemUser{
					Username:  *v.Username,
					Password:  utils.CryptoPassword(*v.Password),
					CNName:    *v.CNName,
					ENName:    *v.ENName,
					Email:     *v.Email,
					Phone:     *v.Phone,
					HidePhone: trans.Uint(utils.StringToUint(*v.HidePhone)),
					Gender:    trans.Uint(utils.StringToUint(*v.Gender)),
					Avatar: func() string {
						if utils.StringToUint(*v.Gender) == 1 {
							return data.RandomMaleAvatar()
						}
						return data.RandomFemaleAvatar()
					}(),
					SystemDepartments: func() []model.SystemDepartment {
						departments := strings.Split(*v.SystemDepartments, ",")
						var result []model.SystemDepartment
						for _, deptId := range departments {
							result = append(result, model.SystemDepartment{
								BaseModel: model.BaseModel{
									Id: utils.StringToUint(deptId),
								},
							})
						}
						return result
					}(),
					SystemJobPositions: func() []model.SystemJobPosition {
						jobPositions := strings.Split(*v.SystemJobPositions, ",")
						var result []model.SystemJobPosition
						for _, posId := range jobPositions {
							result = append(result, model.SystemJobPosition{
								BaseModel: model.BaseModel{
									Id: utils.StringToUint(posId),
								},
							})
						}
						return result
					}(),
					SystemRoleId: utils.StringToUint(*v.SystemRole),
					Description:  *v.Description,
					Creator: func() string {
						username, _ := utils.ExtractStringResultFromContext(ctx, "username")
						cnName, _ := utils.ExtractStringResultFromContext(ctx, "cnName")
						enName, _ := utils.ExtractStringResultFromContext(ctx, "enName")
						userId, _ := utils.ExtractUintResultFromContext(ctx, "userId")
						return fmt.Sprintf("%s,%s,%s,%d", username, cnName, enName, userId)
					}(),
				}

				// 创建用户
				common.SystemLog.Info("开始创建用户：", user.Username)
				if err := common.MySQLDB.Create(&user).Error; err != nil {
					common.MySQLDB.Model(&taskDetail).Where("id = ?", taskDetail.Id).Updates(map[string]interface{}{
						"status": trans.Uint(3),
						"result": err.Error(),
					})
					common.SystemLog.Error("创建用户失败：", err.Error())
				} else {
					// 更新状态和结果
					common.MySQLDB.Model(&taskDetail).Where("id = ?", taskDetail.Id).Updates(map[string]interface{}{
						"status": trans.Uint(2),
						"result": "创建成功",
					})
				}
			}

			// 判断是不是列表最后一个元素
			if i == len(req)-1 {
				time.Sleep(10 * time.Second) // 避免还没有执行完成，就返回结果
				var successCount, failCount int64
				// 查询成功的数量
				common.MySQLDB.Model(&model.SystemUserMultiAddDetail{}).Where("taskId = ?", task.Id).Where("status = ?", trans.Uint(2)).Count(&successCount)
				// 查询失败的数量
				common.MySQLDB.Model(&model.SystemUserMultiAddDetail{}).Where("taskId = ?", task.Id).Where("status = ?", trans.Uint(3)).Count(&failCount)
				// 更新任务状态
				common.MySQLDB.Model(&task).Where("id = ?", task.Id).Updates(map[string]interface{}{
					"successNumber": uint(successCount),
					"failNumber":    uint(failCount),
					"status":        trans.Uint(2),
				})
				common.SystemLog.Info(fmt.Sprintf("批量导入任务完成，成功：%d，失败：%d", successCount, failCount))
			}
		}(idx)
	}
	response.Success()
}

// 修改用户状态
func SystemUserModifyStatusHandler(ctx *gin.Context) {
	// 获取 post 参数
	req := dto.SystemUserModifyStatusRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 单个修改
	ids := []uint{req.Id}
	if err := service.SystemUserModifyStatusService(ctx, ids, req.Operate); err != nil {
		response.FailedWithMessage("修改用户状态失败")
		return
	}
	response.Success()
}

// 批量修改用户状态
func SystemUserMultiModifyStatusHandler(ctx *gin.Context) {
	// 获取 post 参数
	req := dto.SystemUserMultiModifyStatusRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 批量修改
	if err := service.SystemUserModifyStatusService(ctx, req.Ids, req.Operate); err != nil {
		response.FailedWithMessage("修改用户状态失败")
		return
	}
	response.Success()
}
