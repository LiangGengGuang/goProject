package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

type User struct {
	UserName string `json:"userName"`
	Sex      string `json:"sex"`
}

func PostApiInit(c *gin.Engine) {

	//路由分组
	postG := c.Group("/post")
	{
		//form-data 表单提交
		postG.POST("", func(c *gin.Context) {

			userName := c.PostForm("userName")
			sex := c.DefaultPostForm("sex", "女")
			c.JSON(http.StatusOK, gin.H{
				"userName": userName,
				"sex":      sex,
			})
		})

		//json格式提交
		postG.POST("/json1", func(c *gin.Context) {

			json := make(map[string]interface{})
			c.BindJSON(&json)
			userName := json["userName"]
			sex := json["sex"]
			c.JSON(http.StatusOK, gin.H{
				"userName": userName,
				"sex":      sex,
			})
		})

		postG.POST("/json2", func(c *gin.Context) {

			user := User{}
			c.BindJSON(&user)
			c.JSON(http.StatusOK, user)
		})
	}
}
