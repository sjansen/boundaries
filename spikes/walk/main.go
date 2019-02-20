package main

import (
	"fmt"
	"go/build"
	"os"
	"path/filepath"
	"strings"
)

type Walker struct {
	ctx build.Context
}

func NewWalker() *Walker {
	ctx := build.Default
	ctx.UseAllFiles = true
	return &Walker{ctx: ctx}
}

func (w *Walker) PrintPkg(dir string) (err error) {
	pkg, err := w.ctx.ImportDir(dir, 0)
	if err != nil {
		return
	}

	fmt.Println("tags:", pkg.AllTags)
	fmt.Println("files:", pkg.GoFiles)
	fmt.Println("ignored:", pkg.IgnoredGoFiles)
	fmt.Println("invalid:", pkg.InvalidGoFiles)
	fmt.Println("imports:", pkg.Imports)
	fmt.Println("tests:", pkg.TestGoFiles)
	fmt.Println("imports:", pkg.TestImports)
	fmt.Println("xtest:", pkg.XTestGoFiles)
	fmt.Println("imports:", pkg.XTestImports)

	return
}

func main() {
	w := NewWalker()
	os.Chdir("testdata")
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		switch {
		case err != nil:
			return err
		case !info.IsDir():
			return nil
		case strings.HasSuffix(path, "testdata"):
			return nil
		case strings.HasSuffix(path, "vendor"):
			return nil
		}
		fmt.Println(path)
		err = w.PrintPkg(path)
		fmt.Println("")
		if _, ok := err.(*build.NoGoError); ok {
			return nil
		}
		return err
	})
	if err != nil {
		fmt.Println(err.Error())
	}
}
