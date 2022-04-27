package singleList

func DeleteMiddleNode(n Node) bool {
	if n.next == nil {
		return false
	}

	n.value = n.next.value
	n.next = n.next.next
	return true
}
