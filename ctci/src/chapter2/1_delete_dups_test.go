package chapter2

import (
	"reflect"
	"testing"
)

func TestDeleteDups(t *testing.T) {
	var tests = []struct {
		input []int
		want  []int
	}{
		{[]int{0, 1, 1, 1, 2, 3, 3}, []int{0, 1, 2, 3}},
		{[]int{1, 2, 0, 1, 2, 3}, []int{1, 2, 0, 3}},
		{[]int{0}, []int{0}},
		{[]int{0, 1, 2}, []int{0, 1, 2}},
	}
	t.Run("Set algorithm", func(t *testing.T) {
		for _, test := range tests {
			n := InitNode(test.input)
			DeleteDups(&n)
			actual := NodeToArray(&n)
			if !reflect.DeepEqual(actual, test.want) {
				t.Fatalf("Expected: %v, actual: %v\n", test.want, actual)
			}
		}
	})
	t.Run("Runner algorithm", func(t *testing.T) {
		for _, test := range tests {
			n := InitNode(test.input)
			DeleteDupsRunner(&n)
			actual := NodeToArray(&n)
			if !reflect.DeepEqual(actual, test.want) {
				t.Fatalf("Expected: %v, actual: %v\n", test.want, actual)
			}
		}
	})
}

func InitNode(input []int) Node {
	n := Node{nil, input[0]}
	for _, v := range input[1:] {
		n.AppendToTail(v)
	}
	return n
}

func NodeToArray(n *Node) []int {
	var array []int
	for ; n != nil; n = n.next {
		array = append(array, n.data)
	}
	return array
}
