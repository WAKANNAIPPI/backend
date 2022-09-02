package database

import (
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
//ユーザテーブル
type User struct {
	Id        string `gorm:"column:user_id"`
	Name      string `gorm:"column:user_name"`
	Password  string `gorm:"column:user_password"`
	IsDeleted bool   `gorm:"column:is_deleted"` //論理削除フラグ
}
//ユーザitem情報テーブル
type User_item struct {
	Iid      string `gorm:"column:user_item_id"`       //アイテムid
	Uid      string `gorm:"column:user_id"`            //ユーザid
	Quantity int    `gorm:"column:user_item_quantity"` //アイテム数量
}


//Item差分管理用(ユーザアイテムのjsonのやり取りに使う)
type Item_difference struct {
	Iid string`json:"itemid"`
	Quantity string `json:"diff"`
	
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

	//データベースに登録
	db.Create(&u)

}

func GetUserData(u User) User { //ユーザ情報を取得する関数
	db := DBconnect()
	//UserIDを入れてレコードを特定
	user := User{}
	user.Id = u.Id

	//単一レコードを引っ張ってくる
	db.First(&user)

	return user
}

func SetUserItemData(u User, item []) { //クライアントにはアイテム名と更新されたアイテム数をjsonとして渡される前提
	db := DBconnect()

	// json形式 [{"ItemId":"", Quantity:""}]
	for i := 0; i < len(item); i++{
		
	}

}

func GetUserItemData(u User) {

}