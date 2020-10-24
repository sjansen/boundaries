package main

import (
	"flag"
	"fmt"
	"os"
	"sort"

	"golang.org/x/tools/go/packages"
)

type Packages []*packages.Package

func (x Packages) Len() int      { return len(x) }
func (x Packages) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

type ByID struct{ Packages }

func (x ByID) Less(i, j int) bool {
	return x.Packages[i].ID < x.Packages[j].ID
}

func main() {
	flag.Parse()

	cfg := &packages.Config{
		Mode: packages.NeedFiles |
			packages.NeedImports |
			packages.NeedName,
		Tests: true,
	}
	pkgs, err := packages.Load(cfg, flag.Args()...)
	if err != nil {
		fmt.Fprintf(os.Stderr, "load: %v\n", err)
		os.Exit(1)
	}
	if packages.PrintErrors(pkgs) > 0 {
		os.Exit(1)
	}

	sort.Sort(ByID{pkgs})

	for _, pkg := range pkgs {
		imports := make([]string, 0, len(pkg.Imports))
		for k := range pkg.Imports {
			imports = append(imports, k)
		}
		sort.Strings(imports)

		fmt.Println(pkg.ID)
		fmt.Println(" ", pkg.Name)
		fmt.Println("  PkgPath:", pkg.PkgPath)
		fmt.Println("  GoFiles:  ", pkg.GoFiles)
		fmt.Println("  Ignored:", pkg.IgnoredFiles)
		fmt.Println("  Other:  ", pkg.OtherFiles)
		fmt.Println("  Imports:", imports)
		fmt.Println()
	}
}
