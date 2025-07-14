package auth

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	v1 "middle/app/http/api/auth/v1"
	"middle/internal/consts"
	"middle/internal/model"
	"middle/internal/service"
	"middle/utility/code"
	"net/http"
)

func (s *sAuth) FeishuCallback(ctx context.Context, req *v1.FeishuCallbackReq) error {
	//1. check error 如果error不是空或code是空，则代码用户拒绝授权，固定返回access_denied
	if req.Code == "" {
		//@todo 跳转目标地址登录页
		g.Log().Warningf(ctx, "sAuth.access_denied")
		return code.CodeError.New(ctx, code.AuthAccessDenied)
	}
	//2. check state
	stateCacheKey := s.buildStateCacheKey(ctx, req.State)
	ssoAuthVar, err := g.Redis().Get(ctx, stateCacheKey)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.FeishuCallback getCache failed,err=%v", err)
		return code.CodeError.New(ctx, code.SystemServerErr)
	}
	defer func(ctx context.Context, key string) {
		_, err := g.Redis().Del(ctx, key)
		if err != nil {
			g.Log().Errorf(ctx, "sAuth.FeishuCallback delCache failed,err=%v", err)
			return
		}
	}(ctx, stateCacheKey)

	//检测appID存在不，不存在说明是恶意的
	ssoAuth := &model.SSOAuth{}
	err = ssoAuthVar.Struct(ssoAuth)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.FeishuCallback ssoAuthVar Struct failed,err=%v", err)
		return code.CodeError.New(ctx, code.SystemServerErr)
	}

	if ssoAuth == nil {
		//即redis过期或恶意请求
		return code.CodeError.New(ctx, code.AuthThirdCurlExpire)
	}

	//3. 根据code换token
	oauthConfig, err := s.oauthFeishuCommonCfg(ctx)
	if err != nil {
		return code.CodeError.New(ctx, code.SystemServerErr)
	}
	// 使用获取到的 code 获取 token
	//token, err := oauthConfig.Exchange(ctx, req.Code, oauth2.VerifierOption(codeVerifier))
	token, err := oauthConfig.Exchange(ctx, req.Code)
	if err != nil {
		g.Log().Errorf(ctx, "oauthConfig.Exchange() failed with '%v'", err)

		return code.CodeError.New(ctx, code.AuthGetThirdTokenFailed)
	}
	marshal, _ := json.Marshal(token)

	fmt.Printf("===========token = %v\n", string(marshal))
	client := oauthConfig.Client(ctx, token)
	request, err := http.NewRequest("GET", consts.FeishuUserInfoUrl, nil)
	if err != nil {
		//@todo 跳转至来源
		g.Log().Errorf(ctx, "oauthConfig.Exchange() failed with '%v'", err)
		return code.CodeError.New(ctx, code.AuthGetThirdInfoFailed)
	}
	request.Header.Set("Authorization", "Bearer "+token.AccessToken)
	// 使用 token 发起请求，获取用户信息
	resp, err := client.Do(request)
	if err != nil {
		g.Log().Errorf(ctx, "client.Do() failed with err=%v", err)

		return code.CodeError.New(ctx, code.AuthGetThirdInfoFailed)
	}
	defer resp.Body.Close()
	user := &model.FeiShuUserInfo{}
	if err := json.NewDecoder(resp.Body).Decode(&user); err != nil {
		g.Log().Errorf(ctx, "json.Decode() failed with err=%v", err)
		return code.CodeError.New(ctx, code.AuthGetThirdInfoFailed)
	}
	if user.Code != 0 {
		//报错
		g.Log().Errorf(ctx, "获取用户信息接口失败，错误码:%d,错误信息:%s", user.Code, user.Msg)
		return code.CodeError.New(ctx, code.AuthGetThirdInfoFailed)
	}
	//b, _ := json.Marshal(gate)
	//g.Log().Infof(ctx, "访问成功,信息：%s", string(b))

	//登录
	cmsUser, err := service.User().SaveAndReturnInfoByFeishu(ctx, user.Data)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.FeishuCallback SaveAndReturnInfoByFeishu failed with err=%v", err)
		return code.CodeError.New(ctx, code.SystemServerErr)
	}
	//u, _ := json.Marshal(cmsUser)
	//g.Log().Infof(ctx, "写入成功,信息：%s", string(u))

	//将拿到的信息存cookie和session
	err = s.setSession(ctx, cmsUser)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.FeishuCallback setSession failed with err=%v", err)
		return code.CodeError.New(ctx, code.AuthSaveSessionFailed)
	}
	s.loginSucJumpCallback(ctx, ssoAuth, cmsUser.Id)
	return nil
}
