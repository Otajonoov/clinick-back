package utils

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	ErrInvalidToken = errors.New("token is isvalid")
	ErrExpiredToken = errors.New("tokes is expired")
)

type Payload struct {
	Id        uuid.UUID `json:"id"`
	UserName  string    `json:"username"`
	IssuedAt  time.Time `json:"issued_at"`
	ExpiredAt time.Time `json:"expired_at"`
}

func NewPayload(tokenParams *TokenParams) (*Payload, error) {
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	payload := &Payload{
		Id:        tokenId,
		UserName:  tokenParams.UserName,
		IssuedAt:  time.Now(),
		ExpiredAt: time.Now().Add(tokenParams.Duration),
	}

	return payload, nil
}

func (payload *Payload) Valid() error {
	if time.Now().After(payload.ExpiredAt) {
		return ErrExpiredToken
	}
	return nil
}
