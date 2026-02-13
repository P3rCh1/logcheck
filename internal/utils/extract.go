package utils

import (
	"go/ast"
	"go/token"

	"golang.org/x/tools/go/analysis"
)

var edgeMethods = map[TypeInfo]map[string]struct{}{
	{Package: "log/slog"}: {
		"String": {}, "Int": {}, "Int64": {}, "Uint64": {},
		"Float64": {}, "Bool": {}, "Time": {}, "Duration": {},
		"Any": {}, "Group": {},
	},

	{Package: "go.uber.org/zap"}: {
		"String": {}, "Strings": {}, "Stringp": {},
		"Int": {}, "Ints": {}, "Int64": {}, "Int64s": {},
		"Uint": {}, "Uints": {}, "Uint64": {}, "Uint64s": {},
		"Float32": {}, "Float32s": {}, "Float64": {}, "Float64s": {},
		"Bool": {}, "Bools": {},
		"Time": {}, "Times": {},
		"Duration": {}, "Durations": {},
		"Any": {}, "Error": {}, "Namespace": {},
		"Binary": {}, "ByteString": {},
	},
}

type ItemAST struct {
	Data string
	Pos  token.Pos
}

type LogInfo struct {
	MsgParts []ItemAST
	ArgNames []ItemAST
}

func ExtractLogInfo(pass *analysis.Pass, call *ast.CallExpr) *LogInfo {
	if len(call.Args) == 0 {
		return nil
	}

	msg := &LogInfo{}

	msg.MsgParts = extractMsgParts(call.Args[0])

	for _, arg := range call.Args {
		items := extractArgNames(pass, arg)
		msg.ArgNames = append(msg.ArgNames, items...)
	}

	return msg
}

func extractArgNames(pass *analysis.Pass, e ast.Expr) []ItemAST {
	var args []ItemAST

	var unwrap func(ast.Expr)
	unwrap = func(e ast.Expr) {
		switch node := e.(type) {
		case *ast.Ident:
			args = append(args, ItemAST{node.Name, e.Pos()})

		case *ast.SelectorExpr:
			args = append(args, ItemAST{node.Sel.Name, e.Pos()})

		case *ast.IndexExpr:
			if ident, ok := node.X.(*ast.Ident); ok {
				args = append(args, ItemAST{ident.Name, e.Pos()})
			}

		case *ast.BinaryExpr:
			if node.Op == token.ADD {
				unwrap(node.X)
				unwrap(node.Y)
			}

		case *ast.CallExpr:

			if ident, ok := node.Fun.(*ast.Ident); ok {
				args = append(args, ItemAST{ident.Name, e.Pos()})
			} else if sel, ok := node.Fun.(*ast.SelectorExpr); ok {
				if obj := pass.TypesInfo.ObjectOf(sel.Sel); obj != nil && obj.Pkg() != nil {
					typeInfo := TypeInfo{Package: obj.Pkg().Path()}
					if methods, ok := edgeMethods[typeInfo]; ok {
						if _, ok := methods[obj.Name()]; ok {
							for _, arg := range node.Args {
								unwrap(arg)
							}

							return
						}
					}
				}

				args = append(args, ItemAST{sel.Sel.Name, e.Pos()})
			}
		}
	}

	unwrap(e)
	return args
}

func extractMsgParts(e ast.Expr) []ItemAST {
	var parts []ItemAST

	var unwrap func(ast.Expr)
	unwrap = func(e ast.Expr) {
		switch node := e.(type) {
		case *ast.BasicLit:
			if node.Kind == token.STRING {
				parts = append(parts, ItemAST{unquote(node.Value), e.Pos()})
			}

		case *ast.BinaryExpr:
			if node.Op == token.ADD {
				unwrap(node.X)
				unwrap(node.Y)
			}
		}
	}

	unwrap(e)
	return parts
}

func unquote(s string) string {
	if len(s) < 2 {
		return s
	}

	return s[1 : len(s)-1]
}
