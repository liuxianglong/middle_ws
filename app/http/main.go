package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "github.com/gogf/gf/contrib/nosql/redis/v2"

	"github.com/gogf/gf/v2/os/gctx"

	"demogogo/app/http/internal/cmd"
	_ "demogogo/internal/boot"
	_ "demogogo/internal/logic"
	_ "demogogo/packed"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
