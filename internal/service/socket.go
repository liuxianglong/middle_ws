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
	ISocket interface {
		Conn(ctx context.Context)
		Crash(ctx context.Context)
		GetConn(ctx context.Context) *model.WSConn
		IsHeartbeatTimeout(ctx context.Context, currentTime uint64) (timeout bool)
	}
)

var (
	localSocket ISocket
)

func Socket() ISocket {
	if localSocket == nil {
		panic("implement not found for interface ISocket, forgot register?")
	}
	return localSocket
}

func RegisterSocket(i ISocket) {
	localSocket = i
}
