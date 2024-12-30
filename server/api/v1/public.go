package v1

import (
	"net/http"
	"stellar/common"
	"stellar/dto"
	"stellar/pkg/gedis"
	"stellar/pkg/response"
	"strings"

	"github.com/gin-gonic/gin"
)

// 健康检查
func SystemHealthCheckHandler(ctx *gin.Context) {
	ctx.String(http.StatusOK, "OK")
}

// 系统信息
func SystemInformationHandler(ctx *gin.Context) {
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
func SystemVersionHandler(ctx *gin.Context) {
	response.SuccessWithData(map[string]interface{}{
		"systemVersion":   common.SystemVersion,
		"systemGoVersion": common.SystemGoVersion,
	})
}

// Token 校验
func TokenVerificationHandler(ctx *gin.Context) {
	response.Success()
}

// 节点信息
func NodeInformationHandler(ctx *gin.Context) {
	// 节点信息
	var nodes []dto.NodeInfoResponse

	// 连接 Redis
	conn := gedis.NewRedisConnection()
	// 查询所有以 HeartbeatId 开头的键
	keys, err := conn.Redis.Keys(ctx, common.RKP.HeartbeatId+":*").Result()
	if err != nil {
		response.FailedWithMessage("查询节点信息失败")
		return
	}

	// 遍历键，先获取所有节点信息
	for _, key := range keys {
		// 查询键的值
		startTime := conn.GetString(key).Unwrap()
		name := strings.TrimPrefix(key, common.RKP.HeartbeatId+":")
		nodes = append(nodes, dto.NodeInfoResponse{Name: name, StartTime: startTime})
	}

	// 获取 Leader 节点信息
	keys, err = conn.Redis.Keys(ctx, common.RKP.LeaderId+":*").Result()
	if err != nil {
		response.FailedWithMessage("查询 Leader 节点信息失败")
		return
	}

	// 遍历键，获取 Leader 节点信息
	for _, key := range keys {
		name := strings.TrimPrefix(key, common.RKP.LeaderId+":")
		// 修改指定节点的 Roles 信息
		for i, node := range nodes {
			if node.Name == name {
				nodes[i].Roles = append(nodes[i].Roles, "Leader")
			}
		}
	}

	// 获取 Worker 节点信息
	keys, err = conn.Redis.Keys(ctx, common.RKP.WorkerId+":*").Result()
	if err != nil {
		response.FailedWithMessage("查询 Worker 节点信息失败")
		return
	}

	// 遍历键，获取 Worker 节点信息
	for _, key := range keys {
		name := strings.TrimPrefix(key, common.RKP.WorkerId+":")
		// 修改指定节点的 Roles 信息
		for i, node := range nodes {
			if node.Name == name {
				nodes[i].Roles = append(nodes[i].Roles, "Worker")
			}
		}
	}

	// 获取 WebServer 服务信息
	keys, err = conn.Redis.Keys(ctx, common.RKP.WebServerId+":*").Result()
	if err != nil {
		response.FailedWithMessage("查询 WebServer 后端服务信息失败")
		return
	}

	// 遍历键，获取 Web 后端服务信息
	for _, key := range keys {
		name := strings.TrimPrefix(key, common.RKP.WebServerId+":")
		for i, node := range nodes {
			if node.Name == name {
				nodes[i].Roles = append(nodes[i].Roles, "WebServer")
			}
		}
	}

	response.SuccessWithData(nodes)
}
