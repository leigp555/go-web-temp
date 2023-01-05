package database

import (
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

var (
	rdb *redis.Client
	ctx = context.Background()
)

// 初始化连接
func initClient() (err error) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "1.117.141.66:6379",
		Password: "123456abc", // no password set
		DB:       0,           // use default DB
		PoolSize: 100,         // 连接池大小
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdb.Ping(ctx).Result()
	return err
}
