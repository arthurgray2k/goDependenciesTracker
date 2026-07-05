package goDependenciesTracker

import (
	"bytes"
	"testing"
)

func TestGraphBuild(t *testing.T) {
	// Simple sanity test against its own directory
	g, err := BuildGraph("..")
	if err != nil {
		t.Fatalf("Failed to build graph: %v", err)
	}
	if g.Root == "" {
		t.Errorf("Expected root module to be non-empty")
	}
}

func TestTreeBuild(t *testing.T) {
	g := &Graph{
		Nodes: map[string][]string{
			"root": {"dep1", "dep2"},
			"dep1": {"dep3"},
		},
		Root: "root",
	}

	// Depth 1 should only get direct children and not children's children
	tree := g.BuildTree(1)
	if len(tree.Children) != 2 {
		t.Errorf("Expected 2 children at depth 1, got %d", len(tree.Children))
	}
	if len(tree.Children[0].Children) != 0 {
		t.Errorf("Expected 0 children at depth 1 (cutoff), got %d", len(tree.Children[0].Children))
	}

	// Infinite depth should traverse everything
	treeAll := g.BuildTree(-1)
	// Find dep1 which has a child
	var dep1Node *TreeNode
	for _, child := range treeAll.Children {
		if child.Module == "dep1" {
			dep1Node = child
			break
		}
	}

	if dep1Node == nil || len(dep1Node.Children) != 1 {
		t.Errorf("Expected 1 child at depth 2 (infinite), got something else")
	}
}

func TestExports(t *testing.T) {
	tree := &TreeNode{
		Module: "root",
		Children: []*TreeNode{
			{Module: "dep1"},
		},
	}

	var buf bytes.Buffer
	err := ExportJSON(&buf, tree)
	if err != nil || buf.Len() == 0 {
		t.Error("JSON export failed")
	}

	buf.Reset()
	err = ExportCSV(&buf, tree)
	if err != nil || buf.Len() == 0 {
		t.Error("CSV export failed")
	}

	buf.Reset()
	err = ExportMarkdown(&buf, tree)
	if err != nil || buf.Len() == 0 {
		t.Error("Markdown export failed")
	}

	buf.Reset()
	err = ExportTXT(&buf, tree)
	if err != nil || buf.Len() == 0 {
		t.Error("TXT export failed")
	}
}
