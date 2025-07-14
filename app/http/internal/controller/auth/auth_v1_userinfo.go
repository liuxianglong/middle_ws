package auth

import (
	"context"
	"middle/internal/service"

	"middle/app/http/api/auth/v1"
)

func (c *ControllerV1) Userinfo(ctx context.Context, req *v1.UserinfoReq) (res *v1.UserinfoRes, err error) {
	userinfo, err := service.Auth().Userinfo(ctx, req)
	if err != nil {
		return nil, err
	}
	if userinfo != nil {
		res = &v1.UserinfoRes{
			AuthSSOUserinfo: userinfo,
		}
	}
	return res, nil
}
