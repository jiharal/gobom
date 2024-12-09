package gobom

import (
	"encoding/xml"
)

// CycloneDX BOM format version 1.4
type CycloneDXBOM struct {
	XMLName      xml.Name    `xml:"bom"`
	XMLNS        string      `xml:"xmlns,attr"`
	Version      int         `xml:"version,attr"`
	SerialNumber string      `xml:"serialNumber,attr"`
	Metadata     Metadata    `xml:"metadata"`
	Components   []Component `xml:"components>component"`
	Dependencies []BOMRef    `xml:"dependencies>dependency"`
}

type Metadata struct {
	Timestamp string    `xml:"timestamp"`
	Tools     []Tool    `xml:"tools>tool"`
	Component Component `xml:"component"`
}

type Tool struct {
	Vendor  string `xml:"vendor"`
	Name    string `xml:"name"`
	Version string `xml:"version"`
}

type Component struct {
	Type         string        `xml:"type,attr"`
	BOMRef       string        `xml:"bom-ref,attr"`
	Name         string        `xml:"name"`
	Version      string        `xml:"version"`
	Publisher    string        `xml:"publisher,omitempty"`
	Description  string        `xml:"description,omitempty"`
	Licenses     []License     `xml:"licenses>license,omitempty"`
	PackageURL   string        `xml:"purl,omitempty"`
	ExternalRefs []ExternalRef `xml:"externalReferences>reference,omitempty"`
}

type License struct {
	ID   string `xml:"id,omitempty"`
	Name string `xml:"name,omitempty"`
}

type ExternalRef struct {
	Type string `xml:"type,attr"`
	URL  string `xml:"url"`
}

type BOMRef struct {
	Ref          string   `xml:"ref,attr"`
	Dependencies []string `xml:"dependency,omitempty"`
}
