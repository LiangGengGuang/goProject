package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"project/5-models"
	"project/9-database/module"
	"strconv"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/21

func MysqlController(e *gin.Engine) {
	group := e.Group("/mysql")
	{

		group.GET("/queryAll", func(c *gin.Context) {

			c.JSON(http.StatusOK, models.SuccessResult(module.Goods.QueryAll()))
		})

		group.GET("/queryById", func(c *gin.Context) {

			query := c.Query("id")
			id, err := strconv.Atoi(query)
			if err != nil {
				fmt.Printf("参数不存在：%d\n", id)
				c.JSON(http.StatusInternalServerError, models.ErrorResult("id参数不存在"))
				return
			}
			c.JSON(http.StatusOK, models.SuccessResult(module.Goods.QueryById(id)))
		})

		group.POST("/insert", func(c *gin.Context) {
			goods := module.IGoods{}
			c.BindJSON(&goods)
			if module.Goods.Insert(&goods) {
				c.JSON(http.StatusOK, models.SuccessResult("保存成功"))
			} else {
				c.JSON(http.StatusInternalServerError, models.ErrorResult("保存失败"))
			}
		})

		group.PUT("/updateById", func(c *gin.Context) {
			goods := module.IGoods{}
			c.BindJSON(&goods)
			if module.Goods.UpdateById(&goods) {
				c.JSON(http.StatusOK, models.SuccessResult("更新成功"))
			} else {
				c.JSON(http.StatusInternalServerError, models.ErrorResult("更新失败"))
			}
		})

		group.DELETE("/deleteById", func(c *gin.Context) {
			id, err := strconv.Atoi(c.Query("id"))
			if err != nil {
				fmt.Printf("参数不存在：%d\n", id)
				c.JSON(http.StatusInternalServerError, models.ErrorResult("id参数不存在"))
				return
			}
			if module.Goods.DeleteById(id) {
				c.JSON(http.StatusOK, models.SuccessResult("删除成功"))
			} else {
				c.JSON(http.StatusInternalServerError, models.ErrorResult("删除失败"))
			}
		})
	}
}
