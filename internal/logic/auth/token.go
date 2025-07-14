package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "middle/app/http/api/auth/v1"
	"middle/internal/model"
	"middle/internal/service"
	"middle/protobuf/pb"
	"middle/utility/code"
	"time"
)

func (s *sAuth) Token(ctx context.Context, req *v1.TokenReq) (*pb.AuthSSOTokenData, error) {
	cacheKey := s.generateCodeCacheKey(req.Code)

	ssoCode, err := s.getSSOCodeCache(ctx, cacheKey)
	if err != nil {
		return nil, err
	}

	err = s.verifyTokenReq(ctx, req, ssoCode)
	if err != nil {
		return nil, err
	}
	//不要放要验证前，否则会将正常的key给删除
	defer func(ctx context.Context, key string) {
		_, err := g.Redis().Del(ctx, key)
		if err != nil {
			g.Log().Errorf(ctx, "sAuth.Token delCache failed,err=%v", err)
			return
		}
	}(ctx, cacheKey)
	//ssoCode := &model.SSOCode{
	//	ClientSecret: "11",
	//	ClientId:     "44",
	//	Uid:          1,
	//}

	//有一点要注意，返回code时，有可能cookie过期，这是否还要正常返回，还是说cookie过期则报错

	//生成token体
	current := time.Now().Unix()
	validTime := 3600
	expireTime := current + int64(validTime)
	jwtClaims := &model.JWTClaims{
		Uid: ssoCode.Uid,
		Exp: expireTime,
	}
	secretKey, err := s.getSSOSecretKey(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.Token getSSOSecretKey err=%v", err)
		return nil, code.CodeError.New(ctx, code.AuthSSOSecretKeyNoFound)
	}

	token, err := service.Jwt().GenerateJWE(jwtClaims, secretKey)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.Token GenerateJWE err=%v", err)
		return nil, code.CodeError.New(ctx, code.AuthBuildTokenError)
	}

	return &pb.AuthSSOTokenData{
		AccessToken: token,
		TokenType:   "Bearer",
		ExpiresIn:   int32(validTime),
		//Expiry:
	}, nil
}

func (s *sAuth) getSSOCodeCache(ctx context.Context, cacheKey string) (ssoCode *model.SSOCode, err error) {
	ssoCodeVar, err := g.Redis().Get(ctx, cacheKey)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.Token Redis.get err=%v", err)
		return nil, code.CodeError.New(ctx, code.SystemServerErr)
	}
	ssoCode = &model.SSOCode{}

	err = ssoCodeVar.Struct(ssoCode)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.Token ssoCodeVar.Struct err=%v", err)
		return nil, code.CodeError.New(ctx, code.SystemServerErr)
	}
	return ssoCode, nil
}
func (s *sAuth) verifyTokenReq(ctx context.Context, req *v1.TokenReq, ssoCode *model.SSOCode) error {
	if req.GrantType != "authorization_code" {
		return code.CodeError.New(ctx, code.AuthGrantTypeError)
	}

	if ssoCode.ClientId != req.ClientId || ssoCode.ClientSecret != req.ClientSecret {
		return code.CodeError.New(ctx, code.AuthGrantSecretError)
	}
	return nil
}
