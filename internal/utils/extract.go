package utils

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

func ExtractMessage(pass *analysis.Pass, call *ast.CallExpr) (string, token.Pos) {
	if len(call.Args) == 0 {
		return "", token.NoPos
	}

	if arg, ok := call.Args[0].(*ast.BasicLit); ok && arg.Kind == token.STRING {
		return unquote(arg.Value), arg.Pos()
	}

	return "", token.NoPos
}

func unquote(s string) string {
	if len(s) < 2 {
		return s
	}

	return s[1 : len(s)-1]
}
