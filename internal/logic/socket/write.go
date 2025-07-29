package socket

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"middle/internal/service"
	"time"
)

func (s *sSocket) sendMsg(ctx context.Context, msg []byte) {
	if s == nil {
		return
	}
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(ctx, "sSocket.SendMsg panic: %v", err)
		}
	}()
	s.wsConn.MsgSendQueue <- msg
}

// write 向客户端写数据
func (s *sSocket) write(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(ctx, "sSocket.write panic: %v", err)
		}
	}()
	defer func() {
		fmt.Println("发送客户端数据 关闭write")
		service.ClientManager().EventUnregister(ctx, s)

		_ = s.wsConn.Socket.Close()
		//fmt.Println("Client发送数据 defer", c)
	}()

	for {
		select {
		case msg, ok := <-s.wsConn.MsgSendQueue:
			if !ok {
				fmt.Println("Client发送数据 关闭连接")
				return
			}
			err := s.wsConn.Socket.SetWriteDeadline(time.Now().Add(time.Duration(1) * time.Second))
			if err != nil {
				g.Log().Warningf(ctx, "sSocket.write SetWriteDeadline err: %v", err)
				s.Crash(ctx)
				continue
			}
			if err := s.wsConn.Socket.WriteMessage(websocket.BinaryMessage, msg); err != nil {
				s.Crash(ctx)
				continue
			}
		case <-s.wsConn.StopFlag:
			return
		}
	}
}
