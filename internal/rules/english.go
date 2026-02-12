package rules

import (
	"errors"
	"unicode"
)

var ErrNotEnglish = errors.New("log message should be in english")

func CheckEnglish(msg string) error {
	for _, r := range msg {
		if !unicode.IsLetter(r) {
			continue
		}

		if !((r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z')) {
			return ErrNotEnglish
		}
	}

	return nil
}
