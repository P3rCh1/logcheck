package analyzer

import (
	"go/ast"

	"github.com/P3rCh1/logcheck/internal/config"
	"github.com/P3rCh1/logcheck/internal/rules"
	"github.com/P3rCh1/logcheck/internal/utils"
	"golang.org/x/tools/go/analysis"
)

func NewAnalyzer(cfg *config.Config) *analysis.Analyzer {
	return &analysis.Analyzer{
		Name: "logcheck",
		Doc:  "reports invalid log messages",
		Run: func(pass *analysis.Pass) (any, error) {
			return run(pass, cfg)
		},
	}
}

func run(pass *analysis.Pass, cfg *config.Config) (any, error) {
	rules := rules.Rules(cfg)

	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !utils.IsLog(pass, call) {
				return true
			}

			info := utils.ExtractLogInfo(pass, call)
			if info == nil {
				return true
			}

			for _, checkerName := range cfg.EnabledRules {
				if checker, ok := rules[checkerName]; ok {
					if diagnostic := checker(info); diagnostic != nil {
						pass.Report(*diagnostic)
					}
				}
			}

			return true
		})
	}

	return nil, nil
}
