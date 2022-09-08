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
	Password  string `gorm:"column:user_password"`
	Name      string `gorm:"column:user_name"`
	IsDeleted bool   `gorm:"column:is_deleted"` //論理削除フラグ
}

//ユーザitem情報テーブル
type User_item struct {
	Iid      string `gorm:"column:user_item_id"`       //アイテムid
	Uid      string `gorm:"column:user_id"`            //ユーザid
	Quantity int    `gorm:"column:user_item_quantity"` //アイテム数量
}

type User_constellations struct {
	Cid  string `gorm:"column:user_constellation_id"`   //星座ID
	Name string `gorm:"column:user_constellation_name"` //星座の名前
	Uid  string `gorm:"column:user_id"`                 //ユーザーid
	Data int    `gorm:"column:user_constellation_data"` //星座データ
}

//Item差分管理用(ユーザアイテムのjsonのやり取りに使う)
type Item_difference struct {
	Iid  string `json:"itemid"`
	Diff int    `json:"diff"`
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

func SetUserItemData(u User, itemDiff []Item_difference) { //クライアントにはアイテム名と更新されたアイテム数をjsonとして渡される前提
	db := DBconnect()
	// json形式 [{"ItemId":string, Quantity:int}]
	for i := 0; i < len(itemDiff); i++ {
		UserItemBefore := User_item{}
		UserItemBefore.Iid = itemDiff[i].Iid
		UserItemBefore.Uid = u.Id

		UserItemAfter := UserItemBefore

		db.First(&UserItemAfter)

		UserItemAfter.Quantity += itemDiff[i].Diff

		db.Model(&UserItemBefore).Where("user_id = ?", u.Id).Where("user_item_id = ?", itemDiff[i].Iid).Update("user_item_quantity", UserItemAfter.Quantity)
	}

}

func GetUserItemData(u User) []User_item {
	db := DBconnect()

	UserItem := []User_item{}

	db.Find(&UserItem, "user_id=?", u.Id)

	return UserItem
}
