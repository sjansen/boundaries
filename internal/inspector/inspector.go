package inspector

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"
	"strconv"
	"strings"
)

type Inspector struct {
	bctx build.Context
	root string
}

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

func New(root string) (*Inspector, error) {
	root, err := filepath.Abs(root)
	if err != nil {
		return nil, err
	}
	root, err = filepath.EvalSymlinks(root)
	if err != nil {
		return nil, err
	}
	ins := &Inspector{
		bctx: build.Default,
		root: root,
	}
	return ins, nil
}

func (ins *Inspector) Inspect(filename string) (*File, error) {
	abspath, relpath := filename, filename
	if filepath.IsAbs(filename) {
		tmp, err := filepath.Rel(ins.root, filename)
		if err != nil {
			return nil, err
		}
		relpath = tmp
	} else {
		abspath = filepath.Join(ins.root, filename)
	}

	parsed, err := parser.ParseFile(
		token.NewFileSet(),
		abspath, nil,
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(abspath)
	imports, err := ins.imports(dir, parsed)
	if err != nil {
		return nil, err
	}

	tags := buildTagsFromAST(parsed)
	if tmp := buildTagsFromFilename(filename); tmp != "" {
		tags = append(tags, tmp)
	}
	f := &File{
		AbsPath:   abspath,
		RelPath:   relpath,
		Package:   parsed.Name.Name,
		BuildTags: tags,
		Imports:   imports,
	}

	return f, nil
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
