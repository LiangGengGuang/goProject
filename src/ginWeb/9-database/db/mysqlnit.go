package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/9-database/config"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

// MDB Db数据库连接池
var MDB *gorm.DB

//创建mysql连接
func init() {

	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8mb4&parseTime=True&loc=Local", config.GlobalCfg.DbCfg.UserName, config.GlobalCfg.DbCfg.Password, config.GlobalCfg.DbCfg.Uri)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		panic(err)
	}
	sqlDB, err := mdb.DB()
	if err != nil {
		fmt.Println("数据库连接失败：", err)
		panic(err)
	}
	//设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(config.GlobalCfg.DbCfg.MaxConnTime)

	//置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.GlobalCfg.DbCfg.MaxOpenConn)

	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.GlobalCfg.DbCfg.MaxIdleConn)

	MDB = mdb
	fmt.Println("数据库连接成功")
}
