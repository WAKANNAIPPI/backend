package database

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"golang.org/x/crypto/bcrypt"
	"gopkg.in/ini.v1"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

//ユーザテーブル
type User struct {
	Id        string `gorm:"column:user_id" json:"userId"`
	Password  string `gorm:"column:user_password" json:"userPass"`
	Name      string `gorm:"column:user_name" json:"userName"`
	IsDeleted bool   `gorm:"column:is_deleted" json:"-"` //論理削除フラグ
}

//ユーザitem情報テーブル
type User_item struct {
	Iid      string `gorm:"column:user_item_id" json:"itemId"`        //アイテムid
	Uid      string `gorm:"column:user_id" json:"-"`                  //ユーザid
	Quantity int    `gorm:"column:user_item_quantity" json:"itemQty"` //アイテム数量
}

//Item差分管理用(ユーザアイテムのjsonのやり取りに使う)
type UserItemJson struct {
	Iid  string `json:"itemId"`
	Diff int    `json:"itemDiff"`
}

type User_constellations struct {
	Cid  string `gorm:"column:user_constellation_id" json:"consteId"`     //星座ID
	Name string `gorm:"column:user_constellation_name" json:"consteName"` //星座の名前
	Uid  string `gorm:"column:user_id" json:"-"`                          //ユーザーid
}

type Conste_star struct {
	Id    string      `gorm:"column:id" json:"-"`
	Cid   string      `gorm:"column:user_constellation_id" json:"-"`
	SStar StoredStars `gorm:"column:conste_stored_star" json:"storedStars"`
}

type Conste_line struct {
	Id     string      `gorm:"column:id" json:"-"`
	Cid    string      `gorm:"column:conste_stored_star" json:"-"`
	SLines StoredLines `gorm:"column:conste_lines" json:"storedLines"`
}

type UserConstellationJson struct {
	Cid   string      `json:"consteId"`
	Name  string      `json:"consteName"`
	Stars StoredStars `json:"storedStars"`
	Lines StoredLines `json:"storedLines"`
}

type StoredStars struct {
	StarItemId    int `json:"starItemId"`
	StarLocationX int `json:"starLocationX"`
	StarLocationY int `json:"starLocationY"`
}

type StoredLines struct {
	Sx string `json:"sx"`
	Sy string `json:"sy"`
	Fx string `json:"fx"`
	Fy string `json:"fy"`
}

type QuizeDataJson struct {
	No       int    `json:"quizeNumber"`
	Question string `json:"question"`
	Choice1  string `json:"choice1"`
	Choice2  string `json:"choice2"`
	Choice3  string `json:"choice3"`
	Choice4  string `json:"choice4"`
	Ans      int    `json:"ans"`
}

func DBconnect() *gorm.DB {
	//iniファイルをロード
	dbCfg, err := ini.Load("../database/db.ini")
	if err != nil {
		log.Panic(err)
	}

	//GormのアクセスURLを代入
	GormInfo := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbCfg.Section("db").Key("dbUserName").String(),
		os.Getenv("MYSQL_ROOT_PASSWORD"),
		os.Getenv("MYSQL_HOST"),
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
	//単一レコードを引っ張ってくる
	db.Where("user_id = ?", u.Id).First(&user)

	log.Println("databaseReqUser:", u)
	log.Println("databaseUser:", user)

	return user
}

func SetUserItemData(u User, itemDiff []UserItemJson) { //クライアントにはアイテム名と更新されたアイテム数をjsonとして渡される前提
	db := DBconnect()
	// json形式 [{"ItemId":string, Quantity:int}]
	for i := 0; i < len(itemDiff); i++ {
		UserItemBefore := User_item{}
		UserItemBefore.Iid = itemDiff[i].Iid
		UserItemBefore.Uid = u.Id

		UserItemAfter := UserItemBefore

		log.Println("1:", UserItemAfter)

		db.Where("user_item_id = ?", itemDiff[i].Iid).Where("user_id = ?", u.Id).First(&UserItemAfter)
		log.Println("2", UserItemAfter)

		UserItemAfter.Quantity += itemDiff[i].Diff
		log.Println("itemQuantity = ", UserItemAfter)

		db.Model(&UserItemBefore).Where("user_id = ?", u.Id).Where("user_item_id = ?", itemDiff[i].Iid).Update("user_item_quantity", UserItemAfter.Quantity)
	}

}

func GetUserItemData(u User) []User_item {
	db := DBconnect()

	UserItem := []User_item{}

	db.Find(&UserItem, "user_id=?", u.Id)

	return UserItem
}

func CreateUserConstellationData(u User, uc UserConstellationJson) error {
	db := DBconnect()

	UserConstellation := User_constellations{}

	UserConstellation.Uid = u.Id
	UserConstellation.Name = uc.Name
	UserConstellation.Cid = uc.Cid

	err := db.Debug().Create(&UserConstellation).Error

	return err
}

func GetUserConstellationData(u User) []User_constellations {
	db := DBconnect()

	uc := []User_constellations{}

	db.Find(&uc, "user_id=?", u.Id) //あるユーザーの作った星座情報を一括取得

	return uc
}

//ユーザー定義の構造体をGormで扱えるように定義
func (p StoredLines) Value() (driver.Value, error) {

	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err

	}
	return string(bytes), nil
}
func (p *StoredLines) Scan(input interface{}) error {
	switch v := input.(type) {
	case string:
		return json.Unmarshal([]byte(v), p)
	case []byte:
		return json.Unmarshal(v, p)
	default:
		return fmt.Errorf("unsupported type: %T", input)
	}
}
func (p StoredStars) Value() (driver.Value, error) {

	bytes, err := json.Marshal(p)
	if err != nil {
		return nil, err

	}
	return string(bytes), nil
}

func (p *StoredStars) Scan(input interface{}) error {
	switch v := input.(type) {
	case string:
		return json.Unmarshal([]byte(v), p)
	case []byte:
		return json.Unmarshal(v, p)
	default:
		return fmt.Errorf("unsupported type: %T", input)
	}
}
