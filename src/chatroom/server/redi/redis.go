package redi

import (
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func InitRedisPool(address string, password string, db, maxIdle, maxActive int, idleTimeout time.Duration) {
	Client = redis.NewClient(&redis.Options{
		Addr:         address,     // Redis 服务器地址
		Password:     password,    // Redis 访问密码，如果无密码可为空字符串""
		DB:           db,          // Redis 数据库索引，默认为 0
		PoolSize:     maxIdle,     // 最大空闲连接数
		PoolTimeout:  idleTimeout, // 连接池超时时间
		MinIdleConns: 0,           // 最小空闲连接数
	})
	fmt.Println("连接池已经启动")
}
