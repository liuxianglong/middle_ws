package middleware

import (
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"middle/protobuf/pb"
	"net/http"
)

// SsoTokenResponse
// sso——token返回值需特殊处理，目的：为符合oauth2.0标准
func (s *sMiddleware) SsoTokenResponse(r *ghttp.Request) {
	r.Middleware.Next()

	var (
		msg     string
		err     = r.GetError()
		res     = r.GetHandlerResponse()
		codeErr = gerror.Code(err)
	)
	if err != nil {
		if codeErr == gcode.CodeNil {
			codeErr = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			msg = http.StatusText(r.Response.Status)
			switch r.Response.Status {
			case http.StatusNotFound:
				codeErr = gcode.CodeNotFound
			case http.StatusForbidden:
				codeErr = gcode.CodeNotAuthorized
			default:
				codeErr = gcode.CodeUnknown
			}
			// It creates error as it can be retrieved by other middlewares.
			err = gerror.NewCode(codeErr, msg)
			r.SetError(err)
		} else {
			//codeErr = gcode.CodeOK
			//msg = ""
			r.Response.WriteJson(res)
			return
		}
	}

	r.Response.Status = http.StatusBadRequest
	r.Response.WriteJson(pb.AuthSSOTokenFail{
		Code:  int32(codeErr.Code()),
		Error: msg,
	})
	r.Response.ClearBuffer()
	r.ExitAll()
}
