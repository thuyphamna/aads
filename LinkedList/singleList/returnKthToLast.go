package singleList

func (l *List) ReturnKthToLast(k int) interface{} {
	if l.Head == nil {
		return nil
	}

	p1 := l.Head
	p2 := l.Head
	for i := 1; i < k; i++ {
		p1 = p1.next
	}

	for p1.next != nil {
		p1 = p1.next
		p2 = p2.next
	}
	return p2.value
}
