package middleware

import (
	"backend/database"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/koron/go-dproxy"
)

func LoginCheck(ctx *gin.Context) {
	session := sessions.Default(ctx)

	loginUserJson, err := dproxy.New(session.Get("loginUser")).String()

	if err != nil {
		ctx.Status(http.StatusUnauthorized)
		log.Println(err)
		log.Println("dproxy")
		log.Println(session.Get("loginUser"))
		ctx.Abort()
	} else {
		var UserInfo database.User

		err := json.Unmarshal([]byte(loginUserJson), &UserInfo)

		if err != nil {
			ctx.Status(http.StatusUnauthorized)
			log.Println("unmarshal")
			ctx.Abort()
		} else {
			ctx.Next()
		}
	}
}
