// This file is a plugin for golangci-lint
// Build with: go build -o loglint.so -buildmode=plugin plugin/golangci-lint/main.go

package main

import (
	"github.com/weizis/loglint/internal/analyzer"
	"golang.org/x/tools/go/analysis"
)

func AnalyzerPlugin() *analysis.Analyzer {
	return analyzer.Analyzer
}