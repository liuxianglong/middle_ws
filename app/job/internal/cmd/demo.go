package cmd

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	Demo = cDemo{}
)

type cDemo struct {
	g.Meta `name:"demo" brief:"this is a demo"`
}

type cDemoInput struct {
	g.Meta `name:"demo"`
	Action string `v:"required" short:"a" name:"action" brief:"action"`
}

type cDemoOutput struct{}

func (c cDemo) Index(ctx context.Context, in cDemoInput) (out *cDemoOutput, err error) {
	fmt.Println(111)
	return
}
