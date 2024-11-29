package cmd

import (
	"context"
	"demogogo/app/job/internal/cmd/crontab"
	"demogogo/library/cmd"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
)

// 这是框架自带的crontab，适情况而用

var (
	Crontab  = cCrontab{}
	crontabS = &crontabServer{}
)

type crontabServer struct {
	cmd.StartUp
	//List []*nsqc.NsqWorker //所有的消费实例
}

type cCrontab struct {
	g.Meta `name:"crontab" brief:"this is a crontab"`
}

type cCrontabInput struct {
	g.Meta `name:"crontab"`
	//Action string `v:"required" short:"a" name:"action" brief:"action"`
}

type cCrontabOutput struct{}

func (c cCrontab) Index(ctx context.Context, in cCrontabInput) (out *cCrontabOutput, err error) {
	g.Log().Printf(ctx, "crontab start")
	if err := crontabS.Init(ctx); nil != err {
		panic("service init failed")
	}

	_, err = gcron.AddSingleton(ctx, "# 0 */1 * * *", crontab.CrontabDemo.HandleDemoData, "CrontabDemo_HandleDemoData")
	if err != nil {
		g.Log().Errorf(ctx, "%v", err)
		panic(err)
	}

	if err := crontabS.Start(); nil != err {
		panic("service heartbeat start failed")
	}
	g.Log().Printf(ctx, "crontab stop")
	return
}
