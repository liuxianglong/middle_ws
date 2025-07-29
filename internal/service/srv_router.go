// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"middle/internal/model"
)

type (
	ISrvRouter interface {
		InitRouter(ctx context.Context)
		GetRouter(ctx context.Context, routerName string) *model.SrvRouterServer
		// Lookup 监听,更新配置
		Lookup(ctx context.Context)
	}
)

var (
	localSrvRouter ISrvRouter
)

func SrvRouter() ISrvRouter {
	if localSrvRouter == nil {
		panic("implement not found for interface ISrvRouter, forgot register?")
	}
	return localSrvRouter
}

func RegisterSrvRouter(i ISrvRouter) {
	localSrvRouter = i
}
