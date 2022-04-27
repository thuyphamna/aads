package singleList

func (l *List) Partition(val int) *List {
	beforeList := CreateList()
	afterList := CreateList()

	current := l.head
	for current != nil {
		if current.value < val {
			beforeList.InsertToEnd(current.value)
		} else {
			afterList.InsertToEnd(current.value)
		}
		current = current.next
	}

	pointer := beforeList.head
	for pointer != nil {
		pointer = pointer.next
		if pointer.next == nil {
			pointer.next = afterList.head
			break
		}
	}
	return beforeList
}
