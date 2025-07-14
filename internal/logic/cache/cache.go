package cache

import (
	"context"
	"github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"middle/internal/model"
	"middle/internal/service"
)

type (
	sCache struct{}
)

func init() {
	service.RegisterCache(New())
}

func New() service.ICache {
	return &sCache{}
}

func (s *sCache) RedisRegister(ctx context.Context) {
	//获取redis的配置
	redisVars, err := g.Config().Get(ctx, "redis")
	if err != nil {
		panic("redis no config")
	}
	debug := false
	debugVar, _ := g.Config().Get(ctx, "redis-debug")
	if debugVar.Bool() {
		debug = true
	}

	cfgs := map[string]*model.RedisAdapterConfig{}
	err = redisVars.Struct(&cfgs)
	if err != nil {
		g.Log().Errorf(ctx, "redis parse config err,err=%v", err)
		panic("redis parse config err")
	}
	gredis.RegisterAdapterFunc(func(config *gredis.Config) gredis.Adapter {
		r := &model.RedisAdapter{
			Redis: redis.New(config),
			Debug: debug,
		}
		r.AdapterOperation = r // This is necessary.
		return r
	})
	for groupName, cfg := range cfgs {
		gredis.SetConfig(cfg.Config, groupName)
	}
}
