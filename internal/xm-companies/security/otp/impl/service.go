package otp

import (
	"time"

	"github.com/google/uuid"
	"xm-companies/config"
	"xm-companies/internal/xm-companies/security/otp"
)

type Service struct {
	lifetime uint
	keys     map[string]*otp.TokenInfo
}

func NewService(conf *config.Config) *Service {
	return &Service{
		lifetime: conf.TokenLifetime,
		keys:     make(map[string]*otp.TokenInfo),
	}
}

func (s *Service) CreateToken(username string) string {
	token := uuid.NewString()
	s.keys[token] = &otp.TokenInfo{
		Assignee:       username,
		ExpirationTime: time.Now().Add(time.Second * time.Duration(s.lifetime)),
	}

	return token
}

func (s *Service) ValidateToken(t string) (bool, *otp.TokenInfo) {
	s.revokeInvalidKeys()
	value, ok := s.keys[t]
	return ok, value
}

func (s *Service) ValidityPeriod() uint {
	return s.lifetime
}

func (s *Service) revokeInvalidKeys() {
	currentTime := time.Now()
	for k, v := range s.keys {
		if currentTime.After(v.ExpirationTime) {
			delete(s.keys, k)
		}
	}
}
