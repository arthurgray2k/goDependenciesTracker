package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/arthurgray2k/goDependenciesTracker/pkg/goDependenciesTracker"
)

func run(args []string, stdout, stderr io.Writer) int {
	fs := flag.NewFlagSet("goDependenciesTracker", flag.ContinueOnError)
	fs.SetOutput(stderr)

	dir := fs.String("dir", ".", "Go project directory (defaults to current directory)")
	depth := fs.Int("depth", -1, "Maximum depth for dependency tree (-1 for infinite)")
	format := fs.String("format", "cli", "Output format (cli, json, txt, csv, md, markdown)")

	if err := fs.Parse(args); err != nil {
		return 2
	}

	g, err := goDependenciesTracker.BuildGraph(*dir)
	if err != nil {
		fmt.Fprintf(stderr, "Error building graph: %v\n", err)
		return 1
	}

	tree := g.BuildTree(*depth)

	switch *format {
	case "json":
		err = goDependenciesTracker.ExportJSON(stdout, tree)
	case "csv":
		err = goDependenciesTracker.ExportCSV(stdout, tree)
	case "md", "markdown":
		err = goDependenciesTracker.ExportMarkdown(stdout, tree)
	case "txt", "cli":
		err = goDependenciesTracker.ExportTXT(stdout, tree)
	default:
		fmt.Fprintf(stderr, "Unknown format: %s. Defaulting to CLI (txt).\n", *format)
		err = goDependenciesTracker.ExportTXT(stdout, tree)
	}

	if err != nil {
		fmt.Fprintf(stderr, "Error exporting output: %v\n", err)
		return 1
	}

	return 0
}

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}
