package cmd

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"stellar/common"
	"stellar/initialize"
	"stellar/pkg/gedis"
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
		initialize.ClientId()     // 初始化客户端 ID

		// 心跳上报
		go func() {
			cache := gedis.NewRedisConnection()
			key := fmt.Sprintf("%s:%s", common.RKP.HeartbeatId, *common.ClientId)
			for {
				common.SystemLog.Debugf("Start heartbeat check, client id: %s", *common.ClientId)
				cache.Set(key, time.Now().Unix(), gedis.WithExpire(time.Second*15))
				time.Sleep(10 * time.Second)
			}
		}()

		// 竞选 Leader
		if common.Config.System.LeaderElection {
			go func() {
				cache := gedis.NewRedisConnection()
				for {
					// 没有 Leader，则设置自己为 Leader，自己是 Leader，则延长过期时间并加锁
					leaderId := cache.GetString(common.RKP.LeaderId).Unwrap()
					if leaderId == "" {
						common.SystemLog.Debugf("No leader, set leader id: %s", *common.ClientId)
						cache.Set(common.RKP.LeaderId, *common.ClientId, gedis.WithExpire(time.Second*15))
					} else if leaderId == *common.ClientId {
						common.SystemLog.Debugf("I am leader, extend expiration time and lock")
						cache.Set(common.RKP.LeaderId, *common.ClientId, gedis.WithExpire(time.Second*15), gedis.WithXX())
					}
					time.Sleep(10 * time.Second)
				}
			}()
		}

		// 注册 Worker
		if common.Config.System.Worker {
			go func() {
				common.SystemLog.Debugf("Start register worker, client id: %s", *common.ClientId)
				cache := gedis.NewRedisConnection()
				key := fmt.Sprintf("%s:%s", common.RKP.WorkerId, *common.ClientId)
				for {
					cache.Set(key, time.Now().Unix(), gedis.WithExpire(time.Second*15))
					time.Sleep(10 * time.Second)
				}
			}()
		}

		// Web 后端服务
		if common.Config.System.WebServer {
			r := initialize.Router() // 初始化路由
			// 启动服务
			server := http.Server{
				Addr:    fmt.Sprintf("%s:%s", common.Config.System.Host, common.Config.System.Port),
				Handler: r,
			}

			go func() {
				common.SystemLog.Debugf("Start register web server, client id: %s", *common.ClientId)
				cache := gedis.NewRedisConnection()
				key := fmt.Sprintf("%s:%s", common.RKP.WebServerId, *common.ClientId)
				for {
					cache.Set(key, time.Now().Unix(), gedis.WithExpire(time.Second*15))
					time.Sleep(10 * time.Second)
				}
			}()

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
			fmt.Println("Server shutdown gracefully")
		} else {
			// 设置一个保活的主进程
			select {}
		}

		// 如果当前节点竞选成功 Leader，则启动任务调度
		// TODO: 启动任务调度
		// service.CheckAndPublishMetricTask()

		// time.Sleep(10 * time.Second)
		// service.SubscribeMetricTask()
	},
}
