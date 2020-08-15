package chapter2

func DeleteDups(n *Node) {
	set := make(map[int]bool)
	var previous *Node
	for ; n != nil; n = n.next {
		if set[n.data] {
			previous.next = n.next
		} else {
			set[n.data] = true
			previous = n
		}
	}
}

func DeleteDupsRunner(n *Node) {
	current := n
	for ; current != nil; current = current.next {
		runner := current
		for runner.next != nil {
			if current.data == runner.next.data {
				runner.next = runner.next.next
			} else {
				runner = runner.next
			}
		}
	}
}
