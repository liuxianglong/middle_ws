package aliyun

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"testing"
)

func Test_library_aliyun_view(t *testing.T) {
	ctx := context.Background()
	ossCfg, err := g.Cfg().Get(ctx, "oss")
	if err != nil {
		panic("no oss config")
	}
	ossLocalConfig = &KeyLocalConfig{}
	ossCfg.Struct(ossLocalConfig)
	g.DumpJson(ossLocalConfig)
}
