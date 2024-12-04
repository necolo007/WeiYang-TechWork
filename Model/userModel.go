package Model

import (
	"github.com/dgrijalva/jwt-go"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name       string `gorm:"unique;not null"`
	Email      string `gorm:"unique;not null"`
	Password   string `gorm:"not null"`
	Role       string `gorm:"default:'user'"`
	AdminLevel uint   `gorm:"default:0"`
}

type UserClaims struct {
	UserId   uint
	Username string
	Role     string
	jwt.StandardClaims
}