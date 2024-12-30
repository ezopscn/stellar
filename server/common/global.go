package common

import (
	"embed"

	"github.com/casbin/casbin/v2"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量，build 或者 run 的时候会通过参数替换
var (
	SystemVersion               = "unknown"                                       // 系统版本
	SystemGoVersion             = "unknown"                                       // 系统 Go 版本
	SystemConfigType            = "yaml"                                          // 配置文件类型
	SystemConfigDefaultFilename = "stellar.yaml"                                  // 默认配置文件名
	SystemConfigFilename        = ""                                              // 配置文件名，可以通过参数替换
	SystemConfigHost            = ""                                              // 配置文件主机，可以通过参数替换
	SystemConfigPort            = ""                                              // 配置文件端口，可以通过参数替换
	SystemRoleWebServer         = ""                                              // 是否是 Web 后端服务角色，可选值：1、0
	SystemRoleLeaderElection    = ""                                              // 是否是领导者竞选角色，可选值：1、0
	SystemRoleWorker            = ""                                              // 是否是工作者角色，可选值：1、0
	SystemRoleAdminList         = []string{"Administrator", "SuperAdministrator"} // 管理员角色关键字列表
)

// 全局状态
var (
	ClientId        *string // 客户端 ID
	ClientStartTime *string // 客户端启动时间
)

// 系统常量
const (
	SystemApiPrefix          = "/api/v1"                                                                                                                      // API 前缀
	SystemProjectName        = "Stellar"                                                                                                                      // 项目名称
	SystemProjectDescription = "Stellar is a multi-data-source operational monitoring system that integrates both system monitoring and business monitoring." // 项目描述
	SystemDeveloperName      = "DK"                                                                                                                           // 开发者
	SystemDeveloperEmail     = "ezopscn@gmail.com"                                                                                                            // 开发者邮箱
)

// 格式常量
const (
	TimeMillisecondFormat = "2006-01-02 15:04:05.000"    // 毫秒时间格式化
	TimeSecondFormat      = "2006-01-02 15:04:05"        // 秒时间格式化
	TimeDateFormat        = "2006-01-02"                 // 日期时间格式化
	UppercaseLetters      = "ABCDEFGHIJKLMNOPQRSTUVWXYZ" // 大写字母
	LowercaseLetters      = "abcdefghijklmnopqrstuvwxyz" // 小写字母
	Numbers               = "0123456789"                 // 数字
)

// 全局工具
var (
	FS             embed.FS           // 打包的静态资源，用于全局使用
	Config         *Configuration     // 配置文件解析
	SystemLog      *zap.SugaredLogger // 系统日志工具
	AccessLog      *zap.SugaredLogger // 访问日志工具
	MySQLDB        *gorm.DB           // 数据库连接
	RedisCache     *redis.Client      // 缓存连接
	CasbinEnforcer *casbin.Enforcer   // Casbin 策略执行器
)
