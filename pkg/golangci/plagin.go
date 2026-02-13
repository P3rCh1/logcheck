package linters

import (
	"fmt"

	"github.com/P3rCh1/logcheck/internal/analyzer"
	"github.com/P3rCh1/logcheck/internal/config"
	"github.com/P3rCh1/logcheck/internal/rules"
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"
)

func init() {
	register.Plugin("logcheck", New)
}

type PluginLogCheck struct {
	Config *config.Config
}

func New(settings any) (register.LinterPlugin, error) {
	fmt.Println(settings)

	cfg, err := register.DecodeSettings[config.Config](settings)
	if err != nil {
		return nil, fmt.Errorf("decode settings: %w", err)
	}

	checkers := rules.Rules(&cfg)

	for _, name := range cfg.EnabledRules {
		if _, ok := checkers[name]; !ok {
			return nil, fmt.Errorf("unknown rule: %q", name)
		}
	}

	return &PluginLogCheck{Config: &cfg}, nil
}

func (p *PluginLogCheck) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{analyzer.NewAnalyzer(p.Config)}, nil
}

func (p *PluginLogCheck) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
