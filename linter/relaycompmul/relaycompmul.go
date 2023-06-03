package relaycompmul

import "github.com/gqlgo/gqlanalysis"

var Analyzer = &gqlanalysis.Analyzer{
	Name: "relaycompmul",
	Doc:  "relaycompmul finds a relay cursor connection fields which not have @complexity directive or its mul argument.",
	Run:  run,
}

func run(pass *gqlanalysis.Pass) (interface{}, error) {
	panic("unimplemented")
}
