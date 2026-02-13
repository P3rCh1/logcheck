package rules

import (
	"strings"

	"github.com/P3rCh1/logcheck/internal/utils"
	"github.com/cloudflare/ahocorasick"
	"golang.org/x/tools/go/analysis"
)

var SensitiveWordsDefault = []string{
	"token",
	"key",
	"pwd",
	"bearer",
	"password",
	"secret",
}

const SensitiveLeakReport = "log message should not contains sensitive values"

func CheckSensitiveLeak(matcher *ahocorasick.Matcher, info *utils.LogInfo) *analysis.Diagnostic {
	for _, arg := range info.ArgNames {
		if matcher.Contains([]byte(strings.ToLower(arg.Data))) {
			return &analysis.Diagnostic{
				Pos:      arg.Pos,
				End:      arg.End,
				Message:  SensitiveLeakReport,
				Category: StyleCategory,
			}
		}
	}

	return nil
}
