package gobom

import (
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
)

func (g *Generator) isDirectDependency(depName string) bool {
	goModPath := filepath.Join(g.projectPath, "go.mod")
	fset := token.NewFileSet()
	file, err := parser.ParseFile(fset, goModPath, nil, parser.ImportsOnly)
	if err != nil {
		return false
	}

	for _, imp := range file.Imports {
		if strings.Trim(imp.Path.Value, "\"") == depName {
			return true
		}
	}
	return false
}

func (g *Generator) getRepositoryURL(depName string) string {
	if strings.Contains(depName, "github.com") {
		return "https://" + depName
	}
	return depName
}
