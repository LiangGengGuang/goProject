package module

import (
	"context"
	"fmt"
	redis2 "github.com/go-redis/redis/v8"
	"project/9-database/mysql"
	"project/9-database/redis"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/19

type IGoods struct {
	Id         int
	Name       string
	Number     int
	Price      float64
	Unit       string
	CreateTime *time.Time `gorm:"->"` //只允许读取
	UpdateTime *time.Time `gorm:"->"` //只允许读取
}

var Goods *IGoods

func (g *IGoods) TableName() string {
	return "goods"
}

func (g *IGoods) QueryAll() []*IGoods {

	var goods []*IGoods
	result := mysql.DB.Find(&goods)
	if result.Error != nil {
		fmt.Println("QueryAll fail:", result.Error)
		return nil
	}
	return goods
}

func (g *IGoods) QueryById(id int) *IGoods {

	var goods *IGoods

	//result := mysql.DB.Where("id=?", id).Find(&Goods)
	result := mysql.DB.Find(&goods, "id=?", id)
	if result.Error != nil {
		fmt.Println("QueryById fail:", result.Error)
		return nil
	}
	return goods
}

func (g *IGoods) Insert(goods *IGoods) bool {

	if goods == nil {
		return false
	}

	result := mysql.DB.Create(&goods)
	if result.Error != nil {
		fmt.Println("Insert fail:", result.Error)
		return false
	}
	return true
}

func (g *IGoods) UpdateById(goods *IGoods) bool {

	result := mysql.DB.Model(&goods).Updates(goods)
	if result.Error != nil {
		fmt.Println("UpdateById fail:", result.Error)
		return false
	}
	return true
}

func (g *IGoods) DeleteById(id int) bool {

	result := mysql.DB.Where("id =?", id).Delete(&IGoods{})
	if result.Error != nil {
		fmt.Println("DeleteById fail:", result.Error)
		return false
	}
	return true
}

//redis
var ctx = context.Background()

func (g *IGoods) Set(key, val string) {
	err := redis.RDB.Set(ctx, key, val, 60000).Err()
	if err != nil {
		panic(err)
	}
}

func (g *IGoods) Get(key string) string {
	result, err := redis.RDB.Get(ctx, key).Result()
	if err == redis2.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		fmt.Println("redis Get error：", err)
		panic(err)
	}
	return result
}
