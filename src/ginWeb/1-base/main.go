package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Message struct {
	Msg     string `json:"msg"`
	Context string `json:"context"`
}

func main() {

	//创建一个默认的路由引擎
	r := gin.Default()

	// 当客户端以GET方法请求/ping，会执行后面的匿名函数
	r.GET("/ping1", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	r.GET("/ping2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/ping3", func(c *gin.Context) {
		m := &Message{
			Msg:     "success",
			Context: "pong",
		}
		c.JSON(http.StatusOK, m)
	})

	r.GET("/xml", func(c *gin.Context) {
		m := &Message{
			Msg:     "success",
			Context: "pong",
		}

		c.XML(http.StatusOK, m)
	})

	// 启动HTTP服务，默认在0.0.0.0:8080启动服务
	r.Run(":8088")
}
