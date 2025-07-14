package model

type JWTClaims struct {
	Uid uint  `json:"uid"`
	Exp int64 `json:"exp"`
}

type JWEHeader struct {
	Alg string `json:"alg"` // 密钥派生算法
	Enc string `json:"enc"` // 内容加密算法
	Kid string `json:"kid"` // 密钥ID（可选）
}
