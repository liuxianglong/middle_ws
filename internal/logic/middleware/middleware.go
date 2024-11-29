package middleware

import (
	"demogogo/internal/model"
	"demogogo/internal/service"
	"demogogo/utility"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"net/http"

	"demogogo/internal/consts"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
)

type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

func (s *sMiddleware) CORS(r *ghttp.Request) {
	// 开发和测试允许跨域
	corsOptions := r.Response.DefaultCORSOptions()
	runMode, _ := g.Cfg().Get(r.Context(), "server.runMode")
	if runMode.String() != consts.RunModeDev {
		corsOptions.AllowDomain = []string{
			"demo.com",
		}
	}
	corsOptions.AllowOrigin = "*"
	corsOptions.AllowMethods = "GET,POST"
	corsOptions.AllowHeaders = "x-requested-with,content-type,user-token,user-language,lang"
	corsOptions.ExposeHeaders = "date,user-status"
	corsOptions.AllowCredentials = "true"

	r.Response.CORS(corsOptions)
	r.Middleware.Next()
}

func (s *sMiddleware) HandleResponse(r *ghttp.Request) {
	// 执行下一步请求逻辑
	r.Middleware.Next()

	if r.Response.BufferLength() > 0 {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			g.Log().Errorf(r.GetCtx(), "system error, err=%s,code=%v", string(r.Response.Buffer()), r.Response.Status)
			r.Response.ClearBuffer()
			r.Response.Write("System error, please try again")
		}

		return
	}

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
			codeErr = gcode.CodeOK
			msg = "success"
		}
	}

	r.Response.WriteJson(ghttp.DefaultHandlerResponse{
		Code:    codeErr.Code(),
		Message: msg,
		Data:    res,
	})
}

func (s *sMiddleware) Ctx(r *ghttp.Request) {
	customCtx := &model.Context{
		I18n: model.NewI18n(),
		Data: make(g.Map),
		User: &model.ContextUser{},
	}
	lang := r.GetHeader("Lang")
	if !(utility.InArray(lang, consts.AllowLangList)) {
		lang = consts.DefaultI18n
	}

	customCtx.I18n.SetLanguage(lang)
	service.BizCtx().Init(r, customCtx)
	r.SetCtx(gi18n.WithLanguage(r.Context(), lang))
	r.Middleware.Next()
}
