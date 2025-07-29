package socket

import (
	"context"
	"middle/internal/consts"
	"middle/internal/model"
	"middle/internal/service"
)

type (
	sSocket struct {
		wsConn *model.WSConn
	}
)

func init() {
	service.RegisterSocket(New())
}

func New() service.ISocket {
	return &sSocket{}
}

func (s *sSocket) GetConn(ctx context.Context) *model.WSConn {
	return s.wsConn
}

// heartbeat 心跳时间
func (c *sSocket) heartbeat(currentTime uint64) {
	c.wsConn.HeartbeatTime = currentTime

	return
}

func (s *sSocket) isLogin(ctx context.Context) bool {
	if s.wsConn == nil {
		return false
	}
	if s.wsConn.Session == nil {
		return false
	}
	if s.wsConn.Session.Uid == 0 {
		return false
	}

	return true
}

func (s *sSocket) IsHeartbeatTimeout(ctx context.Context, currentTime uint64) (timeout bool) {
	if s.wsConn.HeartbeatTime+consts.WSPongWait <= currentTime {
		timeout = true
	}
	return
}
