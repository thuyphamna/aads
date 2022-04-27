package singleList

func SumList(firstNum *Node, secondNum *Node) *List {
	var currFirst, currSec = firstNum, secondNum
	newList := CreateList()
	num := 0

	for currFirst != nil && currSec != nil {
		val := currFirst.value + currSec.value + num
		if val >= 10 {
			newList.InsertToEnd(val % 10)
			num = 1
		} else {
			newList.InsertToEnd(val)
			num = 0
		}
		currFirst = currFirst.next
		currSec = currSec.next
	}

	if currFirst != nil {
		newList.InsertToEnd(currFirst.value + num)
		num = 0
		currFirst = currFirst.next
		for currFirst != nil {
			newList.InsertToEnd(currFirst.value)
			currFirst = currFirst.next
		}
	}

	if currSec != nil {
		newList.InsertToEnd(currSec.value + num)
		num = 0
		currSec = currSec.next
		for currSec != nil {
			newList.InsertToEnd(currSec.value)
			currSec = currSec.next
		}
	}

	if num != 0 {
		newList.InsertToEnd(num)
	}

	return newList

}
