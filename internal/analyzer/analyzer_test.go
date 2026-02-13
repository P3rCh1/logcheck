package analyzer

import (
	"testing"

	"github.com/P3rCh1/logcheck/internal/config"
	"golang.org/x/tools/go/analysis/analysistest"
)

func TestAnalyzer(t *testing.T) {
	cfg := config.Config{
		AllowedSymbols: "%",
	}

	analysistest.Run(t, analysistest.TestData(), NewAnalyzer(&cfg), ".")
}
