package main

import (
	"fmt"
	"os"

	"github.com/sjansen/boundaries/internal/inspector"
)

func main() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	s, err := inspector.New(wd)
	if err != nil {
		return
	}

	for _, arg := range os.Args[1:] {
		f, err := s.Inspect(arg)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(f.RelPath)
		fmt.Printf("  %#v\n", f.Package)
		for _, line := range f.BuildTags {
			fmt.Printf("  %#v\n", line)
		}
		for _, imp := range f.Imports {
			fmt.Printf("  %-7s  %#v\n", imp.Location, imp.Path)
		}
	}
}
