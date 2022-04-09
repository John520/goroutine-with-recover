package main

import (
	"games-trace/proto/protobuf-import/golang.org/x/tools/go/analysis/singlechecker"
	"go-printf-func-name/pkg/analyzer"
)

func main() {
	singlechecker.Main(analyzer.Analyzer)
}