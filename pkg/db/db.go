package db

import (
	"admin-api/common/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB
// 数据库初始化
func SetUpDbLink() error {
	var err error
	var dbConfig = config.Config.Db
	url := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		dbConfig.Username,
		dbConfig.Password,
		dbConfig.Host,
		dbConfig.Port,
		dbConfig.Db,
		dbConfig.Charset)

	Db, err = gorm.Open(mysql.Open(url), &gorm.Config{
		Logger:                                   logger.Default.LogMode(logger.Info),
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		panic(err)
	}
	if Db.Error != nil {
		panic(Db.Error)
	}

	sqlDb, err := Db.DB()
	sqlDb.SetMaxIdleConns(dbConfig.MaxIdle)
	sqlDb.SetMaxOpenConns(dbConfig.MaxOpen)
	return nil

}
