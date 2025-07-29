package v1

import "github.com/gogf/gf/v2/frame/g"

type IndexReq struct {
	g.Meta `path:"/ws" tags:"index" method:"all" summary:"首页"`
}

type IndexRes struct {
}
