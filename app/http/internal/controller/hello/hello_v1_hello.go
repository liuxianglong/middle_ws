package hello

import (
	"context"

	"middle/app/http/api/hello/v1"
)

func (c *ControllerV1) Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error) {
	return nil, nil
}
