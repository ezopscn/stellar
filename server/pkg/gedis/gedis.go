package gedis

import (
	"context"
	"stellar/common"
	"time"

	"github.com/redis/go-redis/v9"
)

// 用户操作
type RedisConnection struct {
	Redis   *redis.Client
	Context context.Context
}

// 构造函数
func NewRedisConnection() *RedisConnection {
	return &RedisConnection{Redis: common.RedisCache, Context: context.Background()}
}

// 获取单个值为 string 的 Key
func (c *RedisConnection) GetString(key string) *StringResult {
	return NewStringResult(c.Redis.Get(c.Context, key).Result())
}

// 获取单个值为 int 的 Key
func (c *RedisConnection) GetInt(key string) *IntResult {
	return NewIntResult(c.Redis.Get(c.Context, key).Int())
}

// 删除 Key
func (c *RedisConnection) Del(key string) (int64, error) {
	return c.Redis.Del(c.Context, key).Result()
}

// 设置 Key / Value
// 用法：gedis.Set("key", "value", gedis.WithExpire(time.Second * 10), gedis.WithNX())
func (c *RedisConnection) Set(key string, value interface{}, params ...*RedisOperationParameter) *InterfaceResult {
	// 参数列表
	ops := RedisOperationParameters(params)

	// 判断是否设置过期时间，没有则设置永不过期
	expire := ops.Find("expire").UnwrapWithDefaultValue(time.Second * 0).(time.Duration)

	// 判断是否 NX 锁，两种锁只能有一个
	nx := ops.Find("nx").UnwrapWithDefaultValue(nil)
	if nx != nil {
		return NewInterfaceResult(c.Redis.SetNX(c.Context, key, value, expire).Result())
	}

	// 判断是否 XX 锁
	xx := ops.Find("xx").UnwrapWithDefaultValue(nil)
	if xx != nil {
		return NewInterfaceResult(c.Redis.SetXX(c.Context, key, value, expire).Result())
	}

	// 默认
	return NewInterfaceResult(c.Redis.Set(c.Context, key, value, expire).Result())
}
