package config

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type AccessTokenConfig struct {
	secret []byte
}

func NewAccessTokenConfig() *AccessTokenConfig {
	return &AccessTokenConfig{
		secret: []byte("secret-secret-secret"),
	}
}

func (config *AccessTokenConfig) Generate(claims jwt.MapClaims) string {
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(config.secret)
	if err != nil {
		panic("error in signed token")
	}

	return tokenString
}

func (config *AccessTokenConfig) FormatClaims(payload map[string]interface{}) jwt.MapClaims {
	mapClaims := jwt.MapClaims{}
	for key, value := range payload {
		mapClaims[key] = value
	}
	return mapClaims
}

func (config *AccessTokenConfig) Validate(accessToken string) (*bool, error) {
	parsedToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return config.secret, nil
	})
	if err != nil {
		result := false
		return &result, errors.New("invalid token")
	}

	if parsedToken.Valid {
		result := true
		return &result, nil
	} else {
		result := false
		return &result, errors.New("invalid token")
	}
}
