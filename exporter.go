package gobom

import (
	"encoding/json"
	"encoding/xml"
)

// ExportJSON exports the BOM as JSON
func (b *BOM) ExportJSON() ([]byte, error) {
	return json.MarshalIndent(b, "", "  ")
}

// ExportXML exports the BOM as XML
func (b *BOM) ExportXML() ([]byte, error) {
	return xml.MarshalIndent(b, "", "  ")
}
