package rules

import (
	"unicode"
	"unicode/utf8"

	"github.com/P3rCh1/logcheck/internal/utils"
	"golang.org/x/tools/go/analysis"
)

const CapitalizedReport = "log message should not be capitalized"

func CheckLowercase(info *utils.LogInfo) *analysis.Diagnostic {
	for _, msg := range info.MsgParts {
		if len(msg.Data) != 0 {
			r, size := utf8.DecodeRuneInString(msg.Data)
			if unicode.IsUpper(r) {
				quote := `"`
				if msg.IsRawQuote {
					quote = "`"
				}

				fix := quote + string(unicode.ToLower(r)) + msg.Data[size:] + quote

				return &analysis.Diagnostic{
					Pos:      msg.Pos,
					End:      msg.End,
					Message:  CapitalizedReport,
					Category: StyleCategory,
					SuggestedFixes: []analysis.SuggestedFix{
						{
							Message: "convert first letter to lowercase",
							TextEdits: []analysis.TextEdit{
								{
									Pos:     msg.Pos,
									End:     msg.End,
									NewText: []byte(fix),
								},
							},
						},
					},
				}
			}

			return nil
		}
	}

	return nil
}
