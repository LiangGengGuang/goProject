package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

func GetApiInit(c *gin.Engine) {

	//路由分组
	getG := c.Group("/get")
	{
		getG.GET("/", func(c *gin.Context) {

			userName := c.Query("userName")
			page := c.DefaultQuery("page", "1")
			c.JSON(http.StatusOK, gin.H{
				"userName": userName,
				"page":     page,
			})
		})

		//动态路由
		getG.GET("/:param", func(c *gin.Context) {

			param := c.Param("param")
			c.JSON(http.StatusOK, gin.H{
				"param": param,
			})
		})
	}
}
