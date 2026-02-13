package rules

import (
	"github.com/P3rCh1/logcheck/internal/config"
	"github.com/P3rCh1/logcheck/internal/utils"
	"github.com/cloudflare/ahocorasick"
	"golang.org/x/tools/go/analysis"
)

var RulesDefault = []string{"capitalize", "english", "symbols", "sensitive"}

const (
	StyleCategory    = "style"
	SecurityCategory = "security"
)

type ruleFunc func(*utils.LogInfo) *analysis.Diagnostic

func Rules(cfg *config.Config) map[string]ruleFunc {
	if len(cfg.EnabledRules) == 0 {
		cfg.EnabledRules = RulesDefault
	}

	if len(cfg.SensitiveWords) == 0 {
		cfg.SensitiveWords = SensitiveWordsDefault
	}

	allowed := make(map[rune]struct{}, len(cfg.AllowedSymbols))
	for _, r := range cfg.AllowedSymbols {
		allowed[r] = struct{}{}
	}

	return map[string]ruleFunc{
		"capitalize": CheckLowercase,
		"english":    CheckEnglish,
		"symbols": func(info *utils.LogInfo) *analysis.Diagnostic {
			return CheckNoSymbolsAndEmoji(allowed, info)
		},
		"sensitive": func(info *utils.LogInfo) *analysis.Diagnostic {
			return CheckSensitiveLeak(ahocorasick.NewStringMatcher(cfg.SensitiveWords), info)
		},
	}
}
