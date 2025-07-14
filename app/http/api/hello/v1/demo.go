package v1

import "github.com/gogf/gf/v2/frame/g"

type DemoReq struct {
	g.Meta `path:"/demo" tags:"demo" method:"get" summary:"You first hello api"`
}

type DemoRes struct {
	g.Meta `mime:"text/html" example:"string"`
}
