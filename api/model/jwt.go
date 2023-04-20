package model

import (
	"errors"
	"time"
)

type JWTPayload struct {
	AuthId    string    `json:"authId,omitempty"`
	Username  string    `json:"username,omitempty"`
	Firstname string    `json:"firstname,omitempty"`
	Lastname  string    `json:"lastname,omitempty"`
	Email     string    `json:"email,omitempty"`
	AuthMode  AuthMode  `json:"authMode,omitempty"`
	IssuedAt  time.Time `json:"issuedAt"`
	ExpiredAt time.Time `json:"expiredAt"`
}

var (
	errExpiredToken = errors.New("token has expired")
)

func (J JWTPayload) Valid() error {
	if time.Now().After(J.ExpiredAt) {
		return errExpiredToken
	}
	return nil
}
