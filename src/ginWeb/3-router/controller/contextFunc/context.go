package contextFunc

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Description
// @Author lianggengguang
// @Date 2022/5/22

type ContextFunc struct {
}

func (cf *ContextFunc) GetDynamicContext(c *gin.Context, object ...string) {
	if !checkParams(c, nil) {
		return
	}

	m := make(map[string]string)
	for i := range object {
		m[object[i]] = c.Param(object[i])
	}
	c.JSON(http.StatusOK, m)
}

func (cf *ContextFunc) GetContext(c *gin.Context, object ...string) {
	if !checkParams(c, object) {
		return
	}

	m := make(map[string]string)
	for i := range object {
		m[object[i]] = c.Query(object[i])
	}
	c.JSON(http.StatusOK, m)
}

func (cf *ContextFunc) PostFormContext(c *gin.Context, object ...string) {
	if !checkParams(c, object) {
		return
	}

	m := make(map[string]interface{})
	for i := range object {
		m[object[i]] = c.PostForm(object[i])
	}
	c.JSON(http.StatusOK, m)
}

func (cf *ContextFunc) PostJsonContext(c *gin.Context, object interface{}) {
	if c == nil {
		return
	}

	if object == nil {
		object = make(map[string]interface{})
	}
	c.BindJSON(&object)
	c.JSON(http.StatusOK, object)
}

func checkParams(c *gin.Context, object ...interface{}) bool {
	if c == nil || object == nil || len(object) == 0 {
		return false
	}
	return true
}
