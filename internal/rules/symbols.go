package rules

import (
	"go/token"
	"unicode"

	"github.com/P3rCh1/logcheck/internal/utils"
)

const SymbolOrEmojiReport = "log message should contains only letters, numbers and spaces"

func CheckNoSymbolsAndEmoji(info *utils.LogInfo) (string, token.Pos) {
	for _, msg := range info.MsgParts {
		for _, r := range msg.Data {
			if !unicode.IsLetter(r) && !unicode.IsSpace(r) && !unicode.IsDigit(r) {
				return SymbolOrEmojiReport, msg.Pos
			}
		}
	}

	return "", token.NoPos
}
