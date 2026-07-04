package auth

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	jwtSecretKey         = "jwtDefaultKey1234"
	defaultTokenDuration = 24 * time.Hour
)

type JwtService interface {
	GenerateToken(userId uint, name string, email string) (string, error)
	ValidateToken(tokenStr string) (*JwtClaims, error)
}

type jwtService struct {
	secretKey     string
	tokenDuration time.Duration
}

func NewJwtService(secretKey string, tokenDuration time.Duration) JwtService {
	if secretKey == "" {
		secretKey = jwtSecretKey
	}

	if tokenDuration == 0 {
		tokenDuration = defaultTokenDuration
	}
	return &jwtService{
		secretKey:     secretKey,
		tokenDuration: tokenDuration,
	}
}

func (js *jwtService) GenerateToken(userId uint, name string, email string) (string, error) {
	claims := &JwtClaims{
		UserId: userId,
		Name:   name,
		Email:  email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(js.tokenDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "gotickets",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte(js.secretKey))

}

func (js *jwtService) ValidateToken(tokenStr string) (*JwtClaims, error) {
	token, err := jwt.ParseWithClaims(
		tokenStr,
		&JwtClaims{},
		func(token *jwt.Token) (any, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwt.ErrTokenSignatureInvalid
			}

			return []byte(js.secretKey), nil
		},
	)

	if err != nil {
		return nil, fmt.Errorf("unexpected signing method: %w", err)
	}

	claims, ok := token.Claims.(*JwtClaims)

	if !ok || !token.Valid {
		return nil, jwt.ErrTokenInvalidClaims
	}

	return claims, nil
}
