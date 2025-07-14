package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type LogoutReq struct {
	g.Meta `path:"/auth/logout" tags:"auth" method:"get" summary:"登出"`
}

type LogoutRes struct {
}
