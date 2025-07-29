package socket

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"io"
	"math/bits"
	"middle/internal/consts"
	"middle/internal/service"
	"middle/protobuf/pb"
	"time"
)

// read 读客户端消息
func (s *sSocket) read(ctx context.Context) {
	defer func() {
		if err := recover(); err != nil {
			g.Log().Errorf(ctx, "sSocket.read panic: %v", err)
		}
	}()
	defer func() {
		fmt.Println("读取客户端数据 关闭send")
		close(s.wsConn.MsgSendQueue)
	}()

	for {
		select {
		case <-s.wsConn.StopFlag:
			return
		default:
			if s.wsConn.Socket == nil {
				g.Log().Errorf(ctx, "sSocket.Socket is nil")
			}
			messageType, reader, err := s.wsConn.Socket.NextReader()
			if err != nil {
				//fmt.Println("读取客户端数据 错误", c.Addr, err)
				g.Log().Errorf(ctx, "sSocket.read ReadMessage err=%v", err)
				s.Crash(ctx)
				continue
			}

			// 处理程序
			if messageType != websocket.BinaryMessage { //只接受二进制
				g.Log().Error(ctx, "sSocket.read ReadMessage non-binary message ")
				continue
			}
			err = s.wsConn.Socket.SetReadDeadline(time.Now().Add(time.Duration(consts.WSPongWait) * time.Second))
			if err != nil {
				g.Log().Errorf(ctx, "sSocket.read SetReadDeadline err=%v", err)
				s.Crash(ctx)
				return
			}

			request, err := s.parseMessage(ctx, reader)
			if err != nil {
				g.Log().Warningf(ctx, "sSocket.read ReadMessage err=%v", err)
				continue
			}
			//g.Log().Infof(ctx, "sSocket.read ReadMessage: %+v", request)
			if s.isLogin(ctx) {
				if request.Head.Cmd == "login" {
					//@todo 需要调接口做登录处理
					s.wsConn.IsVerified = 1
					currentTime := uint64(time.Now().Unix())
					s.heartbeat(currentTime)
					service.ClientManager().EventLogin(ctx, s)
				}
			} else {
				if request.Head.Cmd == "heartbeat" {
					currentTime := uint64(time.Now().Unix())
					s.heartbeat(currentTime)
				}
			}

			//buf := new(bytes.Buffer)
			//_ = binary.Write(buf, binary.BigEndian, s.wsConn.IsVerified)
			//by := buf.Bytes()
			//s.sendMsg(ctx, []byte("hello"))
			//ProcessData(c, message)
			//fmt.Println(s.wsConn.IsVerified)
		}

	}
}
func (s *sSocket) parseMessage(ctx context.Context, reader io.Reader) (request *pb.WSRequest, err error) {
	//读取头部长度前缀（4字节）
	headerSizeBuf := make([]byte, consts.WSHeaderLengthSize)
	if _, err := io.ReadFull(reader, headerSizeBuf); err != nil {
		return nil, fmt.Errorf("读取头部长度失败: %w", err)
	}
	headerSize := binary.BigEndian.Uint32(headerSizeBuf)
	if headerSize > consts.WSMaxHeaderSize {
		return nil, fmt.Errorf("头部大小超过限制")
	}
	// 3. 读取头部数据
	headerData := make([]byte, headerSize)
	if _, err := io.ReadFull(reader, headerData); err != nil {
		return nil, fmt.Errorf("读取头部数据失败: %w", err)
	}

	// 4. 解析头部
	head := &pb.WSRequestHead{}
	if err := proto.Unmarshal(headerData, head); err != nil {
		return nil, fmt.Errorf("无效的头部数据: %v", err)
	}
	if head.Size < 0 {
		return nil, fmt.Errorf("无效的内容大小")
	}
	if head.Size > consts.WSMaxMessageSize {
		return nil, fmt.Errorf("内容大小超过限制")
	}

	//@todo 测试时关
	//if !s.checkHead(ctx, head) {
	//	return nil, fmt.Errorf("checkHead error")
	//}

	// 5. 解析路由
	routerSrv := service.SrvRouter().GetRouter(ctx, head.Cmd)
	if routerSrv == nil {
		return nil, fmt.Errorf("找不到路由")
	}

	// 6. 读取内容
	content := make([]byte, head.Size)
	if _, err := io.ReadFull(reader, content); err != nil {
		return nil, fmt.Errorf("读取内容失败: %w", err)
	}
	return &pb.WSRequest{
		Head: head,
		Body: content,
	}, nil
}
func (s *sSocket) checkHead(ctx context.Context, head *pb.WSRequestHead) bool {
	//check := int64(msg.Head.Tm) + int64(msg.Head.Rt)
	tm, _ := s.decodeVarInt(head.Tm)
	rt, _ := s.decodeVarInt(head.Rt)
	free1, _ := s.decodeVarInt(head.Free1)
	free2, _ := s.decodeVarInt(head.Free2)
	id := head.Id
	uid := head.Uid
	sendTime := head.SendTime
	size := int64(head.Size)
	check := tm + rt + free1 + free2 + id + uid + sendTime + size

	headCheck, _ := s.decodeVarInt(head.Check)
	if len(head.HeadPad) == 1 && head.HeadPad[0] == '$' ||
		len(head.EndPad) == 1 && head.EndPad[0] == '$' || headCheck != check {
		return false
	}
	return true
}

func (s *sSocket) decodeVarInt(data []byte) (int64, int) {
	var x uint64
	for s := 0; s < bits.Len64(x); s += 7 {
		b := data[0]
		data = data[1:]
		x |= uint64(b&0x7F) << s
		if b < 0x80 {
			return int64(x), s/7 + 1
		}
	}
	return 0, 0
}

func (s *sSocket) Crash(ctx context.Context) {
	if s.wsConn.Session != nil {
		service.ClientManager().EventUnregister(ctx, s)
		s.wsConn.Session = nil
	}
	s.stop(ctx)
}

func (s *sSocket) stop(ctx context.Context) {
	for i := 0; i < consts.WSNeedStopGoroutineNum; i++ {
		s.wsConn.StopFlag <- true
	}
}
