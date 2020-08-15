package chapter2

import (
	"testing"
)

func TestNthToLast(t *testing.T) {

	var tests = []struct {
		node []int
		kth  int
		want int
	}{
		{[]int{0, 1, 2, 3, 4, 5}, 1, 5},
		{[]int{0, 1, 2, 3, 4, 5}, 3, 3},
		{[]int{0, 1, 2, 3, 4, 5}, 6, 0},
		{[]int{0}, 1, 0},
		{[]int{3, 2, -3, 4, 4, 10, 22}, 2, 10},
	}

	t.Run("Recursion algorithm", func(t *testing.T) {
		for _, test := range tests {
			n := InitNode(test.node)
			actual, _ := NthToLast(&n, test.kth)
			if actual.data != test.want {
				t.Fatalf("Expected: %v, actual: %v\n", test.want, actual.data)
			}
		}
	})

	t.Run("Runner algorithm", func(t *testing.T) {
		for _, test := range tests {
			n := InitNode(test.node)
			actual := NthToLastRunner(&n, test.kth)
			if actual.data != test.want {
				t.Fatalf("Expected: %v, actual: %v\n", test.want, actual.data)
			}
		}
	})
}
