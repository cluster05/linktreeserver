package service

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/cluster05/linktree/api/config"
	"github.com/cluster05/linktree/api/model"
)

func generateToken(auth model.Auth) (string, error) {
	payload := model.JWTPayload{
		AuthId:    auth.AuthId,
		Username:  auth.Username,
		Firstname: auth.Firstname,
		Lastname:  auth.Lastname,
		Email:     auth.Email,
		AuthMode:  auth.AuthMode,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(time.Duration(config.AppConfig.TokenExpireDuration) * time.Minute),
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	return jwtToken.SignedString([]byte(config.AppConfig.JWTSecret))
}
