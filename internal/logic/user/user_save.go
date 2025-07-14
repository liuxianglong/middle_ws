package user

import (
	"context"
	"middle/internal/dao"
	"middle/internal/model"
	"middle/internal/model/entity"
	"time"
)

func (s *sUser) SaveAndReturnInfoByFeishu(ctx context.Context, feishuUser *model.FeiShuUserInfoDetail) (*entity.CmsUser, error) {
	//1.通过code查询用户是否存在
	cmsUser, err := dao.CmsUser.GetRecordByFeishuCode(ctx, feishuUser.UserId)
	if err != nil {
		return nil, err
	}
	//2.如果存在则直接返回
	if cmsUser != nil {
		return cmsUser, nil
	}

	//3.如果不存在则新增
	currentTime := time.Now().Unix()
	cmsUser = &entity.CmsUser{
		Name:       feishuUser.Name,
		Email:      feishuUser.Email,
		FeishuCode: feishuUser.UserId,
		CreateAt:   currentTime,
		UpdateAt:   currentTime,
	}
	lastId, err := dao.CmsUser.SaveRecord(ctx, cmsUser)
	if err != nil {
		return nil, err
	}
	cmsUser.Id = uint(lastId)
	return cmsUser, nil
}
