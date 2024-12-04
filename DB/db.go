package DB

import (
	"WeiYangWork/Model"
	"WeiYangWork/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"time"
)

func initDB() error {
	dsn := AppConfig.DataBase.Dsn
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatalf("无法链接到数据库 %v", err.Error())
		return err
	}
	sqldb, _ := db.DB()
	sqldb.SetMaxIdleConns(AppConfig.DataBase.MaxIdleConns)
	sqldb.SetConnMaxLifetime(24 * time.Hour)

	if err = db.AutoMigrate(&Model.User{}, &Model.Team{}, &Model.Message{}); err != nil {
		return err
	}
	global.Db = db
	return nil
}
