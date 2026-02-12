package rules

import (
	"go/token"
	"strings"

	"github.com/P3rCh1/logcheck/internal/utils"
	"github.com/cloudflare/ahocorasick"
)

var (
	banWords = []string{
		"token",
		"key",
		"pwd",
		"bearer",
		"password",
		"secret",
	}

	matcher = ahocorasick.NewStringMatcher(banWords)

	SensitiveLeakReport = "log message should not contains sensitive values"
)

func CheckSensitiveLeak(info *utils.LogInfo) (string, token.Pos) {
	for _, name := range info.ArgNames {
		if matcher.Contains([]byte(strings.ToLower(name.Data))) {
			return SensitiveLeakReport, name.Pos
		}
	}

	return "", token.NoPos
}
