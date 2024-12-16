package initialize

import (
	"stellar/common"

	"github.com/casbin/casbin/v2"
	casbinmodel "github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
)

// Casbin 初始化
func Casbin() {
	// 初始化数据库连接
	adapter, err := gormadapter.NewAdapterByDBUseTableName(common.MySQLDB, "", "system_casbin_rule")
	if err != nil {
		panic("Failed to initialize Casbin adapter: " + err.Error())
	}

	// 读取配置文件
	bs, err := common.FS.ReadFile("rbac.conf")
	if err != nil {
		panic("Failed to read Casbin configuration file: " + err.Error())
	}
	conf := string(bs[:])

	// 从字符串中加载配置
	m, err := casbinmodel.NewModelFromString(conf)
	if err != nil {
		panic("Failed to initialize Casbin model: " + err.Error())
	}

	// 初始化 Casbin
	enforcer, err := casbin.NewEnforcer(m, adapter)
	if err != nil {
		panic("Failed to initialize Casbin: " + err.Error())
	}

	// 加载策略
	err = enforcer.LoadPolicy()
	if err != nil {
		panic("Failed to load Casbin policy: " + err.Error())
	}

	// 设置全局 Casbin 策略执行器
	common.CasbinEnforcer = enforcer
}
