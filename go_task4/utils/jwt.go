// utils/jwt.go
package utils

import (
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte("DFJKLJOI4U54J5IUERKLJKLJE") // 替换为你的密钥

// Claims结构体
type Claims struct {
	ID       uint   `json:"id"`
	UserName string `json:"userName"`
	jwt.RegisteredClaims
}

// 生成 JWT Token
func GenerateToken(username string, userId uint) (string, error) {
	claims := Claims{
		ID:       userId,
		UserName: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(2 * time.Hour)), // 过期时间
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gin-jwt-demo",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

// 解析 Token
func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
