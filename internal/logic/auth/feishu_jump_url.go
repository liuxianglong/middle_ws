package auth

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/oauth2"
	v1 "middle/app/http/api/auth/v1"
	"middle/internal/consts"
	"middle/internal/model"
	"middle/internal/service"
	"middle/protobuf/pb"
	"middle/utility"
	"middle/utility/code"
)

// JumpFeiShuUrl 登录并返回跳转的url
func (s *sAuth) JumpFeiShuUrl(ctx context.Context, req *v1.FeishuJumpReq) (*pb.LoginJumpUrlData, error) {
	if service.BizCtx().Get(ctx).User.IsLogin() { //如果登录了则不再生成url
		return nil, code.CodeError.New(ctx, code.AuthHasLogin)
	}
	oauthConfig, err := s.oauthFeishuCommonCfg(ctx)
	if err != nil {
		return nil, code.CodeError.New(ctx, code.SystemServerErr)
	}
	//把appID进行临时存储
	state, err := s.handleState(ctx, req.SSOAuth, 16)
	if err != nil {
		return nil, code.CodeError.New(ctx, code.SystemServerErr)
	}
	jumpUrl := oauthConfig.AuthCodeURL(state)

	//@todo 这块让前端跳，测试时用这个
	g.RequestFromCtx(ctx).Response.RedirectTo(jumpUrl, 302)
	return &pb.LoginJumpUrlData{
		JumpUrl: jumpUrl,
	}, nil
}

func (s *sAuth) buildStateCacheKey(ctx context.Context, state string) string {
	return fmt.Sprintf(consts.AuthStateCache, state)
}
func (s *sAuth) handleState(ctx context.Context, ssoAuth *model.SSOAuth, length int) (string, error) {
	state := utility.RandomString(length)
	cacheKey := s.buildStateCacheKey(ctx, state)

	err := g.Redis().SetEX(ctx, cacheKey, ssoAuth, 300)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.handleState err,err=%v", err)
		return "", err
	}
	return state, nil
}

func (s *sAuth) oauthFeishuCommonCfg(ctx context.Context) (*oauth2.Config, error) {
	var oauthEndpoint = oauth2.Endpoint{
		AuthURL:  consts.FeishuAuthURL,
		TokenURL: consts.FeishuTokenURL,
	}

	feishuVar, _ := g.Cfg().Get(ctx, "feishu")
	feishucfg := &model.FeiShuCfg{}
	err := feishuVar.Struct(feishucfg)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.oauthFeishuCommonCfg feishuVar.Struct err=%v", err)
		return nil, err
	}

	return &oauth2.Config{
		ClientID:     feishucfg.ClientID,
		ClientSecret: feishucfg.ClientSecret,
		RedirectURL:  feishucfg.RedirectURL, // 请先添加该重定向 URL，配置路径：开发者后台 -> 开发配置 -> 安全设置 -> 重定向 URL -> 添加
		Endpoint:     oauthEndpoint,
		//Scopes:       []string{"offline_access"}, // 如果你不需要 refresh_token，请注释掉该行，否则你需要先申请 offline_access 权限方可使用，配置路径：开发者后台 -> 开发配置 -> 权限管理
	}, nil
}
