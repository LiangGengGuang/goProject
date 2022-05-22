package main

import (
	"github.com/gin-gonic/gin"
	"project/3-router/controller"
)

func main() {

	c := gin.Default()

	//GET请求
	controller.GetApiInit(c)

	//POST请求
	controller.PostApiInit(c)

	c.Run(":8088")

}
