package auth

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"middle/internal/consts"
	"middle/internal/model"
	"middle/internal/model/entity"
	"middle/internal/service"
	"middle/utility"
	"net/url"
)

type (
	sAuth struct {
	}
)

func init() {
	service.RegisterAuth(New())
}

func New() service.IAuth {
	return &sAuth{}
}

// callbackSucAfter 回调成功后的操作
func (s *sAuth) setSession(ctx context.Context, user *entity.CmsUser) (err error) {
	//把信息存cookie
	sessionUser := &model.SessionUser{}
	err = utility.CopyFields(user, sessionUser)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.setSession CopyFields failed, err:%v", err)
		return err
	}

	err = g.RequestFromCtx(ctx).Session.Set(consts.SessionKey, sessionUser)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.setSession SetSession failed, err:%v", err)
		return err
	}
	//同时
	sessionId, _ := g.RequestFromCtx(ctx).Session.Id()

	authUserSessioncacheKey := fmt.Sprintf(consts.AuthUserSession, sessionUser.Id)
	//过期时间需要比cooike时间要长
	err = g.Redis().SetEX(ctx, authUserSessioncacheKey, sessionId, consts.SessionExpireTime+300)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.setSession set AuthUserSession err,err=%v", err)
	}

	return nil
}

// loginSucJumpCallback 登录成功后跳转
func (s *sAuth) loginSucJumpCallback(ctx context.Context, ssoAuth *model.SSOAuth, uid uint) {
	frontUrlVar, _ := g.Cfg().Get(ctx, "server.frontUrl")
	jumpUrl := frontUrlVar.String()
	appID := ssoAuth.ClientId
	if appID == "" {
		//跳到前端页
		//@todo 因为没有前端页，先全跳到/去
		if ssoAuth.RedirectUri != "" {
			jumpUrl = ssoAuth.RedirectUri
		}
		g.RequestFromCtx(ctx).Response.RedirectTo(jumpUrl, 302)
		return
	} else {
		//去app表获取回调地址
		ssoManager, err := service.SsoManager().GetInfoByAppID(ctx, appID)
		if err != nil {
			g.Log().Errorf(ctx, "sAuth.loginSucJumpCallback GetInfoByAppID failed, err:%v", err)
			g.RequestFromCtx(ctx).Response.RedirectTo(jumpUrl, 302)
			return
		}

		//如果查不到则跳到自己的前端页
		if ssoManager == nil {
			g.RequestFromCtx(ctx).Response.RedirectTo(jumpUrl, 302)
			return
		}

		//如果查到则带着code跳回
		ssoCode := s.generateAndStoreCode(ctx, uid, ssoManager)
		if ssoCode != "" {
			g.RequestFromCtx(ctx).Response.RedirectTo(jumpUrl, 302)
			return
		}
		callbackUrl := ssoManager.CallbackUrl
		redirectURL, _ := url.Parse(callbackUrl)
		query := redirectURL.Query()
		query.Set("state", ssoAuth.State)
		query.Set("redirect_uri", ssoAuth.RedirectUri)
		query.Set("code", ssoCode)
		redirectURL.RawQuery = query.Encode()
		g.RequestFromCtx(ctx).Response.RedirectTo(redirectURL.String(), 302)

	}
	return
}

func (s *sAuth) generateAndStoreCode(ctx context.Context, uid uint, ssoManager *entity.CmsAuthSso) string {
	maxRetries := 10
	ssoCodeStruct := &model.SSOCode{
		ClientId:     ssoManager.AppId,
		ClientSecret: ssoManager.AppSecret,
		Uid:          uid,
	}
	for i := 0; i < maxRetries; i++ {
		ssoCode := utility.RandomString(16)
		cacheLockKey := s.generateCodeCacheKey(ssoCode)
		if utility.RedisLock(ctx, g.Redis(), cacheLockKey, ssoCodeStruct, 300) {
			return ssoCode
		}
	}
	g.Log().Warningf(ctx, "sAuth.generateAndStoreCode failed, retry times: %d", maxRetries)
	return ""
}

func (s *sAuth) generateCodeCacheKey(ssoCode string) string {
	return fmt.Sprintf(consts.AuthCodeLockCache, ssoCode)
}

func (s *sAuth) getSSOSecretKey(ctx context.Context) (string, error) {
	ssoSecretVar, err := g.Cfg().Get(ctx, "server.ssoSecret")
	if err != nil {
		return "", err
	}
	return ssoSecretVar.String(), nil
}
