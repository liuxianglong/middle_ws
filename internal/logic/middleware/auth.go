package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

func (s *sMiddleware) Auth(r *ghttp.Request) {
	r.Middleware.Next()
}
