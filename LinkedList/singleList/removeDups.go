package singleList

func (l *List) RemoveDups() {
	tempMap := make(map[interface{}]bool)

	if l.head == nil {
		return
	}

	current := l.head
	var prev *Node
	for current.next != nil {
		_, exist := tempMap[current.value]
		if !exist {
			tempMap[current.value] = true
			prev = current
		} else {
			prev.next = current.next
		}
		current = current.next
	}
}
