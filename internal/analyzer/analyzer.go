package analyzer

import (
	"fmt"
	"go/ast"
	"go/token"

	"github.com/P3rCh1/logcheck/internal/rules"
	"github.com/P3rCh1/logcheck/internal/utils"
	"golang.org/x/tools/go/analysis"
)

const (
	StyleCategory    = "style"
	SecurityCategory = "security"
)

type checker struct {
	check    func(*utils.LogInfo) (string, token.Pos)
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
			check:    rules.CheckSensitiveLeak,
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

			info := utils.ExtractLogInfo(pass, call)
			if info == nil {
				return true
			}

			fmt.Println(info)

			for _, checker := range checkers {
				if reportMsg, pos := checker.check(info); pos != token.NoPos {
					pass.Report(
						analysis.Diagnostic{
							Pos:      pos,
							Message:  reportMsg,
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
