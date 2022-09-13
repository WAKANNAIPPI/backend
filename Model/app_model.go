package model

import (
	"backend/database"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
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
			session.Set("loginUser", user)
			session.Save()
		}
	}
}
