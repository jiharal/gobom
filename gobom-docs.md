
### Development Setup

```bash
# Clone the repository
git clone https://github.com/jiharal/gobom.git

# Change into the directory
cd gobom

# Install dependencies
go mod download

# Run tests
go test ./...
```

## Project Structure

```
gobom/
├── bom/
│   ├── types.go       # Core type definitions
│   ├── generator.go   # BOM generation logic
│   ├── license.go     # License detection
│   ├── utils.go       # Utility functions
│   └── exporter.go    # Export functionality
├── go.mod
├── go.sum
├── LICENSE
└── README.md
```

## Common Use Cases

### Basic BOM Generation
```go
package main

import (
    "fmt"
    "github.com/jiharal/gobom/bom"
)

func main() {
    generator := bom.NewGenerator(".")
    bomData, _ := generator.Generate()
    json, _ := bomData.ExportJSON()
    fmt.Println(string(json))
}
```

### Custom Project Path
```go
generator := bom.NewGenerator("/path/to/your/project")
```

### XML Export for Documentation
```go
xmlData, _ := bomData.ExportXML()
err = ioutil.WriteFile("bom.xml", xmlData, 0644)
```

## Error Handling

The library uses standard Go error handling patterns. Common errors you might encounter:

1. Project Path Errors
```go
_, err := bom.NewGenerator("/invalid/path").Generate()
if err != nil {
    // Handle invalid project path
}
```

2. License Detection Errors
```go
// License will be "Unknown" if detection fails
if dep.License == "Unknown" {
    // Handle unknown license case
}
```

## Best Practices

1. Always check error returns:
```go
generator := bom.NewGenerator(projectPath)
bom, err := generator.Generate()
if err != nil {
    log.Fatalf("Failed to generate BOM: %v", err)
}
```

2. Use absolute paths when possible:
```go
absPath, _ := filepath.Abs("./myproject")
generator := bom.NewGenerator(absPath)
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
