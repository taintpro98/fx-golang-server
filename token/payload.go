package token

import (
	"fx-golang-server/pkg/e"
	"time"

	"github.com/google/uuid"
)

type Payload struct {
	TokenID   string                 `json:"token_id,omitempty"`
	Email     string                 `json:"email"`
	Audience  string                 `json:"aud,omitempty"`
	ExpiresAt int64                  `json:"exp,omitempty"`
	Id        string                 `json:"jti,omitempty"`
	IssuedAt  int64                  `json:"iat,omitempty"` // seconds
	Issuer    string                 `json:"iss,omitempty"`
	NotBefore int64                  `json:"nbf,omitempty"`
	Subject   string                 `json:"sub,omitempty"`
	Refresh   bool                   `json:"refresh"`
	Data      map[string]interface{} `json:"data,omitempty"`
}

func NewPayload(
	userID string,
	expDuration time.Duration,
	data map[string]interface{},
) (*Payload, error) {
	newID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}
	curTime := time.Now()
	var expiresAtTime int64 = 0
	if expDuration > 0 {
		curTime := time.Now()
		expiresAtTime = curTime.Add(expDuration).Unix()
	}

	tokenID := newID.String()
	return &Payload{
		TokenID:   tokenID,
		Subject:   userID,
		ExpiresAt: expiresAtTime,
		IssuedAt:  curTime.Unix(),
		Data:      data,
	}, nil
}

func (c *Payload) Valid() error {
	if c.ExpiresAt < time.Now().Unix() {
		return e.ErrTokenExpired
	}
	return nil
}

func (p *Payload) SetExpires(expDuration time.Duration) {
	p.ExpiresAt = time.Now().Add(expDuration).Unix()
}

func (p *Payload) SetRefresh(id *string, refresh bool) {
	if id != nil {
		p.TokenID = *id
	}
	p.Refresh = refresh
}
