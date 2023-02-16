package token

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type JWTClaims struct {
	*jwt.RegisteredClaims
	UserInfo interface{}
}

type ProviderInterface interface {
	GenerateToken(sub string, userInfo interface{}) (string, error)
	ValidateToken(tokenString string) (map[string]interface{}, error)
}

type Provider struct {
	secret []byte
}

// GenerateToken ...
// Ref: https://medium.com/@nooraldinahmed/very-basic-jwt-authentication-with-golang-3516b21c2740
func (p *Provider) GenerateToken(sub string, userInfo interface{}) (string, error) {
	// Get the token instance with the Signing method
	token := jwt.New(jwt.GetSigningMethod("HS256"))
	// Choose an expiration time. Shorter the better
	exp := time.Now().Add(time.Hour * 24)
	// Add your claims
	token.Claims = &JWTClaims{
		&jwt.RegisteredClaims{
			// Set the exp and sub claims. sub is usually the userID
			ExpiresAt: jwt.NewNumericDate(exp),
			Subject:   sub,
		},
		userInfo,
	}
	// Sign the token with your secret key
	val, err := token.SignedString(p.secret)

	if err != nil {
		// On error return the error
		return "", err
	}
	// On success return the token string
	return val, nil
}

func (p *Provider) ValidateToken(tokenString string) (map[string]interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return p.secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}

func NewTokenProvider(secret []byte) (ProviderInterface, error) {
	return &Provider{
		secret: secret,
	}, nil
}
