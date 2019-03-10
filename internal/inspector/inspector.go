package inspector

import (
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
	abspath := filepath.Join(ins.root, filename)
	parsed, err := parser.ParseFile(
		token.NewFileSet(),
		abspath, nil,
		parser.ParseComments,
	)
	if err != nil {
		return nil, err
	}

	dir := filepath.Dir(abspath)
	var imports []*Import
	for _, spec := range parsed.Imports {
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

	tags := buildTagsFromAST(parsed)
	if tmp := buildTagsFromFilename(abspath); tmp != "" {
		tags = append(tags, tmp)
	}
	f := &File{
		Package:   parsed.Name.Name,
		BuildTags: tags,
		Imports:   imports,
	}

	return f, nil
}
