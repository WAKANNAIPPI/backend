package controller

import (
	"log"
	"net/http"
	model "src/Model"
	"src/middleware"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func GetRouter() *gin.Engine {

	//gin,session,melodyのセットアップ
	r := gin.Default()
	m := melody.New()

	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.GET("/flag", func(ctx *gin.Context) { //定期イベント処理用のAPI
		//クライアントIpが127.0.0.1ならばブロードキャスト実行
		cliantIp := ctx.ClientIP()

		if cliantIp == "127.0.0.1" {
			//イベントフラグをws接続中の全クライアントに送信、値はuint8の「1」
			m.Broadcast([]byte("1"))
		} else {
			ctx.Status(http.StatusUnauthorized)
			ctx.String(401, "ここにはlocalhost以外アクセス出来ません")
		}
	})

	r.GET("/ws/Event", func(ctx *gin.Context) { //ws確立用のAPI

		err := m.HandleRequest(ctx.Writer, ctx.Request)

		if err != nil {
			log.Println(err)
		}
	})

	r.POST("/login", model.Userlogin) //loginAPI

	//ここから先のAPIはログイン済みのユーザでないとアクセスできない。
	//sessionの使用が必要
	AuthUserGroup := r.Group("/auth")
	AuthUserGroup.Use(middleware.LoginCheck)
	{

		AuthUserGroup.GET("/OrigConste/Get", model.GetConsteData)    //　オリジナル星座のデータをレスポンスするAPI
		AuthUserGroup.GET("/UserItem/Get", model.GetUserItem)        // ユーザーのアイテムデータをレスポンスするAPI
		AuthUserGroup.POST("/UserItem/Post", model.PostUserItem)     // ユーザーのアイテムデータ更新内容を受け取るAPI
		AuthUserGroup.GET("/Quize/Get", model.QuizeGet)              // クイズデータをランダムに返すAPI
		AuthUserGroup.POST("/OrigConste/Post", model.PostConsteData) //オリジナル星座データを登録するAPI
	}

	return r
}
