package chapter2

import (
	"testing"
)

func TestAppendToTail(t *testing.T) {
	n := Node{nil, 1}
	n.AppendToTail(2)
	if n.data != 1 {
		t.Errorf("n.data = %v\n", n.data)
	}
	if n.next.data != 2 {
		t.Errorf("n.next.data = %v\n", n.next.data)
	}
}
