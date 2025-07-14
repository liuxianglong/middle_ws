package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "middle/app/http/api/auth/v1"
	"middle/internal/service"
)

func (s *sAuth) Login(ctx context.Context, req *v1.LoginReq) {
	//如果是没登录，则跳到前端页
	user := service.BizCtx().Get(ctx).User
	if !user.IsLogin() {
		//@todo 前端页缺失，可以先跳到jump页
		frontUrlVar, _ := g.Cfg().Get(ctx, "server.frontLoginUrl")

		g.RequestFromCtx(ctx).Response.RedirectTo(frontUrlVar.String(), 302)
		return
	}

	//如果是登录的，则直接跳到配置的回调URL
	s.loginSucJumpCallback(ctx, req.SSOAuth, user.UID)
}
