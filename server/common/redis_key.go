package common

// Redis Key Prefix
type RedisKeyPrefix struct {
	LoginToken      string // 用户登录 Token 前缀
	LoginWrongTimes string // 用户登录错误次数
	HeartbeatId     string // 心跳上报 ID
	WebServerId     string // Web 后端服务 ID
	LeaderId        string // 领导者 ID
	WorkerId        string // 工作者 ID
}

// 配置 Redis Key Prefix
var RKP = RedisKeyPrefix{
	LoginToken:      "Login:Token",
	LoginWrongTimes: "Login:WrongTimes",
	HeartbeatId:     "Heartbeat",
	WebServerId:     "WebServer",
	LeaderId:        "Leader",
	WorkerId:        "Worker",
}
