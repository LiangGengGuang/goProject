package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/9-database/controller"
	"project/9-database/db"
	"project/9-database/logger"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

func main() {

	//启动gin
	logger.Log.Info("gin server starting...")

	gin.SetMode(gin.ReleaseMode)
	c := gin.Default()
	controller.ApiInit(c)
	portStr := fmt.Sprintf(":%d", db.GlobalCfg.Port)
	c.Run(portStr)
}
