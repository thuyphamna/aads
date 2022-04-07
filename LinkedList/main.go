package main

import (
	"LinkedList/singleList"
)

func main() {
	l := singleList.CreateList()

	l.InsertToHead(3)
	l.InsertToHead(4)
	l.InsertToHead(5)

	l.PrintList()

}
