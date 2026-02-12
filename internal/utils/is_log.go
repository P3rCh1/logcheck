package utils

import (
	"go/ast"
	"go/types"

	"golang.org/x/tools/go/analysis"
)

type TypeInfo struct {
	Package string
	Name    string
}

var methods = map[TypeInfo]map[string]struct{}{
	{Package: "log/slog", Name: "Logger"}: {
		"Info": {}, "Debug": {}, "Warn": {}, "Error": {},
	},

	{Package: "log/slog"}: {
		"Info": {}, "Debug": {}, "Warn": {}, "Error": {},
	},

	{Package: "go.uber.org/zap", Name: "Logger"}: {
		"Debug": {}, "Info": {}, "Warn": {}, "Error": {},
		"DPanic": {}, "Panic": {}, "Fatal": {},
	},

	{Package: "go.uber.org/zap", Name: "SugaredLogger"}: {
		"Debug": {}, "Debugf": {}, "Debugw": {}, "Debugln": {},
		"Info": {}, "Infof": {}, "Infow": {}, "Infoln": {},
		"Warn": {}, "Warnf": {}, "Warnw": {}, "Warnln": {},
		"Error": {}, "Errorf": {}, "Errorw": {}, "Errorln": {},
		"DPanic": {}, "DPanicf": {}, "DPanicw": {}, "DPanicln": {},
		"Panic": {}, "Panicf": {}, "Panicw": {}, "Panicln": {},
		"Fatal": {}, "Fatalf": {}, "Fatalw": {}, "Fatalln": {},
	},
}

func IsLog(pass *analysis.Pass, call *ast.CallExpr) bool {
	selector, ok := call.Fun.(*ast.SelectorExpr)
	if !ok {
		return false
	}

	method := selector.Sel.Name
	typeInfo := unwrap(pass, selector.X)

	if typeInfo == nil {
		return false
	}

	curSet, ok := methods[*typeInfo]
	if !ok {
		return false
	}

	_, ok = curSet[method]

	return ok
}

func unwrap(pass *analysis.Pass, expr ast.Expr) *TypeInfo {
	switch e := expr.(type) {
	case *ast.SelectorExpr:
		return unwrap(pass, e.X)

	case *ast.Ident:
		if pkgName, ok := pass.TypesInfo.Uses[e].(*types.PkgName); ok {
			return &TypeInfo{Package: pkgName.Imported().Path()}
		}

		if obj := pass.TypesInfo.ObjectOf(e); obj != nil {
			return getTypeInfo(obj.Type())
		}

	case *ast.CallExpr:
		if tv, ok := pass.TypesInfo.Types[e]; ok {
			return getTypeInfo(tv.Type)
		}
	}

	return nil
}

func getTypeInfo(t types.Type) *TypeInfo {
	if t == nil {
		return nil
	}

	for {
		if ptr, ok := t.(*types.Pointer); ok {
			t = ptr.Elem()
			continue
		}

		break
	}

	named, ok := t.(*types.Named)
	if !ok {
		return nil
	}

	return &TypeInfo{
		Package: named.Obj().Pkg().Path(),
		Name:    named.Obj().Name(),
	}
}
