package gobom

import (
	"crypto/rand"
	"encoding/hex"
	"encoding/xml"
	"fmt"
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

// Utility function to generate a random serial number
func generateSerialNumber() (string, error) {
	bytes := make([]byte, 16)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return fmt.Sprintf("urn:uuid:%s", hex.EncodeToString(bytes)), nil
}

// ExportCycloneDX exports the BOM in CycloneDX format
func (b *BOM) ExportCycloneDX() ([]byte, error) {
	cdx, err := b.ToCycloneDX()
	if err != nil {
		return nil, fmt.Errorf("failed to convert to CycloneDX: %v", err)
	}

	return xml.MarshalIndent(cdx, "", "  ")
}
