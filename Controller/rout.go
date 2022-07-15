package controller

import "github.com/gin-gonic/gin"

func GetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("Data/star/", StarData())
	r.GET("Data/UserItem/", UserData())
	r.GET("Data/Quize/", QuizeData())

	return r
}
