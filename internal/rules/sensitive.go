package rules

import (
	"errors"

	"github.com/cloudflare/ahocorasick"
)

var (
	banWords = []string{
		"token",
		"key",
		"password",
	}

	matcher = ahocorasick.NewStringMatcher(banWords)

	ErrSensitive = errors.New("log message should not contains sensitive values")
)

func CheckSensitive(msg string) error {
	if matcher.Contains([]byte(msg)) {
		return ErrSensitive
	}

	return nil
}
