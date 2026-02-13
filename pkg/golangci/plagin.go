package linters

import (
	"github.com/golangci/plugin-module-register/register"
	"golang.org/x/tools/go/analysis"

	"github.com/P3rCh1/logcheck/internal/analyzer"
)

func init() {
	register.Plugin("logcheck", New)
}

type PluginLogCheck struct{}

func New(settings any) (register.LinterPlugin, error) {
	return &PluginLogCheck{}, nil
}

func (p *PluginLogCheck) BuildAnalyzers() ([]*analysis.Analyzer, error) {
	return []*analysis.Analyzer{
		analyzer.Analyzer,
	}, nil
}

func (p *PluginLogCheck) GetLoadMode() string {
	return register.LoadModeTypesInfo
}
