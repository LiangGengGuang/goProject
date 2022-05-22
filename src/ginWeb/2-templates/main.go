package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Body struct {
	Context string
}

func main() {

	engine := gin.Default()
	engine.LoadHTMLGlob("templs/**/*")

	engine.GET("/posts/index", func(c *gin.Context) {

		body := &Body{
			Context: "/posts/index",
		}
		c.HTML(http.StatusOK, "posts/index.html", gin.H{
			"body": body,
		})
	})

	engine.GET("/users/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "users/index.html", gin.H{
			"context": "/users/index",
		})
	})

	engine.Run(":8088")

}
