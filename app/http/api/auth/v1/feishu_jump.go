package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"middle/internal/model"
	"middle/protobuf/pb"
)

type FeishuJumpReq struct {
	g.Meta `path:"/auth/feishu/jump" tags:"auth" method:"get" summary:"跳转"`
	*model.SSOAuth
}

type FeishuJumpRes struct {
	*pb.LoginJumpUrlData
	//g.Meta `mime:"text/html" example:"string"`
}
