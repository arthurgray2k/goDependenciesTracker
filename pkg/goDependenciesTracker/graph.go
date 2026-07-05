package goDependenciesTracker

import (
	"bufio"
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// Graph represents a dependency graph
type Graph struct {
	Nodes map[string][]string // Module -> List of dependencies
	Root  string
}

// BuildGraph fetches the dependency graph of the go project at the given directory.
func BuildGraph(dir string) (*Graph, error) {
	rootCmd := exec.Command("go", "list", "-m")
	rootCmd.Dir = dir
	rootOut, err := rootCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get root module: %w", err)
	}
	rootMod := strings.TrimSpace(string(rootOut))

	graphCmd := exec.Command("go", "mod", "graph")
	graphCmd.Dir = dir
	graphOut, err := graphCmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to get go mod graph: %w", err)
	}

	g := &Graph{
		Nodes: make(map[string][]string),
		Root:  rootMod,
	}

	scanner := bufio.NewScanner(bytes.NewReader(graphOut))
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		if len(parts) == 2 {
			from := parts[0]
			to := parts[1]
			g.Nodes[from] = append(g.Nodes[from], to)
		}
	}

	return g, nil
}
