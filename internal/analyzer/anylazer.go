package analyzer

import (
	"go/ast"

	"golang.org/x/tools/go/analysis"
)

var Analyzer = &analysis.Analyzer{
	Name: "loglint",
	Doc:  "checks log messages",
	Run:  run,
}

var appConfig = LoadConfig()

func init() {
	SetSensitiveWords(appConfig.SensitiveWords)
}

func run(pass *analysis.Pass) (interface{}, error) {
	for _, file := range pass.Files {
		ast.Inspect(file, func(node ast.Node) bool {
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}

			if !isLoggerCall(call) {
				return true
			}

			msg, lit := extractLogMessage(call)

			if msg == "" {
				return true
			}

			checkRules(pass, call, msg, lit)

			return true
		})
	}

	return nil, nil
}