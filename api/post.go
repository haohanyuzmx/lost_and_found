package api

import (
	"github.com/gin-gonic/gin"
	"lost_and_found/internal"
)

func Post(en *gin.Engine) {
	pg := en.Group("post")
	pg.POST("postlist", internal.ListOfPost)
	pg.POST("findbytag", internal.FindPost)
	pg.Use(internal.TokenVerify)
	pg.POST("postthing", internal.PostThing)
	pg.POST("postpicture", internal.PostPicture)
	pg.POST("pushcomment", internal.PushComment)
	pg.POST("postdetail", internal.ShowPostDetail)
}
