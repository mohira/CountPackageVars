package main

import (
	"CountPackageVars"

	"golang.org/x/tools/go/analysis/unitchecker"
)

func main() { unitchecker.Main(CountPackageVars.Analyzer) }
