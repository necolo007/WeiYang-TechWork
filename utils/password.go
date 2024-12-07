package utils

import (
	M "WeiYangWork/Model"
	"crypto/rand"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

const DefaultSecret = "HduHelpMember"

// GenerateRandomKey 生成随机数密钥
func GenerateRandomKey() string {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	if err != nil {
		log.Fatal("无法生成随机密钥:", err)
	}
	return base64.StdEncoding.EncodeToString(key)
}

// GenerateToken 生成Token，时效一天
func GenerateToken(user M.User) (string, error) {
	ExpirationTime := time.Now().Add(24 * time.Hour).Unix()
	claims := &M.UserClaims{
		UserId:   user.ID,
		Role:     user.Role,
		Username: user.Name,
		StandardClaims: jwt.StandardClaims{
			NotBefore: time.Now().Unix() - 60,
			ExpiresAt: ExpirationTime,
			Issuer:    "cxr",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(DefaultSecret))
}
