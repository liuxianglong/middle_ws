package main

import (
	_ "demogogo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"demogogo/app/job/internal/cmd"
)

func main() {
	ctx := gctx.GetInitCtx()
	command, err := cmd.GetCommand(ctx)
	if err != nil {
		panic(err)
	}
	if command == nil {
		panic("command no found")
	}

	command.Run(ctx)
}
