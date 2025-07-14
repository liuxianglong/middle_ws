package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gsession"
	"middle/app/http/internal/controller/auth"
	"middle/app/http/internal/controller/hello"
	"middle/app/http/internal/controller/index"
	"middle/internal/consts"
	"middle/internal/service"
	"time"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			s := g.Server()
			s.SetSessionIdName(consts.CookieName)

			service.Cache().RedisRegister(ctx)
			//将session绑定到redis
			s.SetSessionMaxAge(consts.SessionExpireTime * time.Second)
			sessionStorage := gsession.NewStorageRedis(g.Redis())
			s.SetSessionStorage(sessionStorage)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
				)
				group.Middleware(service.Middleware().HandleResponse)
				group.Bind(
					auth.NewV1(),
					hello.NewV1(),
				)
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().SsoTokenResponse,
					)
					group.Bind(auth.NewV1().Token)
				})

				// 需要登录
				group.Group("/", func(group *ghttp.RouterGroup) {
					group.Middleware(
						service.Middleware().Auth,
					)
					group.Bind(
						index.NewV1(),
					)
				})
			})
			s.Run()
			return nil
		},
	}
)
