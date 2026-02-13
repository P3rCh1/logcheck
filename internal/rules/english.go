package rules

import (
	"unicode"

	"github.com/P3rCh1/logcheck/internal/utils"
	"golang.org/x/tools/go/analysis"
)

const NotEnglishReport = "log message should be in english"

func CheckEnglish(info *utils.LogInfo) *analysis.Diagnostic {
	for _, msg := range info.MsgParts {
		for _, r := range msg.Data {
			if !unicode.IsLetter(r) {
				continue
			}

			if !unicode.Is(unicode.Latin, r) {
				return &analysis.Diagnostic{
					Pos:      msg.Pos,
					End:      msg.End,
					Message:  NotEnglishReport,
					Category: StyleCategory,
				}
			}
		}
	}

	return nil
}
