package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/OrigStar/Get/", StarData())       //　オリジナル星座のデータをレスポンスするAPI
	r.GET("/UserItem/Get/", GetUserItem())    // ユーザーのアイテムデータをレスポンスするAPI
	r.POST("/UserItem/Post/", PostUserItem()) // ユーザーのアイテムデータ更新内容を受け取るAPI
	r.GET("/Quize/Get/", QuizeGet())          // クイズデータをランダムに返すAPI
	r.POST("Auth/", UserAuth())               // ユーザー認証API
	r.GET("/Pub/", Pub())                     // サーバー側からのイベントをレスポンスするAPI

	return r
}
