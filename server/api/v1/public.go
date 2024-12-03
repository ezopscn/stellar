package v1

import (
	"net/http"
	"stellar/common"
	"stellar/pkg/response"

	"github.com/gin-gonic/gin"
)

// 健康检查
func HealthHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 项目信息
func InfoHandler(ctx *gin.Context) {
	response.SuccessWithData(map[string]interface{}{
		"systemProjectName":        common.SystemProjectName,
		"systemProjectDescription": common.SystemProjectDescription,
		"systemVersion":            common.SystemVersion,
		"systemGoVersion":          common.SystemGoVersion,
		"systemDeveloperName":      common.SystemDeveloperName,
		"systemDeveloperEmail":     common.SystemDeveloperEmail,
	})
}

// 版本信息
func VersionHandler(ctx *gin.Context) {
	response.SuccessWithData(map[string]interface{}{
		"systemVersion":   common.SystemVersion,
		"systemGoVersion": common.SystemGoVersion,
	})
}

// Token 校验
func TokenVerificationHandler(ctx *gin.Context) {
	response.Success()
}
