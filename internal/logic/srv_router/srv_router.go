package srv_router

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"middle/internal/model"
	"middle/internal/service"
	"sync"
	"time"
)

type sSrvRouter struct {
	routerMap sync.Map
	version   string
}

func init() {
	service.RegisterSrvRouter(New())
}

func New() service.ISrvRouter {
	return &sSrvRouter{}
}

func (s *sSrvRouter) getCfg(ctx context.Context) (*model.SrvRouterCfg, error) {
	cc, _ := g.Config().Get(ctx, "srv-router")
	srvRouterCfg := &model.SrvRouterCfg{}
	err := cc.Struct(srvRouterCfg)
	if err != nil {
		return nil, err
	}
	return srvRouterCfg, nil
}

func (s *sSrvRouter) InitRouter(ctx context.Context) {
	srvRouterCfg, err := s.getCfg(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "sLimiter.Init 失败， err=%v", err)
		return
	}
	s.version = srvRouterCfg.Version

	//通过consul拿到对应配置
	s.routerMap = sync.Map{}
	for routerName, v := range srvRouterCfg.Route {
		s.routerMap.Store(routerName, v)
	}

	return
}

func (s *sSrvRouter) GetRouter(ctx context.Context, routerName string) *model.SrvRouterServer {
	v, ok := s.routerMap.Load(routerName)
	if ok {
		return v.(*model.SrvRouterServer)
		//v.(*rate.Limiter)
	}
	return nil
}

// Lookup 监听,更新配置
func (s *sSrvRouter) Lookup(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			srvRouterCfg, err := s.getCfg(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "sLimiter.Lookup 失败， err=%v", err)
				return
			}
			if srvRouterCfg.Version != s.version {
				//1.查询是否对部分路由取消
				var toDelete []interface{}
				s.routerMap.Range(func(key, value interface{}) bool {
					if k, ok := key.(string); ok {
						if _, exists := srvRouterCfg.Route[k]; !exists {
							toDelete = append(toDelete, key)
						}
					}
					return true
				})
				//2. 执行删除
				for _, key := range toDelete {
					s.routerMap.Delete(key)
				}

				for routerName, v := range srvRouterCfg.Route {
					//和现有的做对比，如果没变化则不动
					_, ok := s.routerMap.Load(routerName)
					if ok {
						continue
					}

					s.routerMap.Store(routerName, v)
				}
			}
		}
	}
}
