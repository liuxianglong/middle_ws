package model

import (
	"context"
	"github.com/gogf/gf/contrib/nosql/redis/v2"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

type RedisAdapterConfig struct {
	*gredis.Config
}

type RedisAdapter struct {
	Debug bool `json:"debug"`
	*redis.Redis
}

// Do implements and overwrites the underlying function Do from Adapter.
func (r *RedisAdapter) Do(ctx context.Context, command string, args ...interface{}) (*gvar.Var, error) {
	start := time.Now()
	defer func() {
		if r.Debug {
			end := time.Now()
			g.Log().Debugf(ctx, "[ %d ms] [redis] command=%s, args=%v", end.Sub(start).Milliseconds(), command, args)
		}
	}()

	return r.Redis.Do(ctx, command, args...)
}
