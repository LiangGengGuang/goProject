package module

import (
	"encoding/json"
	"fmt"
	"project/9-database/db"
	"strconv"
	"time"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/19

type IGoods struct {
	Id         int        `json:"id"`
	Name       string     `json:"name"`
	Number     int        `json:"number"`
	Price      float64    `json:"price"`
	Unit       string     `json:"unit"`
	CreateTime *time.Time `gorm:"->" json:"create_time"` //只允许读取
	UpdateTime *time.Time `gorm:"->" json:"update_time"` //只允许读取
}

var Goods *IGoods

func (g *IGoods) TableName() string {
	return "goods"
}

func (g *IGoods) QueryAll() []*IGoods {

	var goods []*IGoods
	result := db.MDB.Find(&goods)
	if result.Error != nil {
		fmt.Println("QueryAll fail:", result.Error)
		return nil
	}
	return goods
}

func (g *IGoods) QueryById(id int) *IGoods {

	var goods *IGoods

	val := g.Get(strconv.Itoa(id))
	if val != "" {
		json.Unmarshal([]byte(val), &goods)
	} else {
		//result := db.MDB.Where("id=?", id).Find(&Goods)
		result := db.MDB.Find(&goods, "id=?", id)
		if result.Error != nil {
			fmt.Println("QueryById fail:", result.Error)
			return nil
		}
		if marshal, err := json.Marshal(goods); err == nil {
			g.Set(strconv.Itoa(id), marshal)
		}
	}
	return goods
}

func (g *IGoods) Insert(goods *IGoods) bool {

	if goods == nil {
		return false
	}

	result := db.MDB.Create(&goods)
	if result.Error != nil {
		fmt.Println("Insert fail:", result.Error)
		return false
	}
	goods = g.QueryById(goods.Id)
	if marshal, err := json.Marshal(goods); err == nil {
		g.Set(strconv.Itoa(goods.Id), marshal)
	}
	return true
}

func (g *IGoods) UpdateById(goods *IGoods) bool {

	result := db.MDB.Model(&goods).Updates(goods)
	if result.Error != nil {
		fmt.Println("UpdateById fail:", result.Error)
		return false
	}
	goods = g.QueryById(goods.Id)
	if marshal, err := json.Marshal(goods); err == nil {
		g.Set(strconv.Itoa(goods.Id), marshal)
	}
	return true
}

func (g *IGoods) DeleteById(id int) bool {

	g.Del(strconv.Itoa(id))

	result := db.MDB.Where("id =?", id).Delete(&IGoods{})
	if result.Error != nil {
		fmt.Println("DeleteById fail:", result.Error)
		return false
	}
	return true
}

func (g *IGoods) Set(key string, val interface{}) {
	result, err := db.RDB.Set(db.RDB.Context(), key, val, 0).Result()
	if err != nil {
		fmt.Println("redis Set error：", err)
	} else {
		fmt.Println("redis Set success：", result)
	}
}

func (g *IGoods) Get(key string) string {
	val, err := db.RDB.Get(db.RDB.Context(), key).Result()
	if err != nil {
		fmt.Println("redis Get error：", err)
		return ""
	}
	return val
}

func (g *IGoods) Del(key string) {
	result, err := db.RDB.Del(db.RDB.Context(), key).Result()
	if err != nil {
		fmt.Println("redis Del error：", err)
	} else {
		fmt.Println("redis Del success：", result)
	}
}
