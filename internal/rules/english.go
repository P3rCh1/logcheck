package rules

import (
	"go/token"
	"unicode"

	"github.com/P3rCh1/logcheck/internal/utils"
)

const NotEnglishReport = "log message should be in english"

func CheckEnglish(info *utils.LogInfo) (string, token.Pos) {
	for _, msg := range info.MsgParts {
		for _, r := range msg.Data {
			if !unicode.IsLetter(r) {
				continue
			}

			if !unicode.Is(unicode.Latin, r) {
				return NotEnglishReport, msg.Pos
			}
		}
	}

	return "", token.NoPos
}
