package db

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"project/9-database/logger"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

// RDB redis连接池
var RDB *redis.Client

func init() {

	//单连接
	addr := fmt.Sprintf("%s:%d", GlobalCfg.RedisCfg.Uri, GlobalCfg.RedisCfg.Port)
	RDB = redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: GlobalCfg.RedisCfg.UserName,
		Password: GlobalCfg.RedisCfg.Password,
		DB:       GlobalCfg.RedisCfg.DB,
	})

	_, err := RDB.Ping(context.Background()).Result()
	if err != nil {
		logger.Log.Errorf("redis connection failed：%v", err)
		panic(err)
	}

	logger.Log.Info("redis connection successful")
}
