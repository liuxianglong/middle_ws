package nsqcmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/nsqio/go-nsq"
	"middle/internal/consts"
	"sync"
)

func NewNsqDemo(ctx context.Context) *demo {
	return &demo{
		ctx: ctx,
	}
}

type demo struct {
	syncWait sync.WaitGroup
	ctx      context.Context
}

func (mr *demo) NsqConsumerMember(msg *nsq.Message) error {
	mr.syncWait.Add(1)
	defer mr.syncWait.Done()
	fmt.Println(string(msg.Body))

	return nil
}

func (mr *demo) Wait() {
	g.Log().Info(mr.ctx, "demo.Wait in")
	mr.syncWait.Wait()
	g.Log().Info(mr.ctx, "demo.Wait out")
}

func (mr *demo) GetTopic() string {
	return consts.TopicDemo
}
