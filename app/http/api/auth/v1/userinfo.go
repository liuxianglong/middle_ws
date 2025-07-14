package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"middle/protobuf/pb"
)

type UserinfoReq struct {
	g.Meta `path:"/auth/userinfo" tags:"auth" method:"get" summary:"通过token获取用户信息"`
}

type UserinfoRes struct {
	*pb.AuthSSOUserinfo
}
