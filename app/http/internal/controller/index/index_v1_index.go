package index

import (
	"context"
	"middle/app/http/api/index/v1"
	"middle/internal/service"
)

func (c *ControllerV1) Index(ctx context.Context, req *v1.IndexReq) (res *v1.IndexRes, err error) {
	//g.RequestFromCtx(ctx).
	//writer := g.RequestFromCtx(ctx).Response.Writer
	//g.RequestFromCtx(ctx).Request
	//
	//wsUpGrader := websocket.Upgrader{
	//	CheckOrigin: func(r *http.Request) bool { //是否允许跨域
	//		return true
	//	},
	//	ReadBufferSize:  consts.WSReadBufferSize,
	//	WriteBufferSize: consts.WSWriteBufferSize,
	//	Error: func(w http.ResponseWriter, r *http.Request, status int, reason error) {
	//
	//	},
	//}
	//
	//ws, err := wsUpGrader.Upgrade(g.RequestFromCtx(ctx).Response.Writer, g.RequestFromCtx(ctx).Request, nil)
	////ws, err := wsUpGrader.Upgrade(r.Response.Writer, r.Request, nil)
	//if err != nil {
	//	g.Log().Errorf(ctx, "websocket upgrade err:%v", err)
	//	//r.Response.Write(err.Error())
	//	return
	//}
	//defer ws.Close()
	//
	//ws.SetReadLimit(consts.WSMaxMessageSize)
	//err = ws.SetReadDeadline(time.Now().Add(time.Duration(consts.WSPongWait) * time.Second))
	//if err != nil {
	//	g.Log().Errorf(ctx, "set read deadline err:%v", err)
	//	return nil, err
	//}
	//
	//for {
	//	msgType, msg, err := ws.ReadMessage()
	//	if err != nil {
	//		break
	//	}
	//	g.Log().Infof(ctx, "received msgType:%d message: %s", msgType, string(msg))
	//	if err = ws.WriteMessage(msgType, []byte("ddd")); err != nil {
	//		break
	//	}
	//}
	//g.Log().Info(ctx, "websocket connection closed")
	service.Socket().Conn(ctx)
	return nil, nil
}
