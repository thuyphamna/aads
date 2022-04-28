package singleList

func Reverse(n *Node) *Node {
	var headNode *Node

	for n != nil {
		newNode := Node{}
		newNode.value = n.value
		newNode.next = headNode
		headNode = &newNode
		n = n.next
	}
	return headNode
}

func CheckIfPalidrome(n *Node) bool {

	headNodeOfReversedList := Reverse(n)
	for n != nil && headNodeOfReversedList != nil {
		if n.value != headNodeOfReversedList.value {
			return false
		}
		n = n.next
		headNodeOfReversedList = headNodeOfReversedList.next
	}
	return true
}
