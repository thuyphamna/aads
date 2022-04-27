package singleList

func (l *List) ReturnKthToLast(k int) interface{} {
	if l.head == nil {
		return nil
	}

	p1 := l.head
	p2 := l.head
	for i := 1; i < k; i++ {
		p1 = p1.next
	}

	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2.value
}
