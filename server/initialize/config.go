package initialize

import (
	"bytes"
	"fmt"
	"os"
	"stellar/common"
	"stellar/pkg/utils"

	"github.com/spf13/viper"
)

// 配置初始化
func Config() {
	// 读取配置文件
	var bs []byte
	var err error
	v := viper.New()
	v.SetConfigType(common.SystemConfigType)
	if common.SystemConfigFilename != "" {
		filename := common.SystemConfigFilename
		fmt.Println("Start loading the specified configuration file:", filename)
		bs, err = os.ReadFile(filename)
	} else {
		filename := common.SystemConfigDefaultFilename
		fmt.Println("Start loading the default configuration file:", filename)
		bs, err = common.FS.ReadFile(filename)
	}

	if err != nil {
		panic("Failed to read the configuration file: " + err.Error())
	}

	// 解析配置
	err = v.ReadConfig(bytes.NewReader(bs))
	if err != nil {
		panic("Failed to parse the configuration file: " + err.Error())
	}

	// 反序列化配置
	err = v.Unmarshal(&common.Config)
	if err != nil {
		panic("Configuration file deserialization failed: " + err.Error())
	}

	// 命令行参数解析覆盖配置文件中的设置
	// 监听地址
	if common.SystemConfigHost != "" {
		if !utils.IsIPv4(common.SystemConfigHost) {
			panic("Invalid host: " + common.SystemConfigHost)
		}
		common.Config.System.Host = common.SystemConfigHost
	}

	// 监听端口
	if common.SystemConfigPort != "" {
		if !utils.IsPort(common.SystemConfigPort) {
			panic("Invalid port: " + common.SystemConfigPort)
		}
		common.Config.System.Port = common.SystemConfigPort
	}

	// WebServer 角色
	if common.SystemRoleWebServer != "" {
		if common.SystemRoleWebServer != "1" && common.SystemRoleWebServer != "0" {
			panic("Invalid web server role: " + common.SystemRoleWebServer)
		}
		if common.SystemRoleWebServer == "1" {
			common.Config.System.WebServer = true
		} else {
			common.Config.System.WebServer = false
		}
	}

	// LeaderElection 角色
	if common.SystemRoleLeaderElection != "" {
		if common.SystemRoleLeaderElection != "1" && common.SystemRoleLeaderElection != "0" {
			panic("Invalid leader election role: " + common.SystemRoleLeaderElection)
		}
		if common.SystemRoleLeaderElection == "1" {
			common.Config.System.LeaderElection = true
		} else {
			common.Config.System.LeaderElection = false
		}
	}

	// Worker 角色
	if common.SystemRoleWorker != "" {
		if common.SystemRoleWorker != "1" && common.SystemRoleWorker != "0" {
			panic("Invalid worker role: " + common.SystemRoleWorker)
		}
		if common.SystemRoleWorker == "1" {
			common.Config.System.Worker = true
		} else {
			common.Config.System.Worker = false
		}
	}
}
