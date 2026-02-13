package rules

import (
	"go/token"

	"github.com/P3rCh1/logcheck/internal/config"
	"github.com/P3rCh1/logcheck/internal/utils"
	"github.com/cloudflare/ahocorasick"
)

var RulesDefault = []string{"capitalize", "english", "symbols", "sensitive"}

const (
	StyleCategory    = "style"
	SecurityCategory = "security"
)

type Checker struct {
	Check    func(*utils.LogInfo) (string, token.Pos)
	Category string
}

func Rules(cfg *config.Config) map[string]Checker {
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

	return map[string]Checker{
		"capitalize": {
			Check:    CheckLowercase,
			Category: StyleCategory,
		},
		"english": {
			Check:    CheckEnglish,
			Category: StyleCategory,
		},
		"symbols": {
			Check: func(info *utils.LogInfo) (string, token.Pos) {
				return CheckNoSymbolsAndEmoji(allowed, info)
			},
			Category: StyleCategory,
		},
		"sensitive": {
			Check: func(info *utils.LogInfo) (string, token.Pos) {
				return CheckSensitiveLeak(ahocorasick.NewStringMatcher(cfg.SensitiveWords), info)
			},
			Category: SecurityCategory,
		},
	}
}
