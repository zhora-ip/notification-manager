package token

import (
	"crypto/rand"
	"encoding/hex"
	"time"
)

type VerificationToken struct {
	Email     string
	Token     string
	ExpiresAt time.Time
}

func GenerateToken(email string) *VerificationToken {
	b := make([]byte, 32)
	rand.Read(b)
	return &VerificationToken{
		Email:     email,
		Token:     hex.EncodeToString(b),
		ExpiresAt: time.Now().Add(24 * time.Hour),
	}
}
