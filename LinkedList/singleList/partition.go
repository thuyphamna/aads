package singleList

func (l *List) Partition(val int) *List {
	beforeList := CreateList()
	afterList := CreateList()

	current := l.Head
	for current != nil {
		if current.value < val {
			beforeList.InsertToEnd(current.value)
		} else {
			afterList.InsertToEnd(current.value)
		}
		current = current.next
	}

	pointer := beforeList.Head
	for pointer != nil {
		pointer = pointer.next
		if pointer.next == nil {
			pointer.next = afterList.Head
			break
		}
	}
	return beforeList
}
