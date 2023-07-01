package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	models "project/5-models"
	"project/9-database/module"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2023/6/25

type UserInfo struct {
	Id       int
	Username string
}

type UserJwt struct {
	UserInfo
	jwt.RegisteredClaims
}

// GenerateJWT 生成密钥
func GenerateJWT(user *module.User) (string, error) {

	userJwt := UserJwt{
		UserInfo{
			user.Id,
			user.Username,
		},
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * time.Minute)), //过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),                       //签发时间
			NotBefore: jwt.NewNumericDate(time.Now()),                       //生效时间
		},
	}
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, userJwt)
	signingString, err := claims.SignedString([]byte("a1b2c3"))
	return signingString, err
}

// ParseJWT 解析密钥
func ParseJWT(token string) (*UserJwt, error) {
	claims, err := jwt.ParseWithClaims(token, &UserJwt{}, func(token *jwt.Token) (interface{}, error) {
		return []byte("a1b2c3"), nil
	})

	if userJwt, ok := claims.Claims.(*UserJwt); ok && claims.Valid {
		return userJwt, nil
	} else {
		return nil, err
	}
}

// Register 注册
func Register(c *gin.Context) {
	var user = module.User{}
	c.BindJSON(&user)

	exist := user.CheckUsername(user.Username)
	if exist {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("username already exists"))
		return
	}
	if user.Register(&user) {
		c.JSON(http.StatusOK, models.SuccessResult("Register successfully"))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Register failed"))
	}
}

// Login 登录
func Login(c *gin.Context) {
	var user = module.User{}
	c.BindJSON(&user)
	if user.Login(&user) {
		token, err := GenerateJWT(&user)
		if err != nil {
			c.JSON(http.StatusInternalServerError, models.ErrorResult("generate token failed"))
			return
		}
		c.JSON(http.StatusOK, models.SuccessResult(token))
	} else {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Login failed"))
	}
}

// Logout 退出
func Logout(c *gin.Context) {
	RemoveCurrentUser(c)
	c.JSON(http.StatusOK, models.SuccessResult("Logout successfully"))
}

// RefreshToken 刷新token
func RefreshToken(c *gin.Context) {
	userInfo := GetCurrentUser(c)
	user := &module.User{
		Id:       userInfo.Id,
		Username: userInfo.Username,
	}
	token, err := GenerateJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, models.ErrorResult("Refresh token token failed"))
		return
	}
	c.JSON(http.StatusOK, models.SuccessResult(token))
}
