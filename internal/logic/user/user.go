package user

import (
	"context"
	"middle/internal/dao"
	"middle/internal/model/entity"
	"middle/internal/service"
)

type (
	sUser struct{}
)

func init() {
	service.RegisterUser(New())
}

func New() service.IUser {
	return &sUser{}
}

func (s *sUser) GetUserById(ctx context.Context, id uint) (*entity.CmsUser, error) {
	cmsUser, err := dao.CmsUser.GetRecordById(ctx, id)
	if err != nil {
		return nil, err
	}
	return cmsUser, nil
}
