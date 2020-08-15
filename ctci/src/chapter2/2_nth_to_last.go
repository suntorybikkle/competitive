package chapter2

func NthToLast(n *Node, k int) (*Node, int) {
	if n == nil {
		return nil, 0
	}
	nn, index := NthToLast(n.next, k)
	index++
	if index == k {
		return n, index
	}
	return nn, index
}

func NthToLastRunner(n *Node, k int) *Node {
	runner := n
	for i := 0; i < k; i++ {
		if runner == nil {
			return nil
		}
		runner = runner.next
	}
	for runner != nil {
		runner = runner.next
		n = n.next
	}
	return n
}
