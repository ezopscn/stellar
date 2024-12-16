package middleware

import (
	"stellar/common"
	"stellar/pkg/response"
	"stellar/pkg/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

// Casbin 中间件
func Casbin(ctx *gin.Context) {
	// sub: 用户角色关键字
	sub, _ := utils.ExtractStringResultFromContext(ctx, "systemRoleKeyword")
	// obj: 请求路径
	obj := strings.TrimPrefix(ctx.Request.RequestURI, common.SystemApiPrefix)
	// act: 请求方法
	act := ctx.Request.Method
	// 执行 Casbin 策略
	pass, _ := common.CasbinEnforcer.Enforce(sub, obj, act)
	// 如果策略不通过，则返回 403 状态码
	if !pass {
		response.FailedWithCode(response.RequestForbidden)
		ctx.Abort()
		return
	}
	ctx.Next()
}
