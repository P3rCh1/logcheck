package rules

import (
	"errors"
	"unicode"
	"unicode/utf8"
)

var (
	ErrEmptyMsg    = errors.New("log message should not be empty")
	ErrCapitalized = errors.New("log message should not be capitalized")
)

func CheckLowercase(msg string) error {
	if len(msg) == 0 {
		return ErrEmptyMsg
	}

	r, _ := utf8.DecodeRuneInString(msg)
	if unicode.IsUpper(r) {
		return ErrCapitalized
	}

	return nil
}
