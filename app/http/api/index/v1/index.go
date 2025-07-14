package v1

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `path:"/" tags:"index" method:"get" summary:"首页"`
}

type IndexRes struct {
}
