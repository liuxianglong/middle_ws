package auth

import (
	"context"
	"middle/internal/service"

	"middle/app/http/api/auth/v1"
)

func (c *ControllerV1) Token(ctx context.Context, req *v1.TokenReq) (res *v1.TokenRes, err error) {
	token, err := service.Auth().Token(ctx, req)
	if err != nil {
		return nil, err
	}
	return &v1.TokenRes{
		AuthSSOTokenData: token,
	}, nil
}
