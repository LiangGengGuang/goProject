package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	models "project/5-models"
)

func SetCookie(c *gin.Context) {
	value := "lgg"
	c.SetCookie("name", value, 3600, "/", "localhost", false, true)
	c.JSON(http.StatusOK, models.SuccessResult(value))
}

func GetCookie(c *gin.Context) {
	cookie, err := c.Cookie("name")
	if err != nil {
		fmt.Println("cookie获取失败:", err)
		c.JSON(http.StatusOK, models.ErrorResult("cookie获取失败"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResult(cookie))
}

func ExpireCookie(c *gin.Context) {

	if _, err := c.Cookie("name"); err != nil {
		fmt.Println("cookie不存在：", err)
		c.JSON(http.StatusOK, models.SuccessResult("设置cookie过期成功"))
		return
	}
	c.SetCookie("name", "", -1, "/", "localhost", false, true)
	c.JSON(http.StatusOK, models.SuccessResult("设置cookie过期成功"))
}

func main() {

	e := gin.Default()
	group := e.Group("/cookie")
	{
		group.GET("/setCookie", func(c *gin.Context) {
			SetCookie(c)
		})

		group.GET("/getCookie", func(c *gin.Context) {
			GetCookie(c)
		})

		group.GET("/expireCookie", func(c *gin.Context) {
			ExpireCookie(c)
		})
	}

	e.Run(":8088")

}
