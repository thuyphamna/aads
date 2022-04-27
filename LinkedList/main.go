package main

import (
	"LinkedList/singleList"
	"fmt"
)

func main() {
	l := singleList.CreateList()

	l.InsertToEnd(1)
	l.InsertToEnd(2)
	l.InsertToEnd(3)
	l.InsertToEnd(4)
	l.InsertToEnd(5)
	l.InsertToEnd(9)
	l.InsertToEnd(8)

	fmt.Println(l.ReturnKthToLast(8))
	// l.RemoveDups()

	l.PrintList()

}
