package main

import (
	"github.com/gin-gonic/gin"
	"project/3-router/controller"
)

func main() {

	c := gin.Default()

	//初始化接口
	controller.ApiInit(c)

	c.Run(":8088")
}
