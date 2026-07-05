package goDependenciesTracker

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// ExportJSON writes the tree in JSON format
func ExportJSON(w io.Writer, tree *TreeNode) error {
	encoder := json.NewEncoder(w)
	encoder.SetIndent("", "  ")
	return encoder.Encode(tree)
}

// ExportTXT writes the tree in an indented text format suitable for CLI
func ExportTXT(w io.Writer, tree *TreeNode) error {
	var walk func(node *TreeNode, indent string, isLast bool)
	walk = func(node *TreeNode, indent string, isLast bool) {
		fmt.Fprintf(w, "%s%s\n", indent, node.Module)
		for i, child := range node.Children {
			walk(child, indent+"  ", i == len(node.Children)-1)
		}
	}
	walk(tree, "", true)
	return nil
}

// ExportCSV writes the tree in CSV format (Module, Dependency, Depth)
func ExportCSV(w io.Writer, tree *TreeNode) error {
	writer := csv.NewWriter(w)
	defer writer.Flush()

	if err := writer.Write([]string{"Module", "Dependency", "Depth"}); err != nil {
		return err
	}

	var walk func(node *TreeNode, depth int) error
	walk = func(node *TreeNode, depth int) error {
		for _, child := range node.Children {
			if err := writer.Write([]string{node.Module, child.Module, fmt.Sprintf("%d", depth+1)}); err != nil {
				return err
			}
			if err := walk(child, depth+1); err != nil {
				return err
			}
		}
		return nil
	}
	return walk(tree, 0)
}

// ExportMarkdown writes the tree in Markdown list format
func ExportMarkdown(w io.Writer, tree *TreeNode) error {
	var walk func(node *TreeNode, depth int)
	walk = func(node *TreeNode, depth int) {
		indent := strings.Repeat("  ", depth)
		fmt.Fprintf(w, "%s- %s\n", indent, node.Module)
		for _, child := range node.Children {
			walk(child, depth+1)
		}
	}
	walk(tree, 0)
	return nil
}
