package singleList

import "fmt"

type Node struct {
	value interface{}
	next  *Node
}

type List struct {
	head *Node
	len  int
}

func CreateList() *List {
	return &List{}
}

func (l *List) InsertToEnd(val interface{}) {
	newNode := Node{}
	newNode.value = val

	if l.head == nil {
		l.head = &newNode
	} else {
		current := l.head
		for current.next != nil {
			current = current.next
		}
		current.next = &newNode
	}
	l.len++
}
func (l *List) InsertToHead(val interface{}) {
	newNode := Node{}
	newNode.value = val
	newNode.next = l.head
	l.head = &newNode

}
func (l *List) GetLen() int {
	return l.len
}

func (l *List) GetHead() (interface{}, error) {
	if l.head == nil {
		return nil, fmt.Errorf("Single list is empty")
	}
	return l.head.value, nil
}

func (l *List) PrintList() error {
	if l.head == nil {
		return fmt.Errorf("Single list is empty")
	}
	current := l.head
	for current.next != nil {
		fmt.Print("->", current.value)
		current = current.next
	}
	fmt.Print("->", current.value)

	return nil
}

//Insert at a specific position
//Search for a node
//Get a node at a specific position
//Delete node from a linked list
//Delete at a specific position
