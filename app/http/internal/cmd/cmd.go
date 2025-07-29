package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"middle/app/http/internal/controller/index"

	//"middle/app/http/internal/controller/index"
	"middle/internal/service"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start ws server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			s := g.Server()

			service.Cache().RedisRegister(ctx)
			service.SrvRouter().InitRouter(ctx)
			go service.SrvRouter().Lookup(ctx)

			service.ClientManager().InitClientManager(ctx)
			go service.ClientManager().Start(ctx)
			go service.ClientManager().ClearTimeoutConnections(ctx)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					service.Middleware().CORS,
				)
				group.Middleware(service.Middleware().HandleResponse)
				group.Bind(
					index.NewV1(),
				)

				// 需要登录
				//group.Group("/", func(group *ghttp.RouterGroup) {
				//	group.Middleware(
				//		service.Middleware().Auth,
				//	)
				//	group.Bind(
				//		index.NewV1(),
				//	)
				//})
			})
			s.Run()
			return nil
		},
	}
)
