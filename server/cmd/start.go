package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"stellar/common"
	"stellar/initialize"
	"stellar/service"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().StringVarP(&common.SystemConfigHost, "host", "", common.SystemConfigHost, "The host to run the server on")
	startCmd.Flags().StringVarP(&common.SystemConfigPort, "port", "", common.SystemConfigPort, "The port to run the server on")
	startCmd.Flags().StringVarP(&common.SystemConfigFilename, "config", "", common.SystemConfigFilename, "The path to the configuration file")
	startCmd.Flags().StringVarP(&common.SystemRoleWebServer, "web-server", "", common.SystemRoleWebServer, "Is it a web server role, optional values: 1 (yes), 0 (no)")
	startCmd.Flags().StringVarP(&common.SystemRoleLeaderElection, "leader-election", "", common.SystemRoleLeaderElection, "Is it a leader election role, optional values: 1 (yes), 0 (no)")
	startCmd.Flags().StringVarP(&common.SystemRoleWorker, "worker", "", common.SystemRoleWorker, "Is it a worker role, optional values: 1 (yes), 0 (no)")
}

// 启动命令
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Print(common.LOGO) // 打印 logo
		initialize.Config()    // 初始化配置

		// 如果所有角色都没有设置，则直接退出
		if !common.Config.System.LeaderElection && !common.Config.System.Worker && !common.Config.System.WebServer {
			fmt.Println("No server role is set, exit")
			return
		}

		initialize.SystemLogger() // 初始化日志
		initialize.AccessLogger() // 初始化访问日志
		initialize.MySQL()        // 初始化 MySQL
		initialize.Redis()        // 初始化 Redis
		initialize.Casbin()       // 初始化 Casbin
		initialize.ClientId()     // 初始化客户端 ID

		go func() {
			// 心跳上报
			service.ReportHeartbeatTask()
		}()

		if common.Config.System.LeaderElection {
			go func() {
				// 竞选 Leader
				service.ElectionLeaderTask()
			}()

			go func() {
				// Leader 读取并发布任务到 Redis 列表中
				service.PublishTaskToRedisListTask()
			}()
		}

		if common.Config.System.Worker {
			go func() {
				// 注册 Worker
				service.RegisterWorkerTask()
			}()

			go func() {
				// Worker Redis 列表中获取任务
				service.ConsumeTaskFromRedisListTask()
			}()
		}

		// Web 后端服务
		if common.Config.System.WebServer {
			// 初始化路由
			r := initialize.Router()
			server := http.Server{
				Addr:    fmt.Sprintf("%s:%s", common.Config.System.Host, common.Config.System.Port),
				Handler: r,
			}

			// 注册 Web 服务
			go func() {
				service.RegisterWebServerTask()
			}()

			// 启动服务
			go func() {
				err := server.ListenAndServe()
				if err != nil && err != http.ErrServerClosed {
					panic("Failed to start the server: " + err.Error())
				}
			}()

			// 监听信号
			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit

			// 优雅关闭
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()
			err := server.Shutdown(ctx)
			if err != nil {
				panic("Failed to shutdown the server: " + err.Error())
			}
			fmt.Println("Server shutdown successfully")
		} else {
			select {} // 设置一个保活的主进程
		}
	},
}
