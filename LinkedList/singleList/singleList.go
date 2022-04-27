package singleList

import "fmt"

type Node struct {
	value int
	next  *Node
}

type List struct {
	Head *Node
	len  int
}

func CreateList() *List {
	return &List{}
}

func (l *List) InsertToEnd(val int) {
	newNode := Node{}
	newNode.value = val

	if l.Head == nil {
		l.Head = &newNode
	} else {
		current := l.Head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
	l.len++
}
func (l *List) InsertToHead(val int) {
	newNode := Node{}
	newNode.value = val
	newNode.next = l.Head
	l.Head = &newNode
	l.len++
}
func (l *List) GetLen() int {
	return l.len
}

func (l *List) PrintList() error {
	if l.Head == nil {
		return fmt.Errorf("Single list is empty")
	}
	current := l.Head
	for current.next != nil {
		fmt.Print("->", current.value)
		current = current.next
	}
	fmt.Print("->", current.value)

	return nil
}

func (l *List) InsertAt(pos int, val int) error {
	newNode := Node{}
	newNode.value = val

	if l.Head == nil && pos != 0 {
		return fmt.Errorf("Single list is empty")
	}

	if pos == 0 {
		newNode.next = l.Head
		l.Head = &newNode
		return nil
	}

	current := l.Head
	for i := 0; i < l.len; i++ {
		if i == pos {
			newNode.next = current.next
			current.next = &newNode
			return nil
		}
		current = current.next
	}
	return fmt.Errorf("Postion is not valid for this list")
}

func (l *List) SearchFor(val int) error {
	if l.Head == nil {
		return fmt.Errorf("Single list is empty")
	}
	current := l.Head
	for current.next != nil {
		if current.value == val {
			fmt.Printf("Found node %v in the list", val)
			return nil
		}
		current = current.next
	}
	return fmt.Errorf("Can't find this value in the list")
}

func (l *List) Delete(val int) error {
	if l.Head == nil {
		return fmt.Errorf("Single list")
	}

	current := l.Head
	if current.value == val {
		l.Head = current.next
		return nil
	}
	for current.next != nil {
		if current.next.value == val {
			if current.next.next != nil {
				current.next = current.next.next
				return nil
			} else {
				current.next = nil
				return nil
			}
		}
		current = current.next
	}
	return fmt.Errorf("Can't find node with that value")
}

func (l *List) DeleteAt(pos int) error {
	if l.Head == nil {
		return fmt.Errorf("Single list is empty")
	}

	if pos >= l.len {
		return fmt.Errorf("Can't find the position in the list")
	}

	current := l.Head
	for i := 0; i < l.len; i++ {
		if pos == 0 {
			l.Head = current.next
			return nil
		}
		if i+1 == pos {
			if current.next.next != nil {
				current.next = current.next.next
			} else {
				current.next = nil
			}
			return nil
		}
		current = current.next
	}
	return nil
}
