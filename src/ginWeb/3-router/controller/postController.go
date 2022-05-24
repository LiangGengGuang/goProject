package controller

import (
	"github.com/gin-gonic/gin"
	"project/3-router/controller/contextFunc"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

type User struct {
	UserName string `json:"userName"`
	Sex      string `json:"sex"`
}

func PostApiInit(c *gin.Engine) {
	cf := contextFunc.ContextFunc{}

	//路由分组
	postG := c.Group("/post")
	{
		//form-data 表单提交
		postG.POST("", func(c *gin.Context) {
			cf.PostFormContext(c, "userName", "age")
		})

		//json格式提交
		postG.POST("/json1", func(c *gin.Context) {
			cf.PostJsonContext(c, nil)
		})

		postG.POST("/json2", func(c *gin.Context) {
			cf.PostJsonContext(c, User{})
		})
	}
}
