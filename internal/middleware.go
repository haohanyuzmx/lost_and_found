package internal

import (
	"github.com/gin-gonic/gin"
	"lost_and_found/service"
	"lost_and_found/util"
	"net/http"
	"strings"
)

func TokenVerify(ctx *gin.Context) {
	tokenHeader := ctx.GetHeader("Authorization")
	if tokenHeader == "" || !strings.HasPrefix(tokenHeader, "Bearer ") {
		definedResp(http.StatusForbidden, ctx, AuthorizedError)
		ctx.Abort()
		return
	}
	token := strings.TrimPrefix(tokenHeader, "Bearer ")
	jwt, err := util.Check(token)
	if err != nil {
		definedResp(http.StatusForbidden, ctx, AuthorizedError)
		ctx.Abort()
		return
	}
	/*e := jwt.Payload.Exp
	exp, _ := strconv.Atoi(e)
	if int64(exp)-time.Now().Unix() < 0 {
		definedResp(http.StatusForbidden, ctx, AuthorizedError)
		ctx.Abort()
		return
	}*/
	id := jwt.Payload.Id
	user := service.GetUser(uint(id))
	ctx.Set("user", user)
	ctx.Next()
}
