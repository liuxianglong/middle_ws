package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"middle/app/http/internal/cmd"
	_ "middle/internal/boot"
	_ "middle/internal/logic"
	_ "middle/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
