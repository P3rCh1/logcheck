package rules

import (
	"errors"
	"unicode"
)

var ErrNotLetter = errors.New("log message should contains only letters with spaces")

func CheckLetters(msg string) error {
	for _, r := range msg {
		if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
			return ErrNotLetter
		}
	}

	return nil
}
