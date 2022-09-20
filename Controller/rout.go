package controller

import (
	model "backend/Model"
	"backend/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login", model.Userlogin)

	AuthUserGroup := r.Group("/auth")
	AuthUserGroup.Use(middleware.LoginCheck)
	{

		//AuthUserGroup.GET("/OrigStar/Get", model.GetStarData)    //　オリジナル星座のデータをレスポンスするAPI
		AuthUserGroup.GET("/UserItem/Get", model.GetUserItem)    // ユーザーのアイテムデータをレスポンスするAPI
		AuthUserGroup.POST("/UserItem/Post", model.PostUserItem) // ユーザーのアイテムデータ更新内容を受け取るAPI
		//AuthUserGroup.GET("/Quize/Get", model.QuizeGet)          // クイズデータをランダムに返すAPI
		//AuthUserGroup.GET("/Pub", model.Pub)                     // サーバー側からのイベントをレスポンスするAPI
		//AuthUserGroup.POST("/OrigStar/Post", model.PostStarData) //オリジナル星座データを登録するAPI
	}

	return r
}
