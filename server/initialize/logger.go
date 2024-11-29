package initialize

import (
	"fmt"
	"os"
	"stellar/common"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

// 日志日期格式调整
func ZapLocalTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(common.TimeMillisecondFormat))
}

// 日志初始化
func NewLogger(c common.LoggerConfiguration) *zap.SugaredLogger {
	config := zap.NewProductionEncoderConfig()       // 新建配置
	config.EncodeTime = ZapLocalTimeEncoder          // 调整时间
	config.EncodeLevel = zapcore.CapitalLevelEncoder // 关闭颜色
	var ws zapcore.WriteSyncer                       // 输出
	if c.Enabled {
		// 日志文件
		now := time.Now()
		filename := fmt.Sprintf("%s-%04d-%02d-%02d.log", c.Path, now.Year(), now.Month(), now.Day())

		// 日志切割规则
		hook := &lumberjack.Logger{
			Filename:   filename,
			MaxSize:    c.MaxSize,
			MaxAge:     c.MaxAge,
			MaxBackups: c.MaxBackups,
			Compress:   c.Compress,
		}

		// 延时关闭
		defer func(hook *lumberjack.Logger) {
			_ = hook.Close()
		}(hook)

		// 输出到控制台和文件
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(hook))
	} else {
		// 只输出到控制台
		ws = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout))
	}

	// 整合日志输出信息
	core := zapcore.NewCore(zapcore.NewConsoleEncoder(config), ws, c.Level)
	logger := zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
	return logger.Sugar()
}

// 系统日志初始化
func SystemLogger() {
	logger := NewLogger(common.Config.Log.System)
	common.SystemLog = logger
}

// 访问日志初始化
func AccessLogger() {
	logger := NewLogger(common.Config.Log.Access)
	common.AccessLog = logger
}
