package gobom

import "encoding/xml"

type Dependency struct {
	Name       string `json:"name" xml:"name"`
	Version    string `json:"version" xml:"version"`
	License    string `json:"license" xml:"license"`
	DirectDep  bool   `json:"direct" xml:"direct"`
	Repository string `json:"repository" xml:"repository"`
}

type BOM struct {
	XMLName      xml.Name     `xml:"bom"`
	GeneratedAt  string       `json:"generated_at" xml:"generatedAt"`
	ProjectName  string       `json:"project_name" xml:"projectName"`
	Dependencies []Dependency `json:"dependencies" xml:"dependencies>dependency"`
}
