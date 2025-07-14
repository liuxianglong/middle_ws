package auth

import (
	"context"
	"middle/internal/service"

	"middle/app/http/api/auth/v1"
)

func (c *ControllerV1) FeishuJump(ctx context.Context, req *v1.FeishuJumpReq) (res *v1.FeishuJumpRes, err error) {
	jumpUrlData, err := service.Auth().JumpFeiShuUrl(ctx, req)
	if err != nil {
		return nil, err
	}
	res = &v1.FeishuJumpRes{
		jumpUrlData,
	}
	return res, nil
}
