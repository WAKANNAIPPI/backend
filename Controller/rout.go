package controller

import (
	model "backend/Model"
	"backend/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {

	//gin,melody,sessionのセットアップ
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/flag", model.EventFlag)   //定期イベント処理用のAPI
	r.POST("/login", model.Userlogin) //loginAPI

	AuthUserGroup := r.Group("/auth")
	AuthUserGroup.Use(middleware.LoginCheck)
	{

		AuthUserGroup.GET("/OrigConste/Get", model.GetConsteData)    //　オリジナル星座のデータをレスポンスするAPI
		AuthUserGroup.GET("/UserItem/Get", model.GetUserItem)        // ユーザーのアイテムデータをレスポンスするAPI
		AuthUserGroup.POST("/UserItem/Post", model.PostUserItem)     // ユーザーのアイテムデータ更新内容を受け取るAPI
		AuthUserGroup.GET("/Quize/Get", model.QuizeGet)              // クイズデータをランダムに返すAPI
		AuthUserGroup.GET("/ws/Event", model.WsConnect)              // サーバー側からのイベントをレスポンスするAPI
		AuthUserGroup.POST("/OrigConste/Post", model.PostConsteData) //オリジナル星座データを登録するAPI
	}

	return r
}
