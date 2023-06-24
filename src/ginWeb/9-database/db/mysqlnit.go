package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"project/9-database/logger"
)

// @Description
// @Author lianggengguang
// @Date 2022/6/16

// MDB Db数据库连接池
var MDB *gorm.DB

//创建mysql连接
func init() {

	dsn := fmt.Sprintf("%s:%s@%s?charset=utf8mb4&parseTime=True&loc=Local", GlobalCfg.DbCfg.UserName, GlobalCfg.DbCfg.Password, GlobalCfg.DbCfg.Uri)
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 100,  //设置批量插入最大数量
		QueryFields:     true, //根据当前 model 的所有字段名称进行 select
	})
	if err != nil {
		logger.Log.Errorf("database connection failed：%v", err)
		panic(err)
	}
	sqlDB, err := mdb.DB()
	if err != nil {
		logger.Log.Errorf("database connection failed：%v", err)
		panic(err)
	}
	//设置了连接可复用的最大时间
	sqlDB.SetConnMaxLifetime(GlobalCfg.DbCfg.MaxConnTime)

	//置打开数据库连接的最大数量
	sqlDB.SetMaxOpenConns(GlobalCfg.DbCfg.MaxOpenConn)

	//设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(GlobalCfg.DbCfg.MaxIdleConn)

	MDB = mdb
	logger.Log.Info("database connection successful")
}
