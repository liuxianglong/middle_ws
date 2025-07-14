package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"middle/internal/consts"
)

func (s *sAuth) Logout(ctx context.Context) {
	//@todo 看是否需要通知其他系统
	g.RequestFromCtx(ctx).Cookie.Remove(consts.CookieName)
}
