package cmd

import (
	"context"
	"fmt"
	"github.com/nsqio/go-nsq"
	"middle/app/job/internal/cmd/nsqcmd"
	"middle/library/cmd"
	nsqc "middle/library/nsq"
	"time"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	Nsq = cNsq{}
	s   = &nsqConsumerServer{}
)

type (
	ConsumerCfg struct { //消费启动模版(方便启动新的消费实例)
		Channel  string
		Num      int
		Consumer IConsumer
	}
)
type IConsumer interface {
	NsqConsumerMember(message *nsq.Message) error
	Wait()
	GetTopic() string
}

type cNsq struct {
	g.Meta `name:"nsq" brief:"nsq consumer"`
}

type cNsqInput struct {
	g.Meta `name:"nsq"`
}

type cNsqOutput struct{}

type nsqConsumerServer struct {
	cmd.StartUp
	List []*nsqc.NsqWorker //所有的消费实例
}

func (c cNsq) Index(ctx context.Context, in cNsqInput) (out *cNsqOutput, err error) {

	if err := s.Init(ctx); nil != err {
		panic("service init failed")
	}
	allConsumers := c.getAllConsumers(ctx)
	for _, cfg := range allConsumers {
		work := nsqc.NewNsqWorker(cfg.Consumer.GetTopic(), cfg.Channel, cfg.Consumer.NsqConsumerMember)
		err := work.ConnectWithConcurrency(ctx, cfg.Num)
		if err != nil {
			panic(fmt.Sprintf("NewNsqWorker.Connect failed, topic %s, err=%v", cfg.Consumer.GetTopic(), err))
		}
		s.List = append(s.List, work)
	}

	if err := s.Start(); nil != err {
		panic("service heartbeat start failed")
	}

	time.Sleep(time.Second * 3)
	for _, work := range s.List {
		work.Stop()
	}

	for _, cfg := range allConsumers {
		cfg.Consumer.Wait()
	}
	return
}

func (c cNsq) getAllConsumers(ctx context.Context) []ConsumerCfg {
	return []ConsumerCfg{
		{
			Channel:  "default",
			Num:      2,
			Consumer: nsqcmd.NewNsqDemo(ctx),
		},
	}
}
