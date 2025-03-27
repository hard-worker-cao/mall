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
func initRedis() {
	//加载conf文件
	config, err1 := ini.Load("conf/app.ini")
	if err1 != nil {
		fmt.Println(err1)
		os.Exit(1)
	}
	redisIP := config.Section("redis").Key("ip").String()
	redisPort := config.Section("redis").Key("port").String()
	RedisDB = redis.NewClient(&redis.Options{
		Addr:     redisIP + ":" + redisPort,
		Password: "",
		DB:       0,
	})

	pong, err2 := RedisDB.Ping(redisCoretxt).Result()

	if err2 != nil {
		fmt.Println(err2)
	}

	fmt.Println(pong)
	//在redis存储中完成set get add等操作
}
