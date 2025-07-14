package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"middle/internal/model"
)

type LoginReq struct {
	g.Meta `path:"/auth/login" tags:"auth" method:"get" summary:"跳转登录中间页"`
	*model.SSOAuth
}

type LoginRes struct {
	//*pb.LoginJumpUrlData
}
