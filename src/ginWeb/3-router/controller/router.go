package controller

import (
	"github.com/gin-gonic/gin"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

func ApiInit(c *gin.Engine) {

	//路由分组
	getG := c.Group("/get")
	{
		getG.GET("/", GetContext)

		//动态路由
		getG.GET("/:param", GetDynamicContext)
	}

	//路由分组
	postG := c.Group("/post")
	{
		//form-data 表单提交
		postG.POST("/", PostFormContext)

		//json格式提交
		//json格式提交
		postG.POST("/json", PostJsonContext)

	}
}
