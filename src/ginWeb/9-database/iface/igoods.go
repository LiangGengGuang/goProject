package iface

import "project/9-database/module"

// @Description
// @Author lianggengguang
// @Date 2022/8/23

type IGoods interface {

	/*
		mysql
	*/
	QueryAll() []*module.Goods

	QueryById(id int) *module.Goods

	Insert(goods *module.Goods) bool

	UpdateById(goods *module.Goods) bool

	DeleteById(id int) bool

	/*
		redis
	*/
	Set(key string, val interface{})

	Get(key string) string

	Del(key string)
}
