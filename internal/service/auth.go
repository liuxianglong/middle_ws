// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	v1 "middle/app/http/api/auth/v1"
	"middle/protobuf/pb"
)

type (
	IAuth interface {
		FeishuCallback(ctx context.Context, req *v1.FeishuCallbackReq) error
		// JumpFeiShuUrl 登录并返回跳转的url
		JumpFeiShuUrl(ctx context.Context, req *v1.FeishuJumpReq) (*pb.LoginJumpUrlData, error)
		Login(ctx context.Context, req *v1.LoginReq)
		Logout(ctx context.Context)
		Token(ctx context.Context, req *v1.TokenReq) (*pb.AuthSSOTokenData, error)
		Userinfo(ctx context.Context, req *v1.UserinfoReq) (*pb.AuthSSOUserinfo, error)
	}
)

var (
	localAuth IAuth
)

func Auth() IAuth {
	if localAuth == nil {
		panic("implement not found for interface IAuth, forgot srv_register?")
	}
	return localAuth
}

func RegisterAuth(i IAuth) {
	localAuth = i
}
