package api

import (
	"github.com/gin-gonic/gin"
	"lost_and_found/internal"
)

func User(en *gin.Engine) {
	ug := en.Group("user")
	ug.GET("showdata", internal.ShowData)
	ug.POST("register", internal.Register)
	ug.POST("login", internal.Login)
	ug.POST("portrait", internal.Portrait)
	ug.Use(internal.TokenVerify)
	ug.POST("upportrait", internal.UploadPortrait)
	ug.POST("upreal", internal.UpdateReal)
	ug.GET("showmydata", internal.ShowMyData)
	ug.GET("push", internal.Push)
}
