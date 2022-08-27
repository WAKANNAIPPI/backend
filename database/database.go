package database

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	Id        string `gorm:"column:user_id"`
	Name      string `gorm:"column:user_name"`
	Password  string `gorm:"column:user_password"`
	IsDeleted bool   `gorm:"column:is_deleted"`
}

func DBconnect() *gorm.DB {
	//iniファイルをロード
	dbCfg, err := ini.Load("db.ini")
	if err != nil {
		log.Panic(err)
	}

	//GormのアクセスURLを代入
	GormInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.Section("db").Key("dbUserName").String(),
		dbCfg.Section("db").Key("dbUserPass").String(),
		dbCfg.Section("db").Key("dbHost").String(),
		dbCfg.Section("db").Key("dbPort").String(),
		dbCfg.Section("db").Key("dbName").String(),
	)

	//GORMによるDBアクセス
	db, err := gorm.Open(mysql.Open(GormInfo), &gorm.Config{})

	if err != nil {
		log.Panic(err)
	}
	//アクセスされたオブジェクトを返す
	return db
}

func CreateUser(u User) { //ユーザー作成関数
	db := DBconnect()

	//passwordのhash化
	uPas := []byte(u.Password)
	uHashedPas, _ := bcrypt.GenerateFromPassword(uPas, 5)
	u.Password = string(uHashedPas)

	db.Create(&u)

}
