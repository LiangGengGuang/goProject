package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

func GetDynamicContext(c *gin.Context) {
	c.JSON(http.StatusOK, c.Param("param"))
}

func GetContext(c *gin.Context) {
	m := make(map[string]string)
	for k := range c.Request.URL.Query() {
		m[k] = c.Query(k)
	}
	c.JSON(http.StatusOK, m)
}

func PostFormContext(c *gin.Context) {

	c.Request.ParseMultipartForm(32 << 20)
	c.JSON(http.StatusOK, c.Request.PostForm)
}

func PostJsonContext(c *gin.Context) {

	m := make(map[string]interface{})
	c.BindJSON(&m)
	c.JSON(http.StatusOK, m)

}
