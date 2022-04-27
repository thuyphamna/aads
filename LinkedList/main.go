package main

import (
	"LinkedList/singleList"
)

func main() {
	l := singleList.CreateList()

	l.InsertToEnd(3)
	l.InsertToEnd(5)
	l.InsertToEnd(8)
	l.InsertToEnd(5)
	l.InsertToEnd(10)
	l.InsertToEnd(2)
	l.InsertToEnd(1)
	// l.RemoveDups()
	newList := l.Partition(5)
	newList.PrintList()

	// l.PrintList()

}
