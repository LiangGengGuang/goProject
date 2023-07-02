package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	models "project/5-models"
	"project/9-database/db"
	"project/9-database/logger"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/21

//请求头token校验
func apiHeadMiddleware(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		logger.Log.Error("Authorization does not exist")
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Authorization does not exist"))
		c.Abort()
		return
	}
	//token黑名单校验
	result, err := db.RDB.Exists(db.RDB.Context(), authToken).Result()
	if err != nil {
		logger.Log.Errorf("blacklist add failed:%v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Authorization does not exist"))
		c.Abort()
		return
	}
	if result == 1 {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Authorization already expired"))
		c.Abort()
		return
	}
	//token解析
	userJWT, err := ParseJWT(authToken)
	if err != nil {
		logger.Log.Errorf("token parse failed：%v", err)
		c.JSON(http.StatusInternalServerError, models.ErrorResult("token parse failed"))
		c.Abort()
		return
	}
	setCurrentUser(c, &userJWT.UserInfo)
}

// GetCurrentUser 获取当前用户信息
func GetCurrentUser(c *gin.Context) *UserInfo {
	session := sessions.Default(c)
	userInfo := session.Get("currentUser").(*UserInfo) // 类型转换一下
	return userInfo
}

// 保存当前用户信息
func setCurrentUser(c *gin.Context, userInfo *UserInfo) {
	session := sessions.Default(c)
	session.Set("currentUser", userInfo)
	// 一定要Save否则不生效，若未使用gob注册User结构体，调用Save时会返回一个Error
	session.Save()
}

func RemoveCurrentUser(c *gin.Context) {
	session := sessions.Default(c)
	session.Delete("currentUser")
}

func ApiInit(e *gin.Engine) {
	goodsGroup := e.Group("/goods")
	goodsGroup.Use(apiHeadMiddleware)
	{
		goodsGroup.GET("/queryAll", QueryAll)

		goodsGroup.GET("/queryById", QueryById)

		goodsGroup.POST("/insert", Insert)

		goodsGroup.PUT("/updateById", UpdateById)

		goodsGroup.DELETE("/deleteById", DeleteById)
	}

	userGroup := e.Group("/auth")
	{
		userGroup.POST("/register", Register)

		userGroup.POST("/login", Login)

		userGroup.POST("/logout", apiHeadMiddleware, Logout)

		userGroup.POST("/refreshToken", apiHeadMiddleware, RefreshToken)
	}
}
