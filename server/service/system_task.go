package service

import (
	"context"
	"fmt"
	"stellar/common"
	"stellar/model"
	"stellar/pkg/gedis"
	"time"
)

// 心跳上报任务
func ReportHeartbeatTask() {
	cache := gedis.NewRedisConnection()
	key := fmt.Sprintf("%s:%s", common.RKP.HeartbeatId, *common.ClientId)
	for {
		common.SystemLog.Infof("Start heartbeat check, client id: %s", *common.ClientId)
		cache.Set(key, *common.ClientStartTime, gedis.WithExpire(time.Second*15))
		time.Sleep(10 * time.Second)
	}
}

// 竞选 Leader 任务
func ElectionLeaderTask() {
	cache := gedis.NewRedisConnection()
	for {
		// 没有 Leader，则设置自己为 Leader，自己是 Leader，则延长过期时间并加锁
		leaderId := cache.GetString(common.RKP.LeaderId).Unwrap()
		if leaderId == "" {
			common.SystemLog.Infof("No leader, set leader id: %s", *common.ClientId)
			cache.Set(common.RKP.LeaderId, *common.ClientId, gedis.WithExpire(time.Second*15))
		} else if leaderId == *common.ClientId {
			common.SystemLog.Debugf("I am leader, extend expiration time and lock")
			cache.Set(common.RKP.LeaderId, *common.ClientId, gedis.WithExpire(time.Second*15), gedis.WithXX())
		}
		time.Sleep(10 * time.Second)
	}
}

// Worker 注册任务
func RegisterWorkerTask() {
	common.SystemLog.Infof("Start register worker, client id: %s", *common.ClientId)
	cache := gedis.NewRedisConnection()
	key := fmt.Sprintf("%s:%s", common.RKP.WorkerId, *common.ClientId)
	for {
		cache.Set(key, time.Now().Unix(), gedis.WithExpire(time.Second*15))
		time.Sleep(10 * time.Second)
	}
}

// Web 服务注册任务
func RegisterWebServerTask() {
	common.SystemLog.Infof("Start register web server, client id: %s", *common.ClientId)
	cache := gedis.NewRedisConnection()
	key := fmt.Sprintf("%s:%s", common.RKP.WebServerId, *common.ClientId)
	for {
		cache.Set(key, time.Now().Unix(), gedis.WithExpire(time.Second*15))
		time.Sleep(10 * time.Second)
	}
}

// 检测并发布任务
func PublishTaskToRedisListTask() {
	cache := gedis.NewRedisConnection()
	for {
		leaderId := cache.GetString(common.RKP.LeaderId).Unwrap()
		if leaderId == *common.ClientId {
			common.SystemLog.Debugf("I am leader, start reading tasks and publishing metrics")
			// 如果是 Leader，则开始读取任务
			var metricTasks []model.MetricTask
			now := time.Now()
			err := common.MySQLDB.Where("nextRunTime <= ? AND status = ?", now, 1).Find(&metricTasks).Error
			if err != nil {
				common.SystemLog.Error("Get metric tasks failed: ", err.Error())
				return
			}

			for _, metricTask := range metricTasks {
				// 发布到消息列表，这样才能确保只有一个消费者能收到消息
				name := metricTask.MetricName + metricTask.MetricLabel
				common.SystemLog.Debugf("Publish metric task: %s", name)
				err := common.RedisCache.LPush(context.Background(), common.RKP.MetricTask, metricTask.Id).Err()
				if err != nil {
					common.SystemLog.Error("Publish metric task failed: ", name, ": ", err.Error())
					continue
				}

				// 更新任务下次执行时间，解析 Cron 表达式
				// parser := cron.NewParser(cron.Second | cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow)
				// nextRunTime, err := parser.Parse(metricTask.CronExpression)
				// if err != nil {
				// 	common.SystemLog.Error("Parse cron expression failed: ", err.Error())
				// 	continue
				// }
				// metricTask.NextRunTime = nextRunTime.Next(now)
				// err = common.MySQLDB.Save(&metricTask).Error
				// if err != nil {
				// 	common.SystemLog.Error("Update metric task next run time failed: ", err.Error())
				// 	continue
				// }
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// Worker 消费任务
func ConsumeTaskFromRedisListTask() {
	common.SystemLog.Debugf("Start consume task, client id: %s", *common.ClientId)
	for {
		common.SystemLog.Debugf("I am worker, consume task, client id: %s", *common.ClientId)
		taskId, err := common.RedisCache.BRPop(context.Background(), 0, common.RKP.MetricTask).Result()
		if err == nil {
			// 查询任务详情
			var metricTask model.MetricTask
			err := common.MySQLDB.Where("id = ?", taskId[1]).First(&metricTask).Error
			if err != nil {
				common.SystemLog.Error("Get metric task detail failed: ", err.Error())
				continue
			}

			// 开始执行任务
			common.SystemLog.Debugf("Start execute metric task: %s", metricTask.Name)

			// 异步执行任务，如果任务没有在规定实际执行完成，则终止任务，并保存执行记录
		}
	}
}
