package main

import (
	"LinkedList/singleList"
)

func main() {
	l := singleList.CreateList()
	s := singleList.CreateList()

	l.InsertToEnd(7)
	l.InsertToEnd(1)
	// l.InsertToEnd(6)
	// l.InsertToEnd(8)

	s.InsertToEnd(5)
	s.InsertToEnd(9)
	s.InsertToEnd(2)

	newList := singleList.SumList(l.Head, s.Head)
	newList.PrintList()
}
