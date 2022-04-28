package main

import (
	"LinkedList/singleList"
	"fmt"
)

func main() {
	l := singleList.CreateList()
	// s := singleList.CreateList()

	l.InsertToEnd(8)
	l.InsertToEnd(7)
	l.InsertToEnd(2)
	l.InsertToEnd(2)
	l.InsertToEnd(7)
	l.InsertToEnd(8)

	// l.PrintList()
	// singleList.CheckIfPalidrome(l.Head)

	fmt.Println(singleList.CheckIfPalidrome(l.Head))
}
