// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package hello

import (
	"context"

	"middle/app/http/api/hello/v1"
)

type IHelloV1 interface {
	Demo(ctx context.Context, req *v1.DemoReq) (res *v1.DemoRes, err error)
	Hello(ctx context.Context, req *v1.HelloReq) (res *v1.HelloRes, err error)
}
