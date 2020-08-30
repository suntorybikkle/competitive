package chapter4

import (
	"testing"
)

func TestDeleteDups(t *testing.T) {
	var tests = []struct {
		nodeCount int
		initGraph map[int][]int
		inputs    [][]int
		wants     []bool
	}{
		{5, map[int][]int{0: {1, 2}, 1: {2, 3}, 3: {4}}, [][]int{{0, 1}, {1, 2}, {3, 4}, {0, 4}, {2, 3}, {3, 0}}, []bool{true, true, true, true, false, false}},
		{3, map[int][]int{0: {1}}, [][]int{{0, 1}, {1, 0}}, []bool{true, false}},
	}
	t.Run("Set algorithm", func(t *testing.T) {
		for _, test := range tests {
			g := InitGraph(test.nodeCount, test.initGraph)
			for i := range test.inputs {
				input := test.inputs[i]
				expected := test.wants[i]
				actual := ExistRoot(&g, g.nodes[input[0]], g.nodes[input[1]])
				if expected != actual {
					t.Fatalf("Expected: %v, actual: %v\n", expected, actual)
				}
			}
		}
	})
}

func InitGraph(nodeCount int, initGraph map[int][]int) Graph {
	g := Graph{}
	for i := 0; i < nodeCount; i++ {
		g.nodes = append(g.nodes, &Node{name: "a"})
	}
	for parent, childs := range initGraph {
		for _, child := range childs {
			g.nodes[parent].children = append(g.nodes[parent].children, g.nodes[child])
		}
	}
	return g
}
