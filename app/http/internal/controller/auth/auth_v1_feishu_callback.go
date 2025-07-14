package auth

import (
	"context"
	"middle/internal/service"

	"middle/app/http/api/auth/v1"
)

func (c *ControllerV1) FeishuCallback(ctx context.Context, req *v1.FeishuCallbackReq) (res *v1.FeishuCallbackRes, err error) {
	service.Auth().FeishuCallback(ctx, req)
	return nil, nil
}
