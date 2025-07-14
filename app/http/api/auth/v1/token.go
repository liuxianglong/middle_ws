package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"middle/protobuf/pb"
)

type TokenReq struct {
	g.Meta       `path:"/auth/token" tags:"auth" method:"post" summary:"获取token"`
	Code         string `v:"required" dc:"授权码" json:"code"`
	ClientId     string `v:"required" dc:"系统给的clientId" json:"client_id"`
	ClientSecret string `v:"required" dc:"系统给的clientSecret" json:"client_secret"`
	GrantType    string `v:"required" dc:"授权类型，固定值authorization_code" json:"grant_type"`
}

type TokenRes struct {
	*pb.AuthSSOTokenData
}
