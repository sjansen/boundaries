package main

import (
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

type File struct {
	AbsPath string
	RelPath string

	Package   string
	BuildTags []string
	Imports   []*Import
}

type Import struct {
	Location Location
	Path     string
}

const (
	GoRoot Location = iota + 1
	GoPath
	Project
	Vendor
)

type Location int

func (l Location) String() string {
	switch l {
	case GoRoot:
		return "GoRoot"
	case GoPath:
		return "GoPath"
	case Project:
		return "Project"
	case Vendor:
		return "Vendor"
	}
	return "(invalid)"
}

type Inspector struct {
	bctx build.Context
	root string
}

func NewInspector(root string) (ins *Inspector, err error) {
	root, err = filepath.Abs(root)
	if err != nil {
		return
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return
	}
	ins = &Inspector{
		bctx: build.Default,
		root: root,
	}
	return
}

func (ins *Inspector) Inspect(path string) (file *File, err error) {
	abspath, relpath := path, path
	if filepath.IsAbs(path) {
		relpath, err = filepath.Rel(ins.root, path)
		if err != nil {
			return
		}
	} else {
		abspath = filepath.Join(ins.root, path)
	}
	dir := filepath.Dir(abspath)

	fset := token.NewFileSet()
	f, err := parser.ParseFile(fset, abspath, nil, parser.ParseComments)
	if err != nil {
		return
	}

	result := &File{
		AbsPath:   abspath,
		RelPath:   relpath,
		Package:   f.Name.Name,
		BuildTags: ins.buildTags(f),
	}

	imports, err := ins.imports(dir, f)
	if err != nil {
		return
	}
	result.Imports = imports

	return result, nil
}

func (ins *Inspector) buildTags(f *ast.File) []string {
	var tags []string
	for _, group := range f.Comments {
		// A +build comment is ignored after or adjoining the package declaration.
		if group.End()+1 >= f.Package {
			break
		}
		// "+build" is ignored within or after a /*...*/ comment.
		if !strings.HasPrefix(group.List[0].Text, "//") {
			break
		}
		for _, c := range group.List {
			if !strings.Contains(c.Text, "+build") {
				continue
			}
			line := c.Text
			line = strings.TrimPrefix(line, "//")
			line = strings.TrimSpace(line)
			if strings.HasPrefix(line, "+build") {
				fields := strings.Fields(line)
				if fields[0] != "+build" {
					continue
				}
				tags = append(tags, line)
			}
		}
	}
	return tags
}

func (ins *Inspector) imports(dir string, f *ast.File) ([]*Import, error) {
	var imports []*Import
	for _, spec := range f.Imports {
		p, err := strconv.Unquote(spec.Path.Value)
		if err != nil {
			return nil, err
		}

		imp := &Import{Path: p}
		pkg, err := ins.bctx.Import(p, dir, build.FindOnly)
		if err != nil {
			return nil, err
		}

		switch {
		case pkg.Goroot:
			imp.Location = GoRoot
		case strings.HasPrefix(pkg.Dir, ins.root):
			if strings.Contains(pkg.ImportPath, "/vendor/") {
				imp.Location = Vendor
			} else {
				imp.Location = Project
			}
		default:
			imp.Location = GoPath
		}
		imports = append(imports, imp)
	}
	return imports, nil
}

func main() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}

	s, err := NewInspector(wd)
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
