// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
)

type (
	IClientManager interface {
		Start(ctx context.Context)
		EventRegister(ctx context.Context, conn ISocket)
		EventLogin(ctx context.Context, conn ISocket)
		EventUnregister(ctx context.Context, conn ISocket)
		ClearTimeoutConnections(ctx context.Context)
	}
)

var (
	localClientManager IClientManager
)

func ClientManager() IClientManager {
	if localClientManager == nil {
		panic("implement not found for interface IClientManager, forgot register?")
	}
	return localClientManager
}

func RegisterClientManager(i IClientManager) {
	localClientManager = i
}
