package otp

import "time"

type Service interface {
	CreateToken(username string) string
	ValidateToken(t string) (bool, *TokenInfo)
	ValidityPeriod() uint
}

type TokenInfo struct {
	Assignee       string
	ExpirationTime time.Time
}
