package consts

import "github.com/golang-module/carbon/v2"

const (
	ContextKey      = "DemoContextKey"
	DefaultI18n     = "ja"
	AppId           = 55
	AppName         = "yoipapa"
	RunModeDev      = "dev"
	DefaultTimeZone = carbon.Japan //如果需要用户时区，这个需改成日本
	UnlimitedTimes  = 999999       //无限次数
	DefaultName     = "无名"
)

const (
	LangZh = "zh"
	LangEn = "en"
	LangJa = "ja"
	LangKo = "ko"
	LangVi = "vi"
)

var AllowLangList = []string{
	LangZh,
	LangEn,
	LangJa,
	LangKo,
	LangVi,
}
