package rules

import (
	"unicode"

	"github.com/P3rCh1/logcheck/internal/utils"
	"golang.org/x/tools/go/analysis"
)

const SymbolOrEmojiReport = "log message should contains only letters, numbers and spaces"

func isAllowed(allowed map[rune]struct{}, r rune) bool {
	_, isAllowedSymbol := allowed[r]

	return unicode.IsLetter(r) ||
		unicode.IsSpace(r) ||
		unicode.IsDigit(r) ||
		isAllowedSymbol
}

func CheckNoSymbolsAndEmoji(allowed map[rune]struct{}, info *utils.LogInfo) *analysis.Diagnostic {
	for _, msg := range info.MsgParts {
		for _, r := range msg.Data {
			if !isAllowed(allowed, r) {
				return &analysis.Diagnostic{
					Pos:            msg.Pos,
					End:            msg.End,
					Message:        SymbolOrEmojiReport,
					Category:       StyleCategory,
				}
			}
		}
	}

	return nil
}
