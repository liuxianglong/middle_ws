package cmd

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

type Command struct {
	*gcmd.Command
}

func (c *Command) Run(ctx context.Context) {
	defer func() {
		if exception := recover(); exception != nil {
			if err, ok := exception.(error); ok {
				g.Log().Print(ctx, err.Error())
			} else {
				panic(gerror.NewCodef(gcode.CodeInternalPanic, "%+v", exception))
			}
		}
	}()

	// 原生的直接中断进程了，在这修改为error报警
	if err := c.RunWithError(ctx); err != nil {
		g.Log().Errorf(ctx, err.Error())
	}
}

func GetCommand(ctx context.Context) (*Command, error) {
	//service.Cache().RedisRegister(ctx)
	initCommand, err := gcmd.NewFromObject(Main)
	if err != nil {
		panic(err)
	}
	// 所有消费进程在这注册
	err = initCommand.AddObject(
		Demo,
		Crontab,
	)
	if err != nil {
		g.Log().Errorf(ctx, "process error,err=%v", err)
		panic(err)
	}
	command := &Command{
		Command: initCommand,
	}
	return command, nil
}
