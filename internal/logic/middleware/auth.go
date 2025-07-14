package middleware

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"middle/internal/consts"
	"middle/internal/service"
	"middle/utility/code"
)

func (s *sMiddleware) Auth(r *ghttp.Request) {
	user := service.BizCtx().Get(r.GetCtx()).User
	loginFlag := false
	if user.IsLogin() {
		//查看session是否对得上
		authUserSessioncacheKey := fmt.Sprintf(consts.AuthUserSession, user.UID)
		sessionIdVar, _ := g.Redis().Get(r.GetCtx(), authUserSessioncacheKey)
		sessionId := sessionIdVar.String()
		existSessionId, _ := r.Session.Id()
		if existSessionId != "" && existSessionId == sessionId {
			loginFlag = true
		}
	}
	if loginFlag {
		r.Middleware.Next()
	} else {
		r.Cookie.Remove(consts.CookieName)
		r.SetError(code.CodeError.New(r.GetCtx(), code.SystemNeedLogin))
	}
	return
}
