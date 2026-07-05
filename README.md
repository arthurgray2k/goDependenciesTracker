# goDependenciesTracker

A lightweight, zero-dependency command-line tool and Go library for analyzing and exporting your Go module's dependency tree or graph. It parses the output of `go mod graph` to build an accurate dependency representation and supports multiple output formats (CLI, JSON, CSV, Markdown, TXT).

## Features

- **CLI Tool:** Explore dependencies of any Go module from your terminal.
- **Go Library:** Import `goDependenciesTracker` to programmatically build dependency trees.
- **Multiple Formats:** Export dependency data to JSON, CSV, Markdown lists, or standard text trees.
- **Depth Limiting:** Limit analysis to `n` levels deep, or perform infinite depth analysis to capture everything.
- **No Third-Party Dependencies:** Uses only the standard Go library, keeping the tool extremely light and robust.

## Installation

### To use the CLI tool
You can run this project directly, or compile the binary:
```bash
go build -o goDependenciesTracker ./cmd/goDependenciesTracker
```

### To use as a library
Import it into your project:
```go
import "github.com/arthurgray2k/goDependenciesTracker/pkg/goDependenciesTracker"
```

## CLI Usage

```bash
# Analyze the current directory (infinite depth)
./goDependenciesTracker

# Output as JSON up to a depth of 2
./goDependenciesTracker -dir=/path/to/project -depth=2 -format=json

# Output as Markdown
./goDependenciesTracker -format=md

# Output as CSV
./goDependenciesTracker -format=csv
```

**Flags:**
- `-dir` (default `.`): Go project directory to analyze.
- `-depth` (default `-1`): Maximum depth for dependency traversal. `-1` means infinite.
- `-format` (default `cli`): Output format. Supported options: `cli`, `json`, `txt`, `csv`, `md`, `markdown`.

## Testing & Coverage

This project features an extensive test suite verifying the graph parsing, tree generation, exporter formats, and CLI invocation. Test coverage is **>80%**.

To run the test suite and verify coverage:
```bash
go test -coverprofile=coverage.out ./...
go tool cover -func=coverage.out
```
