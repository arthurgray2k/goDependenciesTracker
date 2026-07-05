package goDependenciesTracker

// TreeNode represents a node in the dependency tree
type TreeNode struct {
	Module   string      `json:"module"`
	Children []*TreeNode `json:"children,omitempty"`
}

// BuildTree converts a Graph to a Tree from the root, up to maxDepth.
// Pass -1 to maxDepth for infinite depth.
func (g *Graph) BuildTree(maxDepth int) *TreeNode {
	visited := make(map[string]bool)
	return g.buildTree(g.Root, 0, maxDepth, visited)
}

func (g *Graph) buildTree(mod string, currentDepth, maxDepth int, visited map[string]bool) *TreeNode {
	node := &TreeNode{Module: mod}

	if maxDepth != -1 && currentDepth >= maxDepth {
		return node
	}

	if visited[mod] {
		// Prevent infinite loops if cyclic (though Go mods should be DAGs)
		return node
	}

	// Create a new visited map for this path to allow different paths to reach the same node
	pathVisited := make(map[string]bool)
	for k, v := range visited {
		pathVisited[k] = v
	}
	pathVisited[mod] = true

	for _, dep := range g.Nodes[mod] {
		child := g.buildTree(dep, currentDepth+1, maxDepth, pathVisited)
		node.Children = append(node.Children, child)
	}

	return node
}
