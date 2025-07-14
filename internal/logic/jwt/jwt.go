package jwt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"middle/internal/model"
	"middle/internal/service"
	"strings"
	"time"
)

type (
	sJwt struct{}
)

func init() {
	service.RegisterJwt(New())
}

func New() service.IJwt {
	return &sJwt{}
}

func (s *sJwt) GenerateJWE(claims *model.JWTClaims, secret string) (string, error) {
	// 1. 创建头部
	header := &model.JWEHeader{
		Alg: "PBKDF2",  // 密钥派生算法（简化实现）
		Enc: "A256GCM", // AES-256-GCM加密
	}

	// 序列化头部
	headerJSON, err := json.Marshal(header)
	if err != nil {
		return "", fmt.Errorf("failed to marshal header: %w", err)
	}
	encodedHeader := s.base64URLEncode(headerJSON)

	// 2. 序列化声明
	claimsJSON, err := json.Marshal(claims)
	if err != nil {
		return "", fmt.Errorf("failed to marshal claims: %w", err)
	}

	// 3. 派生加密密钥
	key := s.deriveKey(secret)

	// 4. 加密声明
	ciphertext, nonce, err := s.encryptAESGCM(claimsJSON, key)
	if err != nil {
		return "", fmt.Errorf("encryption failed: %w", err)
	}

	// 5. 编码各部分
	encodedCiphertext := s.base64URLEncode(ciphertext)
	encodedNonce := s.base64URLEncode(nonce)

	// 6. 组合 JWE
	jwe := fmt.Sprintf("%s.%s.%s", encodedHeader, encodedNonce, encodedCiphertext)

	return jwe, nil
}

// DecryptJWE 解密并验证 JWT
func (s *sJwt) DecryptJWE(jwe, secret string) (*model.JWTClaims, error) {
	var claims *model.JWTClaims

	// 分割 JWE 的各个部分
	parts := strings.Split(jwe, ".")
	if len(parts) != 3 {
		return claims, errors.New("invalid JWE format")
	}

	// 1. 解码头部
	headerJSON, err := s.base64URLDecode(parts[0])
	if err != nil {
		return claims, fmt.Errorf("failed to decode header: %w", err)
	}

	var header *model.JWEHeader
	if err := json.Unmarshal(headerJSON, &header); err != nil {
		return claims, fmt.Errorf("failed to parse header: %w", err)
	}

	// 验证加密算法
	if header.Enc != "A256GCM" {
		return claims, fmt.Errorf("unsupported encryption algorithm: %s", header.Enc)
	}

	// 2. 解码 nonce
	nonce, err := s.base64URLDecode(parts[1])
	if err != nil {
		return claims, fmt.Errorf("failed to decode nonce: %w", err)
	}

	// 3. 解码密文
	ciphertext, err := s.base64URLDecode(parts[2])
	if err != nil {
		return claims, fmt.Errorf("failed to decode ciphertext: %w", err)
	}

	// 4. 派生解密密钥
	key := s.deriveKey(secret)

	// 5. 解密数据
	plaintext, err := s.decryptAESGCM(ciphertext, key, nonce)
	if err != nil {
		return claims, fmt.Errorf("decryption failed: %w", err)
	}

	// 6. 解析声明
	if err := json.Unmarshal(plaintext, &claims); err != nil {
		return claims, fmt.Errorf("failed to parse claims: %w", err)
	}

	// 7. 验证过期时间
	if claims.Exp < time.Now().Unix() {
		return claims, errors.New("token expired")
	}

	return claims, nil
}

// 实现安全的 Base64URL 编码
func (s *sJwt) base64URLEncode(data []byte) string {
	encoded := base64.RawURLEncoding.EncodeToString(data)
	return encoded
}

// 实现安全的 Base64URL 解码
func (s *sJwt) base64URLDecode(encoded string) ([]byte, error) {
	return base64.RawURLEncoding.DecodeString(encoded)
}

// 从密钥派生加密密钥
func (s *sJwt) deriveKey(secret string) []byte {
	hash := sha256.New()
	hash.Write([]byte(secret))
	return hash.Sum(nil)
}

// 使用 AES-GCM 加密数据
func (s *sJwt) encryptAESGCM(plaintext, key []byte) ([]byte, []byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, nil, err
	}

	ciphertext := gcm.Seal(nil, nonce, plaintext, nil)
	return ciphertext, nonce, nil
}

// 使用 AES-GCM 解密数据
func (s *sJwt) decryptAESGCM(ciphertext, key, nonce []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
