package gobom

import (
	"fmt"
	"strings"
	"time"
)

// ToCycloneDX converts the BOM to CycloneDX format
func (b *BOM) ToCycloneDX() (*CycloneDXBOM, error) {
	serialNumber, err := generateSerialNumber()
	if err != nil {
		return nil, fmt.Errorf("failed to generate serial number: %v", err)
	}

	cdx := &CycloneDXBOM{
		XMLNS:        "http://cyclonedx.org/schema/bom/1.4",
		Version:      1,
		SerialNumber: serialNumber,
		Metadata: Metadata{
			Timestamp: time.Now().UTC().Format(time.RFC3339),
			Tools: []Tool{
				{
					Vendor:  "jiharal",
					Name:    "gobom",
					Version: "1.0.0", // Update this with your version
				},
			},
			Component: Component{
				Type:    "application",
				BOMRef:  fmt.Sprintf("pkg:golang/%s@%s", b.ProjectName, "1.0.0"),
				Name:    b.ProjectName,
				Version: "1.0.0", // You might want to extract this from go.mod
			},
		},
		Components:   make([]Component, 0),
		Dependencies: make([]BOMRef, 0),
	}

	// Convert dependencies to components
	for _, dep := range b.Dependencies {
		component := Component{
			Type:    "library",
			BOMRef:  fmt.Sprintf("pkg:golang/%s@%s", dep.Name, dep.Version),
			Name:    dep.Name,
			Version: dep.Version,
		}

		// Add license if available
		if dep.License != "Unknown" {
			component.Licenses = []License{
				{
					ID: dep.License,
				},
			}
		}

		// Add repository as external reference
		if dep.Repository != "" {
			component.ExternalRefs = []ExternalRef{
				{
					Type: "vcs",
					URL:  dep.Repository,
				},
			}
		}

		// Create package URL (purl)
		component.PackageURL = fmt.Sprintf("pkg:golang/%s@%s",
			strings.TrimPrefix(dep.Name, "github.com/"),
			dep.Version)

		cdx.Components = append(cdx.Components, component)
	}

	// Add dependencies relationships
	mainRef := fmt.Sprintf("pkg:golang/%s@%s", b.ProjectName, "1.0.0")
	deps := make([]string, 0)
	for _, dep := range b.Dependencies {
		if dep.DirectDep {
			deps = append(deps, fmt.Sprintf("pkg:golang/%s@%s", dep.Name, dep.Version))
		}
	}

	cdx.Dependencies = append(cdx.Dependencies, BOMRef{
		Ref:          mainRef,
		Dependencies: deps,
	})

	// Add transitive dependencies
	for _, dep := range b.Dependencies {
		depRef := fmt.Sprintf("pkg:golang/%s@%s", dep.Name, dep.Version)
		transDeps := make([]string, 0)

		// Here you would add logic to determine transitive dependencies
		// This would require enhancing the dependency analysis in the Generator

		if len(transDeps) > 0 {
			cdx.Dependencies = append(cdx.Dependencies, BOMRef{
				Ref:          depRef,
				Dependencies: transDeps,
			})
		}
	}

	return cdx, nil
}
