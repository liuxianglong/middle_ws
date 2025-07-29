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
			srvRouter := service.SrvRouter()
			srvRouter.InitRouter(ctx)
			go srvRouter.Lookup(ctx)

			clientManager := service.ClientManager()
			go clientManager.Start(ctx)
			go clientManager.ClearTimeoutConnections(ctx)
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
