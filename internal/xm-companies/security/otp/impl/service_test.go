package otp

import (
	"fmt"
	"testing"
	"time"
	"xm-companies/config"
	"xm-companies/internal/xm-companies/security/otp"
)

func TestService_CreateToken(t *testing.T) {
	keysMock := make(map[string]*otp.TokenInfo)
	target := &Service{
		lifetime: 1,
		keys:     keysMock,
	}

	token := target.CreateToken("test")

	if tokenInfo, ok := keysMock[token]; !ok || tokenInfo.Assignee != "test" {
		t.Error(fmt.Sprintf("Created token (%s) assigned to another user '%s'", token, tokenInfo.Assignee))
	}
}

func TestService_ValidateToken(t *testing.T) {
	target := NewService(&config.Config{TokenLifetime: 1})

	token := target.CreateToken("test")
	if isValid, _ := target.ValidateToken(token); !isValid {
		t.Error("Right token is invalid!")
	}

	expiredToken := target.CreateToken("test")
	time.Sleep(time.Second * 2)
	if isValid, tokenInfo := target.ValidateToken(expiredToken); isValid {
		t.Error(fmt.Sprintf("Expired token should be revoked: %v", tokenInfo))
	}
}

func TestService_ValidityPeriod(t *testing.T) {
	target := NewService(&config.Config{TokenLifetime: 1})

	if v := target.ValidityPeriod(); v != 1 {
		t.Error(fmt.Sprintf("Expected validityPeriod is '1'. Returned: '%d'", v))
	}
}
