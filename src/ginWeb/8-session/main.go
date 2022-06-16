package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"net/http"
	models "project/5-models"
)

var key = "name"

func NewStore(e *gin.Engine) {

	store, _ := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	e.Use(sessions.Sessions("redis", store))
}

func SetSession(c *gin.Context) {

	value := "lgg"
	session := sessions.Default(c)

	//设置session过期时间(redis过期时间)
	session.Options(sessions.Options{
		MaxAge: 6000,
	})
	session.Set(key, value)
	session.Save()
	c.JSON(http.StatusOK, models.SuccessResult(value))
}

func GetSession(c *gin.Context) {

	session := sessions.Default(c)
	if value := session.Get(key); value != nil {
		c.JSON(http.StatusOK, models.SuccessResult(value))
	} else {
		c.JSON(http.StatusOK, models.SuccessResult("session获取失败"))
	}
}

func main() {

	e := gin.Default()

	//配置session中间件，redis存储引擎
	NewStore(e)

	group := e.Group("/session")
	{
		group.GET("/setSession", func(c *gin.Context) {
			SetSession(c)
		})

		group.GET("/getSession", func(c *gin.Context) {
			GetSession(c)
		})
	}

	e.Run(":8088")
}
