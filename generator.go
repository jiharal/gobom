package gobom

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

// Generator represents a BOM generator instance
type Generator struct {
	projectPath string
}

// NewGenerator creates a new BOM generator for the specified project path
func NewGenerator(projectPath string) *Generator {
	return &Generator{
		projectPath: projectPath,
	}
}

// Generate creates a new BOM for the project
func (g *Generator) Generate() (*BOM, error) {
	bom := &BOM{
		GeneratedAt:  time.Now().UTC().Format(time.RFC3339),
		ProjectName:  filepath.Base(g.projectPath),
		Dependencies: make([]Dependency, 0),
	}

	// Get module information
	cmd := exec.Command("go", "list", "-m", "all")
	cmd.Dir = g.projectPath
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("error getting module info: %v", err)
	}

	// Parse module output
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if line == "" || !strings.Contains(line, " ") {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) < 2 {
			continue
		}

		dep := Dependency{
			Name:       parts[0],
			Version:    parts[1],
			DirectDep:  g.isDirectDependency(parts[0]),
			Repository: g.getRepositoryURL(parts[0]),
			License:    g.detectLicense(parts[0]),
		}

		bom.Dependencies = append(bom.Dependencies, dep)
	}

	return bom, nil
}
