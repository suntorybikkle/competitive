package chapter3

import (
	"reflect"
	"testing"
)

func TestNew(t *testing.T) {
	var tests = []struct {
		numberOfStacks int
		defaultSize    int
		want           []int
	}{
		{3, 2, []int{0, 0, 0, 0, 0, 0}},
		{1, 1, []int{0}},
	}
	for _, test := range tests {
		actual := NewMultiStack(test.numberOfStacks, test.defaultSize)
		if !reflect.DeepEqual(actual.values, test.want) {
			t.Fatalf("Expected: %v, actual: %v\n", test.want, actual.values)
		}
	}
}

func TestPush(t *testing.T) {
	var tests = []struct {
		numberOfStacks int
		defaultSize    int
		stackNums      []int
		values         []int
		want           []int
	}{
		{3, 3, []int{0, 0}, []int{1, 1}, []int{1, 1, 0, 0, 0, 0, 0, 0, 0}},
		{3, 3, []int{0, 1}, []int{1, 1}, []int{1, 0, 0, 1, 0, 0, 0, 0, 0}},
		{2, 3, []int{1, 1, 1, 1}, []int{1, 2, 3, 4}, []int{4, 0, 0, 1, 2, 3}},
		{2, 3, []int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}, []int{4, 5, 0, 1, 2, 3}},
		{3, 3, []int{0, 0, 0, 2, 2, 2, 2}, []int{1, 1, 1, 2, 3, 4, 5}, []int{5, 1, 1, 1, 0, 0, 2, 3, 4}},
	}
	for _, test := range tests {
		multiStack := NewMultiStack(test.numberOfStacks, test.defaultSize)
		for i := range test.values {
			multiStack.Push(test.stackNums[i], test.values[i])
		}
		actual := multiStack.values
		if !reflect.DeepEqual(actual, test.want) {
			t.Fatalf("Expected: %v, actual: %v\n", test.want, actual)
		}
	}
}

func TestPop(t *testing.T) {
	multiStack := NewMultiStack(3, 3)
	multiStack.Push(0, 1)
	multiStack.Push(0, 2)
	multiStack.Push(1, 1)
	actual := multiStack.Peek(0)
	expected := 2
	if actual != expected {
		t.Fatalf("Expected: %v, actual: %v\n", expected, actual)
	}
	actual = multiStack.Peek(1)
	expected = 1
	if actual != expected {
		t.Fatalf("Expected: %v, actual: %v\n", expected, actual)
	}
	actual, _ = multiStack.Pop(0)
	expected = 2
	if actual != expected {
		t.Fatalf("Expected: %v, actual: %v\n", expected, actual)
	}
	actual = multiStack.Peek(0)
	expected = 1
	if actual != expected {
		t.Fatalf("Expected: %v, actual: %v\n", expected, actual)
	}
}
