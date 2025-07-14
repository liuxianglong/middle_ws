package nsq

import (
	"context"
	"github.com/nsqio/go-nsq"
	"middle/internal/consts"
	"middle/library/tool"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"

	"fmt"
	"time"
)

// NsqClient 返回cnsq client 对象
//func NsqClient() *Client {
//	return NewNsqClient()
//}

// NewNsqWorker 实例化 NsqWorker
func NewNsqWorker(topic, channel string, handler NsqHandleMessage) *NsqWorker {
	return &NsqWorker{
		topic:   topic,
		channel: channel,
		handler: handler,
	}
}

// NsqHandleMessage 定义NewNsqWorker的回调方式
type NsqHandleMessage func(message *nsq.Message) error

// NsqWorker 基于topic消费nsd
type NsqWorker struct {
	client  *nsq.Consumer
	topic   string
	channel string
	handler NsqHandleMessage
}

// HandleMessage nsq 基类的回调接口
func (s *NsqWorker) HandleMessage(message *nsq.Message) error {
	//反序列化数据为JSON
	return s.handler(message)
}

// Connect 建立连接函数
func (s *NsqWorker) Connect(ctx context.Context) error {
	var err error
	ip, _ := tool.IP.LocalIPv4s()
	g.Log().Infof(ctx, "localIPv4s = %s", ip)
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second * 30 //设置重连时间
	cfg.HeartbeatInterval = time.Second * 5
	cfg.Hostname = ip
	nsds, err := GetNsdGroup(ctx, s.topic)
	if err != nil {
		return err
	}
	cfg.MaxInFlight = len(nsds)
	g.Log().Infof(ctx, "nsq.Consumer MaxInFlight=%d", cfg.MaxInFlight)
	s.client, err = nsq.NewConsumer(s.topic, s.channel, cfg) // 新建一个消费者
	if err != nil {
		return err
	}
	s.client.SetLogger(nil, 0) //屏蔽系统日志
	s.client.AddHandler(s)     // 添加消费者接口
	s.client.SetBehaviorDelegate(s)
	//建立NSQLookupd连接
	configVars, _ := g.Cfg().Get(ctx, consts.NsqConfigLookupName)
	config := configVars.Strings()
	g.Log().Infof(ctx, "nsq.NsqConfigLookupName %v\n", config)
	if len(config) == 0 {
		panic(gerror.New("empty lookup for nsq"))
	}
	if err := s.client.ConnectToNSQLookupds(config); err != nil {
		g.Log().Infof(ctx, "nsq.NsqConfigLookupName  err = %s\n", err.Error())
		return err
	}
	g.Log().Infof(ctx, "nsq.NsqConfigLookupName %v ok\n", s.topic)

	return nil
}

// DiscoveryFilter 对现在线上的NSD ip 进行转换
// todo... 伴伴环境需要填写具体的转化IP
func (s NsqWorker) Filter(addrs []string) []string {
	values := []string{}
	replace := map[string]string{}
	for i := 0; i < len(addrs); i++ {
		addr := addrs[i]
		if val, ok := replace[addr]; ok {
			values = append(values, val)
		} else {
			values = append(values, addr)
		}
	}
	return values
}

func (s *NsqWorker) Close(ctx context.Context) error {
	if !s.client.IsStarved() {
		s.client.ChangeMaxInFlight(0)
		time.Sleep(time.Second * 1)
	}

	configVar, _ := g.Cfg().Get(ctx, consts.NsqConfigLookupName)
	config := configVar.Strings()
	fmt.Println("disconnect: ", config)
	if len(config) > 0 {
		for _, Host := range config {
			_ = s.client.DisconnectFromNSQLookupd(Host)
		}
	}

	return nil
}

// Stop 停止
func (s *NsqWorker) Stop() {
	s.client.Stop()
}

// ConnectWithConcurrency 建立连接函数(可以指定并发消费的消息数)
func (s *NsqWorker) ConnectWithConcurrency(ctx context.Context, cNum int) error {
	var err error
	ip, _ := tool.IP.LocalIPv4s()
	cfg := nsq.NewConfig()
	cfg.LookupdPollInterval = time.Second * 30 //设置重连时间
	cfg.HeartbeatInterval = time.Second * 5
	cfg.Hostname = ip
	nsds, err := GetNsdGroup(ctx, s.topic)
	if err != nil {
		return err
	}
	cfg.MaxInFlight = len(nsds)
	if cfg.MaxInFlight < cNum {
		cfg.MaxInFlight = cNum
	}
	g.Log().Infof(ctx, "nsq.Consumer MaxInFlight=%d", cfg.MaxInFlight)
	s.client, err = nsq.NewConsumer(s.topic, s.channel, cfg) // 新建一个消费者
	if err != nil {
		return err
	}

	//var logger *log.Logger
	//if gmode.IsDevelop() {
	//	logger = log.New(os.Stdout, "[nsq-log]", log.Ldate|log.Ltime|log.Lshortfile)
	//}

	s.client.SetLogger(nil, nsq.LogLevelDebug) //屏蔽系统日志
	s.client.AddConcurrentHandlers(s, cNum)    // 添加消费者接口
	s.client.SetBehaviorDelegate(s)
	//建立NSQLookupd连接
	configVars, _ := g.Cfg().Get(ctx, consts.NsqConfigLookupName)
	config := configVars.Strings()
	g.Log().Infof(ctx, "nsq.NsqConfigLookupName %v\n", config)
	if len(config) == 0 {
		panic(gerror.New("empty lookup for nsq"))
	}
	if err := s.client.ConnectToNSQLookupds(config); err != nil {
		g.Log().Infof(ctx, "nsq.NsqConfigLookupName  err = %s\n", err.Error())
		return err
	}
	g.Log().Infof(ctx, "nsq.NsqConfigLookupName  ok\n")

	return nil
}

func MessageIDToString(msgId nsq.MessageID) string {
	return string(msgId[:])
}
