package model

import (
	"backend/database"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Userlogin(ctx *gin.Context) {
	reqUser := database.User{}
	//クライアントからのjsonデータをユーザー構造体にbinding
	err := ctx.BindJSON(reqUser)

	if err != nil {
		ctx.Status(http.StatusBadRequest)
	} else {

		pass := reqUser.Password
		user := database.GetUserData(reqUser)
		hashPass := user.Password

		err = bcrypt.CompareHashAndPassword([]byte(hashPass), []byte(pass))

		if err != nil {
			ctx.Status(http.StatusBadRequest)
		} else {
			session := sessions.Default(ctx)

			session.Set("loginUser", reqUser)
			session.Save()

			ctx.Status(http.StatusOK)
		}
	}
}
