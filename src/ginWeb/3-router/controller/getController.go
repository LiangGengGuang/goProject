package controller

import (
	"github.com/gin-gonic/gin"
	"project/3-router/controller/contextFunc"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

func GetApiInit(c *gin.Engine) {
	cf := contextFunc.ContextFunc{}

	//路由分组
	getG := c.Group("/get")
	{
		getG.GET("/", func(c *gin.Context) {
			cf.GetContext(c, "userName", "age")
		})

		//动态路由
		getG.GET("/:param", func(c *gin.Context) {
			cf.GetDynamicContext(c, "param")
		})
	}
}
