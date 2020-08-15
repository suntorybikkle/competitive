package chapter2

type Node struct {
	next *Node
	data int
}

func (node *Node) AppendToTail(d int) {
	end := Node{nil, d}
	n := node
	for ; n.next != nil; n = n.next {
	}
	n.next = &end
}
