package controller

import (
	"github.com/gin-gonic/gin"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/21

func ApiInit(e *gin.Engine) {
	group := e.Group("/goods")
	{
		group.GET("/queryAll", QueryAll)

		group.GET("/queryById", QueryById)

		group.POST("/insert", Insert)

		group.PUT("/updateById", UpdateById)

		group.DELETE("/deleteById", DeleteById)
	}
}
