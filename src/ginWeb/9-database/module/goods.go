package module

import (
	"encoding/json"
	"project/9-database/db"
	"project/9-database/logger"
	"project/9-database/utils"
	"strconv"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/19

type Goods struct {
	Id         int           `json:"id"`
	Name       string        `json:"name"`
	Number     int           `json:"number"`
	Price      float64       `json:"price"`
	Unit       string        `json:"unit"`
	Creator    int           `json:"creator"`
	CreateTime *utils.MyTime `gorm:"->" json:"create_time"` //只允许读取
	Editor     int           `json:"editor"`
	UpdateTime *utils.MyTime `gorm:"->" json:"update_time"` //只允许读取
}

func (g *Goods) TableName() string {
	return "goods"
}

func (g *Goods) QueryAll() []*Goods {

	var goods []*Goods
	result := db.MDB.Order("id desc").Find(&goods)
	if result.Error != nil {
		logger.Log.Errorf("QueryAll fail: %v", result.Error)
		return nil
	}
	return goods
}

func (g *Goods) QueryById(id int) *Goods {

	var goods *Goods

	val := g.Get(strconv.Itoa(id))
	if val != "" {
		json.Unmarshal([]byte(val), &goods)
		return goods
	}
	result := db.MDB.Where("id=?", id).Find(&goods)
	//result := db.MDB.Find(&goods, "id=?", id)
	//db.MDB.Select("name","number").Find(&goods, "id=?", id) 只显示特定字段
	if result.Error != nil {
		logger.Log.Errorf("QueryById id: %d fail: %v", id, result.Error)
		return nil
	}
	if goods.Id != 0 {
		if marshal, err := json.Marshal(goods); err == nil {
			g.Set(strconv.Itoa(id), marshal)
		}
		return goods
	} else {
		return nil
	}
}

func (g *Goods) Insert(goods *Goods) bool {

	if goods == nil {
		return false
	}

	result := db.MDB.Create(&goods)
	if result.Error != nil {
		logger.Log.Errorf("Insert fail: %v", result.Error)
		return false
	}
	return true
}

func (g *Goods) UpdateById(goods *Goods) bool {

	result := db.MDB.Model(&goods).Updates(goods)
	if result.Error != nil {
		logger.Log.Errorf("UpdateById id: %d fail:%v", goods.Id, result.Error)
		return false
	}
	goods = g.QueryById(goods.Id)
	return true
}

func (g *Goods) DeleteById(id int) bool {

	g.Del(strconv.Itoa(id))

	result := db.MDB.Where("id =?", id).Delete(&Goods{})
	if result.Error != nil {
		logger.Log.Errorf("DeleteById id: %d fail:%v", id, result.Error)
		return false
	}
	return true
}

func (g *Goods) Set(key string, val interface{}) {
	result, err := db.RDB.Set(db.RDB.Context(), key, val, 0).Result()
	if err != nil {
		logger.Log.Errorf("redis Set error：%v", err)
	} else {
		logger.Log.Infof("redis key: %s set success：%v", key, result)
	}
}

func (g *Goods) Get(key string) string {
	val, err := db.RDB.Get(db.RDB.Context(), key).Result()
	if err != nil {
		return ""
	}
	return val
}

func (g *Goods) Del(key string) {
	result, err := db.RDB.Del(db.RDB.Context(), key).Result()
	if err != nil {
		logger.Log.Errorf("redis key: %s del error：%v", key, err)
	} else {
		logger.Log.Infof("redis key: %s del success：%v", key, result)
	}
}
