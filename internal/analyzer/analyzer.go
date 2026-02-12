package analyzer

import (
	"go/ast"

	"github.com/P3rCh1/logcheck/internal/rules"
	"github.com/P3rCh1/logcheck/internal/utils"
	"golang.org/x/tools/go/analysis"
)

const (
	StyleCategory    = "style"
	SecurityCategory = "security"
)

type checker struct {
	check    func(string) error
	category string
}

var (
	Analyzer = &analysis.Analyzer{
		Name: "logcheck",
		Doc:  "reports invalid log messages",
		Run:  run,
	}

	checkers = []checker{
		{
			check:    rules.CheckLowercase,
			category: StyleCategory,
		},
		{
			check:    rules.CheckEnglish,
			category: StyleCategory,
		},
		{
			check:    rules.CheckLetters,
			category: StyleCategory,
		},
		{
			check:    rules.CheckSensitive,
			category: SecurityCategory,
		},
	}
)

func run(pass *analysis.Pass) (any, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(n ast.Node) bool {
			call, ok := n.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !utils.IsLog(pass, call) {
				return true
			}

			msg, pos := utils.ExtractMessage(pass, call)
			if msg == "" {
				return true
			}

			for _, checker := range checkers {
				if err := checker.check(msg); err != nil {
					pass.Report(
						analysis.Diagnostic{
							Pos:      pos,
							Message:  err.Error(),
							Category: checker.category,
						},
					)
				}
			}

			return true
		})
	}

	return nil, nil
}
