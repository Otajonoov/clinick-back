package utils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt"
	"gitlab.com/QuvonchbekOtajonov/clinic-back/config"
)

type TokenParams struct {
	UserName string
	Duration time.Duration
}

func CreateToken(cfg *config.Config, tokenParams *TokenParams) (string, *Payload, error) {
	payload, err := NewPayload(tokenParams)
	if err != nil {
		return "", nil, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(cfg.SecretKey))
	return token, payload, err
}

func VerifyToken(cfg *config.Config, token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(cfg.SecretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)

	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
