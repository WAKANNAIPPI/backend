package model

import (
	"backend/database"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
	"golang.org/x/crypto/bcrypt"
)

func Userlogin(ctx *gin.Context) {
	reqUser := database.User{}
	//クライアントからのjsonデータをユーザー構造体にbinding
	err := ctx.BindJSON(&reqUser)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		log.Println(err)
	} else {

		pass := reqUser.Password
		user := database.GetUserData(reqUser)
		//DBから取得してきたpasswordはハッシュ値
		hashPass := user.Password

		//password比較
		err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))

		if err != nil {
			ctx.Status(http.StatusBadRequest)
			log.Println(err)
		} else {
			//sessionのセットアップ
			session := sessions.Default(ctx)

			//セッションにuserIDを格納
			log.Println("userdata:", user)
			log.Println("reqdata:", reqUser)
			sessionUser, err := json.Marshal(user)

			if err == nil {
				session.Set("loginUser", string(sessionUser))
				session.Save()

				log.Println("session Log", session.Get("loginUser"))
			} else {
				ctx.Status(http.StatusInternalServerError)
			}
		}
	}
}

func GetUserItem(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//構造体インスタンスの生成
	user := database.User{}

	//sessionから取ったユーザ情報の構造体へのマッピング
	userJson, err := dproxy.New(session.Get("loginUser")).String()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}

	//dbからuserItem情報を取得してjson形式で返却
	userItems := database.GetUserItemData(user)
	log.Println(user, ":", userItems)

	ctx.JSON(200, userItems)
}

func PostUserItem(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//構造体インスタンスの生成
	user := database.User{}

	//sessionから取ったユーザ情報の構造体へのマッピング
	userJson, err := dproxy.New(session.Get("loginUser")).String()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}

	var ItemDiff []database.UserItemJson

	err = ctx.BindJSON(&ItemDiff)

	log.Println(ItemDiff)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
		ctx.Abort()
	}

	database.SetUserItemData(user, ItemDiff)
}

func GetConsteData(ctx *gin.Context) {
	session := sessions.Default(ctx)

	//構造体インスタンスの生成
	user := database.User{}

	//sessionから取ったユーザ情報の構造体へのマッピング
	userJson, err := dproxy.New(session.Get("loginUser")).String()
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}
	err = json.Unmarshal([]byte(userJson), &user)
	if err != nil {
		ctx.Status(http.StatusInternalServerError)
		ctx.Abort()
	}

	userconste := database.GetUserConstellationData(user)

	ctx.JSON(200, userconste)
}
