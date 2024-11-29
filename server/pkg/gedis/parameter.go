package gedis

import (
	"fmt"
	"time"
)

// 参数处理
type RedisOperationParameter struct {
	Name  string
	Value interface{}
}

// 构造函数
func NewRedisOperationParameter(name string, value interface{}) *RedisOperationParameter {
	return &RedisOperationParameter{Name: name, Value: value}
}

// 多参数
type RedisOperationParameters []*RedisOperationParameter

// 查找用法参数
func (params RedisOperationParameters) Find(name string) *InterfaceResult {
	for _, param := range params {
		if param.Name == name {
			return NewInterfaceResult(param.Value, nil)
		}
	}
	return NewInterfaceResult(nil, fmt.Errorf("unsupported method: %s", name))
}

// 设置过期时间
func WithExpire(t time.Duration) *RedisOperationParameter {
	return &RedisOperationParameter{
		Name:  "expire",
		Value: t,
	}
}

// 设置 NX 锁，Key 不存在才能设置
func WithNX() *RedisOperationParameter {
	return &RedisOperationParameter{
		Name:  "nx",
		Value: struct{}{},
	}
}

// 设置 XX 锁，Key 存在才能设置
func WithXX() *RedisOperationParameter {
	return &RedisOperationParameter{
		Name:  "xx",
		Value: struct{}{},
	}
}
