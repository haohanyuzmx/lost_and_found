package model

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var (
	DB *gorm.DB
)

const (
	dns = "root:@tcp(127.0.0.1:3306)/lost_and_found?charset=utf8mb4&parseTime=True&loc=Local"
)

func init() {
	db, err := gorm.Open(mysql.Open(dns), &gorm.Config{DisableForeignKeyConstraintWhenMigrating: true /*,Logger: logger.Default.LogMode(logger.Info)*/})
	if err != nil {
		log.Panicln(err)
		return
	}
	//
	if db.AutoMigrate(&File{}, &Comment{}, &LostAndFound{}, &NoticeForLost{}, &User{}) != nil {
		log.Println(err)
		return
	}

	DB = db
	DB.Debug()
}
