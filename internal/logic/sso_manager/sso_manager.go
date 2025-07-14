package sso_manager

import (
	"context"
	"middle/internal/dao"
	"middle/internal/model/entity"
	"middle/internal/service"
)

type (
	sSsoManager struct{}
)

func init() {
	service.RegisterSsoManager(New())
}

func New() service.ISsoManager {
	return &sSsoManager{}
}

func (s *sSsoManager) GetInfoByAppID(ctx context.Context, appID string) (*entity.CmsAuthSso, error) {
	//@todo 所有相关做缓存管理 我们提前将所有的appID存储
	return dao.CmsAuthSso.GetRecordByAppID(ctx, appID)
}
