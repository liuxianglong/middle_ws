package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type FeishuCallbackReq struct {
	g.Meta `path:"/auth/feishu/callback" tags:"auth" method:"get" summary:"飞书回调"`
	Code   string `dc:"code" json:"code"`
	State  string `dc:"唯一标识符" json:"state"`
	Error  string `dc:"错误值" json:"error"`
}

type FeishuCallbackRes struct {
	//*pb.LoginJumpUrlData
	//g.Meta `mime:"text/html" example:"string"`
}
