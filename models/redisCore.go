package models

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gopkg.in/ini.v1"
	"os"
)

var redisCoretxt = context.Background()

var (
	RedisDB *redis.Client
)

// 配置连接
func init() {
	//加载conf文件
	config, err := ini.Load("conf/redis.ini")
	if err != nil {
		ErrHandler(err)
		os.Exit(1)
	}
	redisIP := config.Section("redis").Key("ip").String()
	redisPort := config.Section("redis").Key("port").String()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":" + redisPort,
		Password: "",
		DB:       0,
	})

	pong, err := RedisDB.Ping(redisCoretxt).Result()

	if err != nil {
		ErrHandler(err)
	}

	fmt.Println(pong)
	//在redis存储中完成set get add等操作
}
