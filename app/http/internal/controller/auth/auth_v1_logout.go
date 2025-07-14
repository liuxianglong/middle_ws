package auth

import (
	"context"
	"middle/internal/service"

	"middle/app/http/api/auth/v1"
)

func (c *ControllerV1) Logout(ctx context.Context, req *v1.LogoutReq) (res *v1.LogoutRes, err error) {
	service.Auth().Logout(ctx)
	return nil, nil
}
