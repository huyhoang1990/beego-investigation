package service

import (
	"fmt"
	"regexp"
)

func NewPasswordService() PasswordService {
	return &passwordService{
		minLength: 8,
	}
}

type PasswordService interface {
	ValidateStrength(password string) error
}

type passwordService struct {
	minLength int
}

var (
	lower   = regexp.MustCompile(`[a-z]`)
	upper   = regexp.MustCompile(`[A-Z]`)
	number  = regexp.MustCompile(`[0-9]`)
	special = regexp.MustCompile(`[^0-9a-zA-Z]`)
)

func (svc *passwordService) ValidateStrength(password string) error {
	length := len([]rune(password))
	if length < svc.minLength {
		return fmt.Errorf("password must be larger or equal to 8 character")
	}

	typ := 0
	if lower.MatchString(password) {
		typ++
	}
	if upper.MatchString(password) {
		typ++
	}
	if number.MatchString(password) {
		typ++
	}
	if special.MatchString(password) {
		typ++
	}

	if typ < 4 {
		return fmt.Errorf("password is not strong enough")
	}
	return nil
}
