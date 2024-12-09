# GoBOM - Go Bill of Materials Generator

GoBOM is a Go library for generating Bill of Materials (BOM) for Go projects. It automatically detects and tracks dependencies, licenses, and project relationships.

[![Go Reference](https://pkg.go.dev/badge/github.com/jiharal/gobom.svg)](https://pkg.go.dev/github.com/jiharal/gobom)
[![Go Report Card](https://goreportcard.com/badge/github.com/jiharal/gobom)](https://goreportcard.com/report/github.com/jiharal/gobom)

## Features

- Automatic dependency detection
- License identification
- Direct/transitive dependency classification
- Repository URL detection
- Multiple export formats (JSON, XML)
- SPDX license identifier support

## Installation

```bash
go get github.com/jiharal/gobom
```

## Quick Start

```go
package main

import (
    "fmt"
    "log"

    "github.com/jiharal/gobom/bom"
)

func main() {
    // Create a new BOM generator
    generator := bom.NewGenerator(".")

    // Generate BOM
    bomData, err := generator.Generate()
    if err != nil {
        log.Fatal(err)
    }

    // Export as JSON
    jsonData, err := bomData.ExportJSON()
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(string(jsonData))
}
```

## API Reference

### Types

#### BOM

```go
type BOM struct {
    GeneratedAt  string       // Timestamp when BOM was generated
    ProjectName  string       // Name of the analyzed project
    Dependencies []Dependency // List of project dependencies
}
```

#### Dependency

```go
type Dependency struct {
    Name       string // Package name
    Version    string // Package version
    License    string // Detected license
    DirectDep  bool   // Whether it's a direct dependency
    Repository string // Repository URL
}
```

### Generator

#### NewGenerator

```go
func NewGenerator(projectPath string) *Generator
```

Creates a new BOM generator for the specified project path.

Example:

```go
generator := bom.NewGenerator("/path/to/project")
```

#### Generate

```go
func (g *Generator) Generate() (*BOM, error)
```

Generates a BOM for the project.

Example:

```go
bom, err := generator.Generate()
if err != nil {
    log.Fatal(err)
}
```

### Export Functions

#### ExportJSON

```go
func (b *BOM) ExportJSON() ([]byte, error)
```

Exports the BOM as formatted JSON.

Example:

```go
jsonData, err := bom.ExportJSON()
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(jsonData))
```

#### ExportXML

```go
func (b *BOM) ExportXML() ([]byte, error)
```

Exports the BOM as formatted XML.

Example:

```go
xmlData, err := bom.ExportXML()
if err != nil {
    log.Fatal(err)
}
fmt.Println(string(xmlData))
```

## Example Output

### JSON Format

```json
{
  "generated_at": "2024-12-09T10:30:00Z",
  "project_name": "example-project",
  "dependencies": [
    {
      "name": "github.com/example/package",
      "version": "v1.2.3",
      "license": "MIT",
      "direct": true,
      "repository": "https://github.com/example/package"
    }
  ]
}
```

### XML Format

```xml
<?xml version="1.0" encoding="UTF-8"?>
<bom>
  <generatedAt>2024-12-09T10:30:00Z</generatedAt>
  <projectName>example-project</projectName>
  <dependencies>
    <dependency>
      <name>github.com/example/package</name>
      <version>v1.2.3</version>
      <license>MIT</license>
      <direct>true</direct>
      <repository>https://github.com/example/package</repository>
    </dependency>
  </dependencies>
</bom>
```

## License Detection

GoBOM supports detection of the following licenses:

- MIT
- Apache-2.0
- GPL-2.0
- GPL-3.0
- BSD-2-Clause
- BSD-3-Clause
- LGPL-3.0
- MPL-2.0
- ISC

The library looks for license files in the following patterns:

- LICENSE
- LICENSE.txt
- LICENSE.md
- COPYING
- COPYING.txt
- COPYING.md

It also supports SPDX license identifiers in source files.

## Contributing

Contributions are welcome! Here's how you can help:

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request
