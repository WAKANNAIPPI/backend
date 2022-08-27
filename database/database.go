package database

import (
	"fmt"
	"log"

	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBconnect() *gorm.DB {
	dbCfg, err := ini.Load("db.ini")
	if err != nil {
		log.Panic(err)
	}

	GormInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.Section("db").Key("dbUserName").String(),
		dbCfg.Section("db").Key("dbUserPass").String(),
		dbCfg.Section("db").Key("dbHost").String(),
		dbCfg.Section("db").Key("dbPort").String(),
		dbCfg.Section("db").Key("dbName").String(),
	)

	db, err := gorm.Open(mysql.Open(GormInfo), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}

	return db
}
