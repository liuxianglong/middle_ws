package tool

import (
	"strconv"
	"strings"
)

// Str 导出str对象
var Str = &str{}

type str struct{}

// UnescapeUnicode Unicode 变成字符串
func (*str) UnescapeUnicode(raw []byte) ([]byte, error) {
	str, err := strconv.Unquote(strings.Replace(strconv.Quote(string(raw)), `\\u`, `\u`, -1))
	if err != nil {
		return nil, err
	}
	return []byte(str), nil
}

// EscapeUnicode 字符串变成\u Unicode
func (*str) EscapeUnicode(raw string) string {
	textQuoted := strconv.QuoteToASCII(raw)
	return textQuoted[1 : len(textQuoted)-1]
}
