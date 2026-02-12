package rules

import (
	"go/token"
	"unicode"
	"unicode/utf8"

	"github.com/P3rCh1/logcheck/internal/utils"
)

const (
	CapitalizedReport = "log message should not be capitalized"
)

func CheckLowercase(info *utils.LogInfo) (string, token.Pos) {
	for _, msg := range info.MsgParts {
		if len(msg.Data) != 0 {
			r, _ := utf8.DecodeRuneInString(msg.Data)
			if unicode.IsUpper(r) {
				return CapitalizedReport, msg.Pos
			}

			return "", token.NoPos
		}
	}

	return "", token.NoPos
}
