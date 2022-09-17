package model

import (
	"backend/database"
	"encoding/json"
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
