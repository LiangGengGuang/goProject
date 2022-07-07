package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/9-database/config"
	"project/9-database/controller"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

func main() {

	e := gin.Default()

	//mysql
	controller.MysqlController(e)

	//redis
	//controller.RedisController(e)

	portStr := fmt.Sprintf(":%d", config.GlobalCfg.Port)
	e.Run(portStr)
}
