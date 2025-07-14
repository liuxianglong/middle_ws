package hello

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"middle/app/http/api/hello/v1"
)

func (c *ControllerV1) Demo(ctx context.Context, req *v1.DemoReq) (res *v1.DemoRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
