package main

import (
	"github.com/gqlgo/gqlanalysis/multichecker"

	"github.com/Warashi/compgen/linter/relaycompmul"
)

func main() {
	multichecker.Main(
		relaycompmul.Analyzer,
	)
}
