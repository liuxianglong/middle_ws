// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"middle/internal/model"
)

type (
	IJwt interface {
		GenerateJWE(claims *model.JWTClaims, secret string) (string, error)
		// DecryptJWE 解密并验证 JWT
		DecryptJWE(jwe string, secret string) (*model.JWTClaims, error)
	}
)

var (
	localJwt IJwt
)

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot srv_register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}
