package cache

import "fmt"

type Node struct {
	key   int
	value int
	next  *Node
	prev  *Node
}

type DoublyLinkList struct {
	tail *Node
	head *Node
}

func CreateDoublyLinkList() *DoublyLinkList {
	return &DoublyLinkList{}
}

func (dl *DoublyLinkList) RemoveFromFront() *Node {
	oldHead := dl.head
	if dl.head == nil {
		return nil
	} else if dl.head == dl.tail {
		dl.head = nil
		dl.tail = nil
	} else {
		dl.head = dl.head.next
		dl.head.prev = nil
	}
	return oldHead
}

func (dl *DoublyLinkList) AddToEnd(n *Node) {
	newNode := n
	if dl.head == nil {
		dl.head = newNode
		dl.tail = newNode
	} else {
		curr := dl.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
		newNode.prev = curr
		dl.tail = newNode
	}
}

func (dl *DoublyLinkList) Front() *Node {
	return dl.head
}

func (dl *DoublyLinkList) MoveNodeToEnd(n *Node) {
	prevNode := n.prev
	nextNode := n.next

	if nextNode == nil {
		return
	}

	if prevNode != nil && nextNode != nil {
		prevNode.next = nextNode
		nextNode.prev = prevNode
		n.next = nil
	}

	if prevNode == nil {
		dl.head = nextNode
		nextNode.prev = nil
		n.next = nil
	}
	dl.AddToEnd(n)
}

func (dl *DoublyLinkList) PrintList() {
	curr := dl.head
	for curr != nil {
		fmt.Println(curr.key, curr.value)
		curr = curr.next
	}
}
