package model

import (
	"github.com/gorilla/websocket"
	"middle/internal/consts"
)

type WSConn struct {
	//ws连接
	Socket *websocket.Conn
	//停止信号
	StopFlag chan bool
	// 消息发送队列
	MsgSendQueue chan []byte
	// 心跳超时时间
	HbTime int32
	// 用户上次心跳时间
	HeartbeatTime uint64
	// 首次连接事件时间
	FirstTime  uint64
	IsVerified int32
	Session    *Session
}

type Session struct {
	Uid int64
}

func NewWSConn(socket *websocket.Conn, firstTime uint64) (client *WSConn) {
	client = &WSConn{
		Socket:        socket,
		MsgSendQueue:  make(chan []byte, 50),
		FirstTime:     firstTime,
		HeartbeatTime: firstTime,
		HbTime:        consts.WSPongWait,
		StopFlag:      make(chan bool, consts.WSNeedStopGoroutineNum),
	}
	return
}
