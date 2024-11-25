package main

import (
	"os"

	"github.com/rauzh/go-complexity-analysis"
	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() {
	os.Setenv("GOCOMPLEXITY_EXIT_CODE", "228")
	unitchecker.Main(complexity.Analyzer)
}
