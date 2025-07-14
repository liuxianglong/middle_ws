// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package auth

import (
	"context"

	"middle/app/http/api/auth/v1"
)

type IAuthV1 interface {
	FeishuCallback(ctx context.Context, req *v1.FeishuCallbackReq) (res *v1.FeishuCallbackRes, err error)
	FeishuJump(ctx context.Context, req *v1.FeishuJumpReq) (res *v1.FeishuJumpRes, err error)
	Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error)
	Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error)
	Token(ctx context.Context, req *v1.TokenReq) (res *v1.TokenRes, err error)
	Userinfo(ctx context.Context, req *v1.UserinfoReq) (res *v1.UserinfoRes, err error)
}
