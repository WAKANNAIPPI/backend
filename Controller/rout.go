package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func GetRouter() *gin.Engine {
	r := gin.Default()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.POST("/login/post", model.Userlogin())

	AuthUserGroup := r.Group("/auth")
	AuthUserGroup.Use(middleware.LoginCheck())
	{

		r.GET("/OrigStar/Get/", model.GetStarData())    //　オリジナル星座のデータをレスポンスするAPI
		r.GET("/UserItem/Get/", model.GetUserItem())    // ユーザーのアイテムデータをレスポンスするAPI
		r.POST("/UserItem/Post/", model.PostUserItem()) // ユーザーのアイテムデータ更新内容を受け取るAPI
		r.GET("/Quize/Get/", model.QuizeGet())          // クイズデータをランダムに返すAPI
		r.GET("/Pub/", model.Pub())                     // サーバー側からのイベントをレスポンスするAPI
		r.POST("/OrigStar/Post/", model.PostStarData())
	}

	return r
}
