package models

type Token struct {
	Token     string `json:"token" db:"token"`
	ExpiresAt string `json:"expiresAt" db:"expires_at"`
}
