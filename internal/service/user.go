// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"context"
	"middle/internal/model"
	"middle/internal/model/entity"
)

type (
	IUser interface {
		GetUserById(ctx context.Context, id uint) (*entity.CmsUser, error)
		SaveAndReturnInfoByFeishu(ctx context.Context, feishuUser *model.FeiShuUserInfoDetail) (*entity.CmsUser, error)
	}
)

var (
	localUser IUser
)

func User() IUser {
	if localUser == nil {
		panic("implement not found for interface IUser, forgot srv_register?")
	}
	return localUser
}

func RegisterUser(i IUser) {
	localUser = i
}
