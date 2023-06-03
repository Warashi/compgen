package relaycompmul_test

import (
	"testing"

	"github.com/Warashi/compgen/linter/relaycompmul"
	"github.com/gqlgo/gqlanalysis/analysistest"
)

func Test(t *testing.T) {
	testdata := analysistest.TestData(t)
	analysistest.Run(t, testdata, relaycompmul.Analyzer, "a")
}
