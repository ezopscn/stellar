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
		cache.Set(key, time.Now().Unix(), gedis.WithExpire(time.Second*15))
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
				// 发布到消息队列，确保只有一个订阅者能收到消息
				name := metricTask.MetricName + metricTask.MetricLabel
				common.SystemLog.Debugf("Detect metric task: %s", name)
				err := common.RedisCache.LPush(context.Background(), common.RKP.MetricTask, metricTask.Id).Err()
				if err != nil {
					common.SystemLog.Error("Publish metric task failed: ", name, ": ", err.Error())
					continue
				}

				// 更新任务下次执行时间
			}
		}
		time.Sleep(1 * time.Second)
	}
}

// Worker 订阅任务
func SubscribeTaskFromRedisListTask() {
	common.SystemLog.Debugf("Start subscribe task, client id: %s", *common.ClientId)
	for {
		common.SystemLog.Debugf("I am worker, subscribe task, client id: %s", *common.ClientId)
		taskId, err := common.RedisCache.BRPop(context.Background(), 0, common.RKP.MetricTask).Result()
		if err == nil {
			common.SystemLog.Debugf("Get metric task: %s", taskId[1])
		}
		time.Sleep(1 * time.Second)
	}
}
