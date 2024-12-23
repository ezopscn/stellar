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

}

// 批量添加用户接口
func SystemUserMutiAddHandler(ctx *gin.Context) {
	// 获取 post 参数
	req := []dto.SystemUserMutiAddRequest{}
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
	task := model.SystemUserMutiAddTask{
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
			taskDetail := model.SystemUserMutiAddDetail{
				TaskId:       task.Id,
				Username:     *v.Username,
				CNName:       *v.CNName,
				ENName:       *v.ENName,
				Email:        *v.Email,
				Phone:        *v.Phone,
				HidePhone:    *v.HidePhone,
				Gender:       *v.Gender,
				Departments:  *v.Departments,
				JobPositions: *v.JobPositions,
				Role:         *v.Role,
				Description:  *v.Description,
			}

			// 设置状态为进行中
			taskDetail.Status = trans.Uint(1)
			taskDetail.Result = "用户创建中"
			common.MySQLDB.Save(&taskDetail)

			// 错误列表
			errList := []string{}

			// 用户名
			if v.Username == nil || *v.Username == "" || len(*v.Username) > 30 || len(*v.Username) < 3 || !regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9_]*$`).MatchString(*v.Username) {
				errList = append(errList, "用户名格式错误")
			}

			// 密码
			if v.Password == nil || *v.Password == "" || len(*v.Password) > 30 || len(*v.Password) < 8 {
				errList = append(errList, "密码格式错误")
			}

			// 中文名
			if v.CNName == nil || *v.CNName == "" || len(*v.CNName) > 30 || len(*v.CNName) < 2 {
				errList = append(errList, "中文名格式错误")
			}

			// 英文名
			if v.ENName == nil || *v.ENName == "" || len(*v.ENName) > 30 || len(*v.ENName) < 2 || !regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(*v.ENName) {
				errList = append(errList, "英文名格式错误")
			}

			// 邮箱
			if v.Email == nil || *v.Email == "" || !utils.IsEmail(*v.Email) {
				errList = append(errList, "邮箱格式错误")
			}

			// 手机号
			if v.Phone == nil || *v.Phone == "" || !utils.IsPhoneNumber(*v.Phone) {
				errList = append(errList, "手机号格式错误")
			}

			// 隐藏手机号
			if v.HidePhone == nil || *v.HidePhone == "" || (*v.HidePhone != "0" && *v.HidePhone != "1") {
				errList = append(errList, "隐藏手机号格式错误，只能是 0 或 1")
			}

			// 性别
			if v.Gender == nil || *v.Gender == "" || (*v.Gender != "0" && *v.Gender != "1" && *v.Gender != "2") {
				errList = append(errList, "性别格式错误，只能是 0 或 1 或 2")
			}

			// 部门
			if v.Departments == nil || *v.Departments == "" || !regexp.MustCompile(`^(\d+)(,\d+)*$`).MatchString(*v.Departments) {
				errList = append(errList, "部门格式错误")
			}

			// 职位
			if v.JobPositions == nil || *v.JobPositions == "" || !regexp.MustCompile(`^(\d+)(,\d+)*$`).MatchString(*v.JobPositions) {
				errList = append(errList, "职位格式错误")
			}

			// 角色
			if v.Role == nil || *v.Role == "" || !regexp.MustCompile(`^(\d+)$`).MatchString(*v.Role) {
				errList = append(errList, "角色格式错误")
			}

			// 描述
			if len(*v.Description) > 200 {
				errList = append(errList, "描述长度不能超过 200 个字符")
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
						departments := strings.Split(*v.Departments, ",")
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
						jobPositions := strings.Split(*v.JobPositions, ",")
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
					SystemRoleId: utils.StringToUint(*v.Role),
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
				common.MySQLDB.Model(&model.SystemUserMutiAddDetail{}).Where("taskId = ?", task.Id).Where("status = ?", trans.Uint(2)).Count(&successCount)
				// 查询失败的数量
				common.MySQLDB.Model(&model.SystemUserMutiAddDetail{}).Where("taskId = ?", task.Id).Where("status = ?", trans.Uint(3)).Count(&failCount)
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
func SystemUserStatusModifyHandler(ctx *gin.Context) {
	// 获取 post 参数
	req := dto.SystemUserStatusModifyRequest{}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.FailedWithMessage(err.Error())
		return
	}

	// 默认 ID 为 1 的用户不能修改状态，属于系统预设用户
	if utils.IsUintInSlice(1, req.Ids) {
		response.FailedWithMessage("系统预设用户不能修改状态")
		return
	}

	// 判断操作类型
	if req.Operate == "disable" {
		if err := common.MySQLDB.Model(&model.SystemUser{}).Where("id IN (?)", req.Ids).Update("status", trans.Uint(0)).Error; err != nil {
			response.FailedWithMessage("禁用用户失败")
			return
		}
	} else if req.Operate == "enable" {
		if err := common.MySQLDB.Model(&model.SystemUser{}).Where("id IN (?)", req.Ids).Update("status", trans.Uint(1)).Error; err != nil {
			response.FailedWithMessage("启用用户失败")
			return
		}
	} else {
		response.FailedWithMessage("不支持的操作类型")
		return
	}
	response.Success()
}
