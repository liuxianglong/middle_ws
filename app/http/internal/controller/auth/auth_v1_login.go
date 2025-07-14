package auth

import (
	"context"
	"middle/internal/service"

	"middle/app/http/api/auth/v1"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	//如果已登录则直接跳转
	service.Auth().Login(ctx, req)
	return nil, nil
}
