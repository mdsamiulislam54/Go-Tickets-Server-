package auth

import "github.com/golang-jwt/jwt/v5"

type JwtClaims struct {
	UserId uint   `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}