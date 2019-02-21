package main

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/analysis"
	"golang.org/x/tools/go/analysis/passes/inspect"
	"golang.org/x/tools/go/analysis/singlechecker"
)

var Analyzer = &analysis.Analyzer{
	Name:             "boundaries",
	Doc:              "Check code organization",
	Run:              run,
	Requires:         []*analysis.Analyzer{inspect.Analyzer},
	RunDespiteErrors: true,
}

func run(pass *analysis.Pass) (interface{}, error) {
	fmt.Println(pass.Pkg.Name())
	fmt.Println(pass.OtherFiles)
	for _, f := range pass.Files {
		fmt.Println(pass.Fset.File(f.Pos()).Name())
		for _, group := range f.Comments {
			// A +build comment is ignored after or adjoining the package declaration.
			if group.End()+1 >= f.Package {
				break
			}

			// "+build" is ignored within or after a /*...*/ comment.
			if !strings.HasPrefix(group.List[0].Text, "//") {
				break
			}

			// Check each line of a //-comment.
			for _, c := range group.List {
				if !strings.Contains(c.Text, "+build") {
					continue
				}

				line := c.Text
				line = strings.TrimPrefix(c.Text, "//")
				line = strings.TrimSpace(line)

				if strings.HasPrefix(line, "+build") {
					fields := strings.Fields(line)
					if fields[0] != "+build" {
						continue
					}
					//filename := pass.Fset.File(f.Pos()).Name()
					//fmt.Println(filename)
					//fmt.Println("    ", line)
					pass.Reportf(c.Pos(), "%s", line)
				}
			}
		}
	}
	return nil, nil
}

func main() {
	singlechecker.Main(Analyzer)
}
