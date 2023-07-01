package module

import (
	"project/9-database/db"
	"project/9-database/logger"
	"project/9-database/utils"
)

// @Description
// @Author lianggengguang
// @Date 2023/6/25

type User struct {
	Id         int           `json:"id"`
	Username   string        `json:"username"`
	Pwd        string        `json:"pwd"`
	CreateTime *utils.MyTime `gorm:"->" json:"create_time"` //只允许读取
	UpdateTime *utils.MyTime `gorm:"->" json:"update_time"` //只允许读取
}

func (u *User) TableName() string {
	return "user"
}

func (u *User) CheckUsername(username string) bool {
	var user *User
	affected := db.MDB.Where("username =?", username).Find(&user).RowsAffected
	if affected > 0 {
		return true
	}
	return false
}

// Register 注册
func (u *User) Register(user *User) bool {
	if user.Username == "" || user.Pwd == "" {
		logger.Log.Error("register failed:username or pwd does not nil")
		return false
	}
	result := db.MDB.Create(&user)
	if result.Error != nil {
		logger.Log.Errorf("register failed: %v", result.Error)
		return false
	}
	return true
}

// Login 登录
func (u *User) Login(user *User) bool {
	if user.Username == "" || user.Pwd == "" {
		logger.Log.Error("login failed:username or pwd does not nil")
		return false
	}
	affect := db.MDB.Where("username=? and pwd=? ", user.Username, user.Pwd).Find(&user).RowsAffected
	if affect == 0 {
		logger.Log.Error("login failed: incorrect username or password")
		return false
	}
	return true
}
