package controller

import (
	model "src/Model"
	"src/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	//gin,sessionのセットアップ
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/flag", model.EventFlag)     //定期イベント処理用のAPI
	r.GET("/ws/Event", model.WsConnect) //ws確立用のAPI
	r.POST("/login", model.Userlogin)   //loginAPI

	//ここから先のAPIはログイン済みのユーザでないとアクセスできない。
	//sessionの使用が必要
	AuthUserGroup := r.Group("/auth")
	AuthUserGroup.Use(middleware.LoginCheck)
	{

		AuthUserGroup.GET("/OrigConste/Get", model.GetConsteData)        //　オリジナル星座のデータをレスポンスするAPI
		AuthUserGroup.GET("/UserItem/Get", model.GetUserItem)            // ユーザーのアイテムデータをレスポンスするAPI
		AuthUserGroup.POST("/UserItem/Post", model.PostUserItem)         // ユーザーのアイテムデータ更新内容を受け取るAPI
		AuthUserGroup.GET("/Quize/Get", model.QuizeGet)                  // クイズデータをランダムに返すAPI
		AuthUserGroup.POST("/OrigConste/Post", model.PostConsteData)     //オリジナル星座データを登録するAPI
		AuthUserGroup.POST("/OrigConste/Update", model.UpdateConsteData) //オリジナル星座の編集API
		AuthUserGroup.GET("/UserInfo/Get", model.UserInfoGet)
	}

	return r
}
