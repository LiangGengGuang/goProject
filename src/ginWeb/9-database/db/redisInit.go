package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"project/9-database/config"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

// RDB redis连接池
var RDB *redis.Client

func init() {

	//单连接
	addr := fmt.Sprintf("%s:%d", config.GlobalCfg.RedisCfg.Uri, config.GlobalCfg.RedisCfg.Port)
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: config.GlobalCfg.RedisCfg.UserName,
		Password: config.GlobalCfg.RedisCfg.Password,
		DB:       config.GlobalCfg.RedisCfg.DB,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		fmt.Printf("redis连接失败：%v", err)
		panic(err)
	}

	fmt.Println("redis连接成功")
}
