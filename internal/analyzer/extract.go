package analyzer

import "go/ast"

func isLoggerCall(call *ast.CallExpr) bool {
    sel, ok := call.Fun.(*ast.SelectorExpr)
    if !ok {
        return false
    }

    methodName := sel.Sel.Name
    pkgIdent, ok := sel.X.(*ast.Ident)
    if !ok {
        return false
    }
    pkgName := pkgIdent.Name

    if pkgName == "log" {
        switch methodName {
        case "Print", "Printf", "Println", "Fatal", "Fatalf", "Panic", "Panicf":
            return true
        }
    }

    if pkgName == "slog" {
        switch methodName {
        case "Info", "Error", "Warn", "Debug", "InfoContext", "ErrorContext", "WarnContext", "DebugContext":
            return true
        }
    }

    
    if pkgName == "logger" { 
         switch methodName {
        case "Info", "Error", "Warn", "Debug", "Fatal",
             "Infof", "Errorf", "Warnf", "Debugf", "Fatalf":
            return true
        }
    }
    
    switch methodName {
    case "Info", "Error", "Warn", "Debug", "Fatal",
         "Infof", "Errorf", "Warnf", "Debugf", "Fatalf":
        return true
    }

    return false
}

func extractLogMessage(call *ast.CallExpr) (string, *ast.BasicLit) {

	if len(call.Args) == 0 {
		return "", nil
	}

	lit, ok := call.Args[0].(*ast.BasicLit)

	if !ok {
		return "", nil
	}

	value := lit.Value

	if len(value) < 2 {
		return value, lit
	}

	return value[1 : len(value)-1], lit
}