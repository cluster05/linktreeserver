package middleware

import (
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	
	"github.com/cluster05/linktree/api/appresponse"
	"github.com/cluster05/linktree/api/config"
	"github.com/cluster05/linktree/api/model"
)

var (
	errInvalidToken = errors.New("token is invalid")
	errExpiredToken = errors.New("token has expired")
)

func Auth(c *gin.Context) {
	jwtToken := extractToken(c)

	payload, err := verifyToken(jwtToken)

	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewAuthorizationError(err.Error()))
		c.Abort()
		return
	}

	user, err := json.Marshal(*payload)
	if err != nil {
		c.JSON(http.StatusOK, appresponse.NewInternalError(err.Error()))
	}
	c.Header("user", string(user))
	c.Next()
}

func extractToken(c *gin.Context) string {
	queryToken := c.Query("token")
	if queryToken != "" {
		return queryToken
	}

	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func verifyToken(reqToken string) (*model.JWTPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errInvalidToken
		}
		return []byte(config.AppConfig.JWTSecret), nil
	}

	jwtToken, err := jwt.ParseWithClaims(reqToken, &model.JWTPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, errExpiredToken) {
			return nil, errExpiredToken
		}
		return nil, errInvalidToken
	}

	payload, ok := jwtToken.Claims.(*model.JWTPayload)
	if !ok {
		return nil, errInvalidToken
	}

	return payload, nil
}
