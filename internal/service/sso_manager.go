// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"middle/internal/model/entity"
)

type (
	ISsoManager interface {
		GetInfoByAppID(ctx context.Context, appID string) (*entity.CmsAuthSso, error)
	}
)

var (
	localSsoManager ISsoManager
)

func SsoManager() ISsoManager {
	if localSsoManager == nil {
		panic("implement not found for interface ISsoManager, forgot srv_register?")
	}
	return localSsoManager
}

func RegisterSsoManager(i ISsoManager) {
	localSsoManager = i
}
