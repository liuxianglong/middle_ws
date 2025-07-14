package model

import (
	"context"
	"github.com/gogf/gf/v2/i18n/gi18n"
	"middle/internal/consts"
)

type I18n struct {
	language string
	manager  *gi18n.Manager
}

func NewI18n() *I18n {
	i18n := &I18n{
		manager: gi18n.New(),
	}
	i18n.SetLanguage(consts.DefaultI18n)
	return i18n
}

func (n *I18n) SetLanguage(language string) {
	n.manager.SetLanguage(language)
	n.language = language
}

// Language 获取当前会话的语言
func (n *I18n) Language() string {
	return n.language
}

// T 根据key返回当前会话的语言
func (n *I18n) T(ctx context.Context, content string) string {
	return n.manager.Translate(ctx, content)
}

// Tf 根据key返回当前会话的语言，支持格式化
func (n *I18n) Tf(ctx context.Context, format string, values ...interface{}) string {
	return n.manager.TranslateFormat(ctx, format, values...)
}
