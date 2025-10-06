package main

import "testing"

// Definition for a Node.
type Node struct {
	Val       int
	Neighbors []*Node
}

// cloneGraph clones a connected undirected graph
func cloneGraph(node *Node) *Node {
	// TODO: Implement your solution here
	return nil
}

// Helper function to create a graph from adjacency list
func createGraph(adjList [][]int) *Node {
	if len(adjList) == 0 {
		return nil
	}

	nodes := make([]*Node, len(adjList))
	for i := range nodes {
		nodes[i] = &Node{Val: i + 1}
	}

	for i, neighbors := range adjList {
		for _, neighbor := range neighbors {
			nodes[i].Neighbors = append(nodes[i].Neighbors, nodes[neighbor-1])
		}
	}

	return nodes[0]
}

// Helper function to convert graph to adjacency list
func graphToAdjList(node *Node) [][]int {
	if node == nil {
		return [][]int{}
	}

	visited := make(map[*Node]bool)
	var result [][]int

	var dfs func(*Node)
	dfs = func(n *Node) {
		if visited[n] {
			return
		}
		visited[n] = true

		var neighbors []int
		for _, neighbor := range n.Neighbors {
			neighbors = append(neighbors, neighbor.Val)
		}
		result = append(result, neighbors)

		for _, neighbor := range n.Neighbors {
			dfs(neighbor)
		}
	}

	dfs(node)
	return result
}

func TestCloneGraph(t *testing.T) {
	testCases := []struct {
		name     string
		adjList  [][]int
		expected [][]int
	}{
		{
			name:     "Simple connected graph",
			adjList:  [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
			expected: [][]int{{2, 4}, {1, 3}, {2, 4}, {1, 3}},
		},
		{
			name:     "Single node",
			adjList:  [][]int{{}},
			expected: [][]int{{}},
		},
		{
			name:     "Empty graph",
			adjList:  [][]int{},
			expected: [][]int{},
		},
		{
			name:     "Two connected nodes",
			adjList:  [][]int{{2}, {1}},
			expected: [][]int{{2}, {1}},
		},
		{
			name:     "Complex graph",
			adjList:  [][]int{{2, 3}, {1, 4}, {1, 4}, {2, 3}},
			expected: [][]int{{2, 3}, {1, 4}, {1, 4}, {2, 3}},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			original := createGraph(tc.adjList)
			cloned := cloneGraph(original)
			result := graphToAdjList(cloned)

			// Note: This is a simplified comparison. In a real scenario,
			// you'd want to ensure the cloned graph has the same structure
			// but different memory addresses
			if len(result) != len(tc.expected) {
				t.Errorf("cloneGraph(%v) = %v; expected %v", tc.adjList, result, tc.expected)
			}
		})
	}
}
