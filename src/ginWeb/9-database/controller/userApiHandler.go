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

var goods = module.NewGoods()

func QueryAll(c *gin.Context) {
	c.JSON(http.StatusOK, models.SuccessResult(goods.QueryAll()))
}

func QueryById(c *gin.Context) {

	query := c.Query("id")
	id, err := strconv.Atoi(query)
	if err != nil {
		fmt.Printf("参数不存在：%d\n", id)
		c.JSON(http.StatusInternalServerError, models.ErrorResult("id参数不存在"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResult(goods.QueryById(id)))
}

func Insert(c *gin.Context) {
	c.BindJSON(&goods)
	if goods.Insert(&goods) {
		c.JSON(http.StatusOK, models.SuccessResult("保存成功"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("保存失败"))
	}
}

func UpdateById(c *gin.Context) {

	c.BindJSON(&goods)
	if goods.UpdateById(&goods) {
		c.JSON(http.StatusOK, models.SuccessResult("更新成功"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("更新失败"))
	}
}

func DeleteById(c *gin.Context) {

	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		fmt.Printf("参数不存在：%d\n", id)
		c.JSON(http.StatusInternalServerError, models.ErrorResult("id参数不存在"))
		return
	}
	if goods.DeleteById(id) {
		c.JSON(http.StatusOK, models.SuccessResult("删除成功"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("删除失败"))
	}
}
