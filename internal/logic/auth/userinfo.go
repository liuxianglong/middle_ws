package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "middle/app/http/api/auth/v1"
	"middle/internal/service"
	"middle/protobuf/pb"
	"middle/utility/code"
)

func (s *sAuth) Userinfo(ctx context.Context, req *v1.UserinfoReq) (*pb.AuthSSOUserinfo, error) {
	authHeader := g.RequestFromCtx(ctx).Header.Get("Authorization")
	if len(authHeader) <= 7 {
		return nil, code.CodeError.New(ctx, code.AuthMissingAuthorization)
	}

	// 验证JWT令牌
	// 去掉 "Bearer "
	secretKey, err := s.getSSOSecretKey(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.Userinfo getSSOSecretKey err=%v", err)
		return nil, code.CodeError.New(ctx, code.AuthSSOSecretKeyNoFound)
	}
	claims, err := service.Jwt().DecryptJWE(authHeader[7:], secretKey)

	if err != nil {
		return nil, code.CodeError.New(ctx, code.AuthInvalidToken)
	}

	user, err := service.User().GetUserById(ctx, claims.Uid)
	if err != nil {
		g.Log().Errorf(ctx, "sAuth.Userinfo GetUserById err=%v", err)
		return nil, code.CodeError.New(ctx, code.SystemServerErr)
	}
	if user == nil {
		return nil, code.CodeError.New(ctx, code.AuthUserNoFound)
	}
	return &pb.AuthSSOUserinfo{
		Uid:  int32(user.Id),
		Name: user.Name,
	}, nil
}
