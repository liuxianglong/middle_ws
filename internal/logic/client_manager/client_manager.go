package client_manager

import (
	"context"
	"fmt"
	"middle/internal/service"
	"sync"
	"time"
)

type (
	sClientManager struct {
		Clients     map[service.ISocket]bool
		Users       map[int64]service.ISocket
		Register    chan service.ISocket // 连接连接处理
		ClientsLock sync.RWMutex         // 读写锁
		UsersLock   sync.RWMutex         // 读写锁
		Unregister  chan service.ISocket // 连接连接处理
		Login       chan service.ISocket // 用户登录处理
	}
)

func init() {
	service.RegisterClientManager(New())
}

func New() service.IClientManager {
	return &sClientManager{
		Clients:    make(map[service.ISocket]bool),
		Users:      make(map[int64]service.ISocket),
		Register:   make(chan service.ISocket, 100),
		Unregister: make(chan service.ISocket, 100),
	}
}

//func (s *sClientManager) InitClientManager(ctx context.Context) (clientManager service.IClientManager) {
//	clientManager = &sClientManager{
//		Clients:    make(map[service.ISocket]bool),
//		Users:      make(map[int64]service.ISocket),
//		Register:   make(chan service.ISocket, 100),
//		Unregister: make(chan service.ISocket, 100),
//	}
//	return
//}

func (s *sClientManager) Start(ctx context.Context) {
	for {
		select {
		case conn := <-s.Register:
			s.addClient(ctx, conn)
		case conn := <-s.Login:
			s.addUsers(ctx, conn)
		case conn := <-s.Unregister:
			// 断开连接事件
			s.unRegister(ctx, conn)
		}
	}
}

func (s *sClientManager) EventRegister(ctx context.Context, conn service.ISocket) {
	s.Register <- conn
}

func (s *sClientManager) EventLogin(ctx context.Context, conn service.ISocket) {
	s.Login <- conn
}

func (s *sClientManager) EventUnregister(ctx context.Context, conn service.ISocket) {
	s.Unregister <- conn
}

func (s *sClientManager) addClient(ctx context.Context, conn service.ISocket) {
	s.ClientsLock.RLock()
	defer s.ClientsLock.RUnlock()
	s.Clients[conn] = true
}

func (s *sClientManager) delClient(ctx context.Context, conn service.ISocket) {
	s.ClientsLock.RLock()
	defer s.ClientsLock.RUnlock()

	delete(s.Clients, conn)
}

func (s *sClientManager) addUsers(ctx context.Context, conn service.ISocket) {
	s.UsersLock.RLock()
	defer s.UsersLock.RUnlock()

	// 连接存在，在添加
	wsConn := conn.GetConn(ctx)
	s.Users[wsConn.Session.Uid] = conn
}

func (s *sClientManager) delUsers(ctx context.Context, conn service.ISocket) bool {
	s.UsersLock.RLock()
	defer s.UsersLock.RUnlock()
	wsConn := conn.GetConn(ctx)

	if _, ok := s.Users[wsConn.Session.Uid]; ok {
		delete(s.Users, wsConn.Session.Uid)
		return true
	}
	return false
}

func (s *sClientManager) unRegister(ctx context.Context, conn service.ISocket) {
	s.delClient(ctx, conn)
	deleteResult := s.delUsers(ctx, conn)
	if deleteResult == false {
		// 不是当前连接的客户端
		return
	}
}

func (s *sClientManager) ClearTimeoutConnections(ctx context.Context) {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			currentTime := uint64(time.Now().Unix())
			for client := range s.getClients() {
				if client.IsHeartbeatTimeout(ctx, currentTime) {
					wsConn := client.GetConn(ctx)
					var uid int64
					var heartbeatTime uint64
					if wsConn != nil {
						heartbeatTime = wsConn.HeartbeatTime
						if wsConn.Session != nil {
							uid = wsConn.Session.Uid
						}

					}
					fmt.Println("心跳时间超时 关闭连接", uid, heartbeatTime)
					client.Crash(ctx)
				}
			}
		}
	}
}

func (s *sClientManager) getClients() (clients map[service.ISocket]bool) {
	clients = make(map[service.ISocket]bool)
	s.clientsRange(func(client service.ISocket, value bool) (result bool) {
		clients[client] = value
		return true
	})
	return
}

// clientsRange 遍历
func (s *sClientManager) clientsRange(f func(client service.ISocket, value bool) (result bool)) {
	s.ClientsLock.RLock()
	defer s.ClientsLock.RUnlock()
	for key, value := range s.Clients {
		result := f(key, value)
		if result == false {
			return
		}
	}
	return
}
