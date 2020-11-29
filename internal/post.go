package internal

import (
	"github.com/gin-gonic/gin"
	"log"
	"lost_and_found/model"
	"lost_and_found/service"
	"lost_and_found/util"
)

type CommentInfo struct {
	Kind    int    `json:"kind"`
	PostId  uint   `json:"post_id"`
	Comment string `json:"comment"`
}
type ListRequest struct {
	Start int `json:"start"`
	End   int `json:"end"`
	Kind  int `json:"kind"`
}

func PostThing(ctx *gin.Context) {
	get, _ := ctx.Get("user")
	user := get.(model.User)
	var pi service.PostInfo
	err := ctx.ShouldBindJSON(&pi)
	if err != nil {
		log.Println("28", err)
		msg := getRespMsgByCode(ParamWrong, err)
		definedResp(200, ctx, msg)
		return
	}
	pi.UserID = user.ID
	id := service.Talk(pi)
	msg := getRespMsgByCode(Success, map[string]uint{"id": id})
	definedResp(200, ctx, msg)
}
func PostPicture(ctx *gin.Context) {
	strid := ctx.PostForm("id")
	strkind := ctx.PostForm("kind")
	kind := util.StringPtrInt(strkind)
	id := util.StringPtrUint(strid)
	form, err := ctx.MultipartForm()
	if err != nil {
		log.Println("28", err)
		msg := getRespMsgByCode(ParamWrong, err)
		definedResp(200, ctx, msg)
		return
	}
	files := form.File["file"]
	for _, file := range files {
		service.SaveFile(file, kind, id)
	}
	definedSuccess(ctx)
}
func ShowPostDetail(ctx *gin.Context) {
	strkind := ctx.PostForm("kind")
	strid := ctx.PostForm("id")
	id := util.StringPtrUint(strid)
	kind := util.StringPtrInt(strkind)
	post := service.ShowPostDetail(kind, id)
	msg := getRespMsgByCode(Success, post)
	definedResp(200, ctx, msg)
}
func PushComment(ctx *gin.Context) {
	get, _ := ctx.Get("user")
	user := get.(model.User)
	var ci CommentInfo
	err := ctx.ShouldBindJSON(&ci)
	if err != nil {
		msg := getRespMsgByCode(ParamWrong, err)
		definedResp(200, ctx, msg)
		return
	}
	service.Comment(user.ID, ci.PostId, ci.Kind, ci.Comment)
	definedSuccess(ctx)
}
func ListOfPost(ctx *gin.Context) {
	var lq ListRequest
	_ = ctx.ShouldBindJSON(&lq)
	posts := service.FindPosts(lq.Start, lq.End, lq.Kind)
	msg := getRespMsgByCode(Success, posts)
	definedResp(200, ctx, msg)
}
func FindPost(ctx *gin.Context) {
	var t service.Tag
	_ = ctx.ShouldBindJSON(&t)
	posts := service.SelectPost(t)
	msg := getRespMsgByCode(Success, posts)
	definedResp(200, ctx, msg)
}
