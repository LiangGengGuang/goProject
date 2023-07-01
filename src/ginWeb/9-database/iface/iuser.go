package iface

import (
	"project/9-database/module"
)

// @Description
// @Author lianggengguang
// @Date 2023/6/25

type IUser interface {
	CheckUsername(username string) bool

	// Register 注册
	Register(user *module.User) bool

	// Login 登录
	Login(user *module.User) bool
}
