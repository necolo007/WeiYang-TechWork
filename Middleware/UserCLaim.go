package Middleware

import (
	"WeiYangWork/Model"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// AuthMiddleware JWT验证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "未识别到token!"})
			c.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(authHeader, &Model.UserClaims{}, func(token *jwt.Token) (interface{}, error) {
			return []byte("HduHelpMember"), nil
		})
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token!", "mistake": err.Error()})
			c.Abort()
			return
		}
		if claims, ok := token.Claims.(*Model.UserClaims); ok && token.Valid {
			if time.Now().Unix() > claims.ExpiresAt {
				c.JSON(http.StatusUnauthorized, gin.H{"error": "Token已过期!"})
				c.Abort()
				return
			}
			c.Set("UserClaims", claims)
			c.Next()
		} else {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Token!"})
			c.Abort()
			return
		}
	}
}
