package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"project/5-models"
	"project/9-database/logger"
	"project/9-database/module"
	"strconv"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/21

func QueryAll(c *gin.Context) {
	var goods = module.Goods{}
	c.JSON(http.StatusOK, models.SuccessResult(goods.QueryAll()))
}

func QueryById(c *gin.Context) {
	var goods = module.Goods{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"id": id}).Error("Parameter does not exist")
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Parameter:id does not exist"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResult(goods.QueryById(id)))
}

func Insert(c *gin.Context) {
	var goods = module.Goods{}
	c.BindJSON(&goods)
	userId := GetCurrentUser(c).Id
	goods.Creator = userId
	goods.Editor = userId
	if goods.Insert(&goods) {
		c.JSON(http.StatusOK, models.SuccessResult("Successfully saved"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("save failed"))
	}
}

func UpdateById(c *gin.Context) {
	var goods = module.Goods{}
	c.BindJSON(&goods)
	goods.Editor = GetCurrentUser(c).Id
	if goods.UpdateById(&goods) {
		c.JSON(http.StatusOK, models.SuccessResult("update success"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("update failed"))
	}
}

func DeleteById(c *gin.Context) {
	var goods = module.Goods{}
	id, err := strconv.Atoi(c.Query("id"))
	if err != nil {
		logger.Log.WithFields(logrus.Fields{"id": id}).Error("Parameter does not exist")
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Parameter:id does not exist"))
		return
	}
	if goods.DeleteById(id) {
		c.JSON(http.StatusOK, models.SuccessResult("Successfully deleted"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("delete failed "))
	}
}
