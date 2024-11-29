package crontab

import (
	"context"
	"fmt"
)

var (
	CrontabDemo = cDemo{}
)

type cDemo struct {
}

func (r cDemo) HandleDemoData(ctx context.Context) {
	fmt.Println("crontab demo HandleDemoData start")

	fmt.Println("crontab demo HandleDemoData done")
	return
}
