package internal

import (
	"github.com/gin-gonic/gin"
	"lost_and_found/model"
	"lost_and_found/service"
	"lost_and_found/util"
)

type UserBase struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
type RealBase struct {
	PhoneNum int64  `json:"phone_num"`
	StuNum   int64  `json:"stu_num"`
	RealName string `json:"real_name"`
	QQNum    int64  `json:"qq_num"`
}

func Register(ctx *gin.Context) {
	var ub UserBase
	err := ctx.ShouldBindJSON(&ub)
	if err != nil {
		msg := getRespMsgByCode(ParamWrong, err)
		definedResp(200, ctx, msg)
	}
	err = service.CreatUser(ub.Username, ub.Password)
	if err != nil {
		msg := getRespMsgByCode(DataError, err)
		definedResp(200, ctx, msg)
	}
	msg := getRespMsgByCode(Success, ub)
	definedResp(200, ctx, msg)
}
func Login(ctx *gin.Context) {
	var ub UserBase
	err := ctx.ShouldBindJSON(&ub)
	if err != nil {
		msg := getRespMsgByCode(ParamWrong, err)
		definedResp(200, ctx, msg)
		return
	}
	user, err := service.CheckUser(ub.Username, ub.Password)
	if err != nil {
		msg := getRespMsgByCode(DataError, err)
		definedResp(200, ctx, msg)
		return
	}
	jwt := util.NewJWT(user.ID)
	msg := getRespMsgByCode(Success, jwt.Token)
	definedResp(200, ctx, msg)
}
func UploadPortrait(ctx *gin.Context) {
	get, _ := ctx.Get("user")
	user := get.(model.User)
	file, err := ctx.FormFile("portrait")
	if err != nil {
		msg := getRespMsgByCode(ServiceError, err)
		definedResp(200, ctx, msg)
		return
	}
	service.SaveFile(file, 0, user.ID)
	user.GetPortrait()
	ctx.Set("user", user)
	definedSuccess(ctx)
}
func UpdateReal(ctx *gin.Context) {
	get, _ := ctx.Get("user")
	user := get.(model.User)
	var rb RealBase
	err := ctx.ShouldBindJSON(&rb)
	if err != nil {
		msg := getRespMsgByCode(ParamWrong, err)
		definedResp(200, ctx, msg)
	}
	{
		user.RealName = rb.RealName
		user.PhoneNum = rb.PhoneNum
		user.QQNum = rb.QQNum
		user.StuNum = rb.StuNum
	}
	service.SaveUser(&user)
	ctx.Set("user", user)
	definedSuccess(ctx)
}
func ShowData(ctx *gin.Context) {
	strid := ctx.Query("id")
	id := util.StringPtrUint(strid)
	user := service.GetUser(id)
	msg := getRespMsgByCode(Success, user)
	definedResp(200, ctx, msg)
}
func ShowMyData(ctx *gin.Context) {
	get, _ := ctx.Get("user")
	user := get.(model.User)
	msg := getRespMsgByCode(Success, user)
	definedResp(200, ctx, msg)
}
func Push(ctx *gin.Context) {
	get, _ := ctx.Get("user")
	user := get.(model.User)
	thing := service.FindMyThing(user)
	msg := getRespMsgByCode(Success, thing)
	definedResp(200, ctx, msg)
}
func Portrait(ctx *gin.Context) {
	strid := ctx.PostForm("id")
	id := util.StringPtrUint(strid)
	user := service.GetUser(id)
	msg := getRespMsgByCode(Success, map[string]string{"portrait": user.Portrait.Uri})
	definedResp(200, ctx, msg)
}
