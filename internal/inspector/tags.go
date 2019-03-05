package inspector

import (
	"go/ast"
	"strings"
)

// SOURCE: go tool dist list

var archTags = []string{
	"386", "amd64", "amd64p32",
	"arm", "arm64",
	"mips", "mips64", "mips64le", "mipsle",
	"ppc64", "ppc64le",
	"s390x",
	"wasm",
}

var osTags = []string{
	"aix",
	"android",
	"darwin",
	"dragonfly",
	"freebsd",
	"js",
	"linux",
	"nacl",
	"netbsd",
	"openbsd",
	"plan9",
	"solaris",
	"windows",
}

func buildTagsFromAST(f *ast.File) []string {
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
				tags = append(tags, fields[1:]...)
			}
		}
	}
	return tags
}

func buildTagsFromFilename(filename string) string {
	switch {
	case strings.HasSuffix(filename, "_test.go"):
		filename = filename[:len(filename)-8]
	case !strings.HasSuffix(filename, ".go"):
		return ""
	default:
		filename = filename[:len(filename)-3]
	}

	var tags []string
	filenameLength := len(filename)
	for _, tag := range archTags {
		tagLength := len(tag)
		if filenameLength <= tagLength || !strings.HasSuffix(filename, tag) {
			continue
		}
		if filename[filenameLength-tagLength-1] == '_' {
			filename = filename[:filenameLength-tagLength-1]
			filenameLength = len(filename)
			tags = append(tags, tag)
			break
		}
	}
	for _, tag := range osTags {
		tagLength := len(tag)
		if filenameLength <= tagLength || !strings.HasSuffix(filename, tag) {
			continue
		}
		if filename[filenameLength-tagLength-1] == '_' {
			tags = append(tags, tag)
			break
		}
	}
	return strings.Join(tags, ",")
}
