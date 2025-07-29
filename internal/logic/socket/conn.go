package socket

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gorilla/websocket"
	"middle/internal/consts"
	"middle/internal/model"
	"middle/internal/service"
	"time"
)

func (s *sSocket) Conn(ctx context.Context) {

	ws, err := (&websocket.Upgrader{
		//CheckOrigin: func(r *http.Request) bool {
		//	fmt.Println("升级协议", "ua:", r.Header["User-Agent"], "referer:", r.Header["Referer"])
		//	return true
		//},
		ReadBufferSize:  consts.WSReadBufferSize,
		WriteBufferSize: consts.WSWriteBufferSize,
	}).Upgrade(g.RequestFromCtx(ctx).Response.Writer, g.RequestFromCtx(ctx).Request, nil)

	if err != nil {
		g.Log().Errorf(ctx, "websocket upgrade err:%v", err)

		return
	}

	ws.SetReadLimit(consts.WSMaxMessageSize + consts.WSMaxHeaderSize + 1024)
	err = ws.SetReadDeadline(time.Now().Add(time.Duration(consts.WSPongWait) * time.Second))
	if err != nil {
		g.Log().Errorf(ctx, "set read deadline err:%v", err)
		return
	}
	currentTime := uint64(time.Now().Unix())
	s.wsConn = model.NewWSConn(ws, currentTime)
	go s.read(ctx)
	go s.write(ctx)
	service.ClientManager().EventRegister(ctx, s)
}
