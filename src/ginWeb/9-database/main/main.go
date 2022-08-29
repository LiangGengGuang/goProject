package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"project/9-database/controller"
	"project/9-database/db"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

func main() {

	c := gin.Default()

	controller.ApiInit(c)

	portStr := fmt.Sprintf(":%d", db.GlobalCfg.Port)
	c.Run(portStr)
}
