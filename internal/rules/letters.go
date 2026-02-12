package rules

import (
	"go/token"
	"unicode"

	"github.com/P3rCh1/logcheck/internal/utils"
)

const NotLetterReport = "log message should contains only letters with spaces"

func CheckLetters(info *utils.LogInfo) (string, token.Pos) {
	for _, msg := range info.MsgParts {
		for _, r := range msg.Data {
			if !unicode.IsLetter(r) && !unicode.IsSpace(r) {
				return NotLetterReport, msg.Pos
			}
		}
	}

	return "", token.NoPos
}
