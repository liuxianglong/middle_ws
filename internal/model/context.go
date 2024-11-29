package model

import (
	"github.com/gogf/gf/v2/frame/g"
)

const (
	// TrackKey 上下文传递
	TrackKey = "TraceId"
)

type Context struct {
	User *ContextUser // User in context.
	Data g.Map
	I18n *I18n
}

type ContextUser struct {
	UID int64
	//AppID    AppID //用户对应的APPID
	Salt     string
	Platform string //用户平台
	Time     uint32 //令牌生成时间
	Agent    string //用户的Agent
	Channel  string //用户所属渠道
}

//func (ctx *ContextUser) GetLanguage() string {
//	lang := strings.ToLower(ctx.Language)
//	if _, ok := languages[lang]; ok {
//		return lang
//	}
//	return defaultLang
//}

// GetChannel 返回当前用户所属渠道，来自用户登录时产生的令牌里
func (ctx *ContextUser) GetChannel() string {
	return ctx.Channel
}

// GetAgent 返回用户所属平台，ios|android|pc
func (ctx *ContextUser) GetAgent() string {
	return ctx.Agent
}

// IsLogined 判断当前用户是否登录
func (ctx *ContextUser) IsLogined() bool {
	return ctx.UID > 0 && ctx.Time > 0
}
