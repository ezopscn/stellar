package common

import "go.uber.org/zap/zapcore"

// 配置文件结构体
type Configuration struct {
	System SystemConfiguration `mapstructure:"system" json:"system"`
	Log    LogConfiguration    `mapstructure:"log" json:"log"`
	MySQL  MySQLConfiguration  `mapstructure:"mysql" json:"mysql"`
	Redis  RedisConfiguration  `mapstructure:"redis" json:"redis"`
	JWT    JWTConfiguration    `mapstructure:"jwt" json:"jwt"`
	Login  LoginConfiguration  `mapstructure:"login" json:"login"`
}

// 系统配置
type SystemConfiguration struct {
	Host           string `mapstructure:"host" json:"host"`
	Port           string `mapstructure:"port" json:"port"`
	WebServer      bool   `mapstructure:"web-server" json:"webServer"`
	LeaderElection bool   `mapstructure:"leader-election" json:"leaderElection"`
	Worker         bool   `mapstructure:"worker" json:"worker"`
}

// 日志配置
type LogConfiguration struct {
	System LoggerConfiguration `mapstructure:"system" json:"system"`
	Access LoggerConfiguration `mapstructure:"access" json:"access"`
}

// 日志类型配置
type LoggerConfiguration struct {
	Enabled    bool          `mapstructure:"enabled" json:"enabled"`
	Level      zapcore.Level `mapstructure:"level" json:"level"`
	Path       string        `mapstructure:"path" json:"path"`
	MaxSize    int           `mapstructure:"max-size" json:"MaxSize"`
	MaxAge     int           `mapstructure:"max-age" json:"maxAge"`
	MaxBackups int           `mapstructure:"max-backups" json:"maxBackups"`
	Compress   bool          `mapstructure:"compress" json:"compress"`
}

// 数据库配置
type MySQLConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     string `mapstructure:"database" json:"database"`
	Username     string `mapstructure:"username" json:"username"`
	Password     string `mapstructure:"password" json:"password"`
	Charset      string `mapstructure:"charset" json:"charset"`
	Collation    string `mapstructure:"collation" json:"collation"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
	ExtraParam   string `mapstructure:"extra-param" json:"extraParam"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"maxIdleTime"`
}

// Redis 配置
type RedisConfiguration struct {
	Host         string `mapstructure:"host" json:"host"`
	Port         int    `mapstructure:"port" json:"port"`
	Database     int    `mapstructure:"database" json:"database"`
	Password     string `mapstructure:"password" json:"password"`
	Timeout      int    `mapstructure:"timeout" json:"timeout"`
	MaxOpenConns int    `mapstructure:"max-open-conns" json:"maxOpenConns"`
	MinIdleConns int    `mapstructure:"min-idle-conns" json:"minIdleConns"`
	MaxIdleConns int    `mapstructure:"max-idle-conns" json:"maxIdleConns"`
	MaxIdleTime  int    `mapstructure:"max-idle-time" json:"maxIdleTime"`
}

// JWT 配置
type JWTConfiguration struct {
	Realm   string `mapstructure:"realm" json:"realm"`
	Key     string `mapstructure:"key" json:"key"`
	Timeout int    `mapstructure:"timeout" json:"timeout"`
}

// 登录配置
type LoginConfiguration struct {
	WrongTimes     int  `mapstructure:"wrong-times" json:"wrongTimes"`
	LockTime       int  `mapstructure:"lock-time" json:"lockTime"`
	MultiDevices   bool `mapstructure:"multi-devices" json:"multiDevices"`
	ResetTokenTime int  `mapstructure:"reset-token-time" json:"resetTokenTime"`
}
