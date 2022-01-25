package CountPackageVars

import (
	"golang.org/x/tools/go/analysis"
)

const doc = "CountPackageVars is ..."

var Analyzer = &analysis.Analyzer{
	Name: "CountPackageVars",
	Doc:  doc,
	Run:  run,
}

func run(pass *analysis.Pass) (interface{}, error) {

	return nil, nil
}
