package config

import (
	"encoding/json"
	"io/ioutil"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/17

type GlobalConfig struct {
	DbCfg    *MysqlConfig
	RedisCfg *RedisConfig
}

type MysqlConfig struct {
	Uri         string
	UserName    string
	Password    string
	MaxConnTime time.Duration
	MaxOpenConn int
	MaxIdleConn int
}
type RedisConfig struct {
	Uri      string
	Port     int
	UserName string
	Password string
}

var GlobalCfg *GlobalConfig

func (cfg *GlobalConfig) Reload() {

	//文件的路径：go.mod的相对路径
	file, err := ioutil.ReadFile("9-database/main/staticFile/config.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(file, &GlobalCfg)
	if err != nil {
		panic(err)
	}
}

func init() {

	dbCfg := &MysqlConfig{
		Uri:         "",
		UserName:    "",
		Password:    "",
		MaxConnTime: 1,
		MaxOpenConn: 1,
		MaxIdleConn: 1,
	}

	redisCfg := &RedisConfig{
		Uri:      "",
		Port:     6379,
		UserName: "",
		Password: "",
	}
	GlobalCfg = &GlobalConfig{
		DbCfg:    dbCfg,
		RedisCfg: redisCfg,
	}
	GlobalCfg.Reload()
}
