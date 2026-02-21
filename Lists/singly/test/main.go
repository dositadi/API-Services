package main

import (
	"fmt"
	"lists"
)

func printList(l *lists.List) {
	if l.Head == nil {
		return
	}

	for l.Head != nil {
		fmt.Print(l.Head.Data, " --> ")
		l.Head = l.Head.Next
	}
	fmt.Println(nil)
}

func main() {
	list := &lists.List{}

	lists.InsertBack(list, 0)
	lists.InsertBack(list, 2)
	lists.InsertFront(list, 1)
	lists.InsertRandom(list, 1, 4)
	lists.InsertRandom(list, 2, 1)
	lists.InsertFront(list, 1)
	lists.InsertBack(list, 1)

	lists.InsertRandom(list, 8, 6)
	lists.InsertBack(list, 1)

	lists.ListReverse(list)

	//lists.SortListAscending(list)
	//lists.SortListDescending(list)

	//fmt.Println(lists.SearchList(list, 2))

	lists.DeleteSpecificItem(list, 1)

	printList(list)

	/* lists.DeleteFront(list)
	lists.DeleteFront(list)

	lists.DeleteBack(list)
	lists.DeleteBack(list)

	lists.InsertBack(list, 2)
	lists.InsertFront(list, 0)
	lists.InsertRandom(list, 1, 4)

	lists.DeleteRandom(list, 2)
	lists.DeleteRandom(list, 3)
	lists.DeleteRandom(list, 4)
	lists.DeleteRandom(list, 5) */

	//printList(list)

	//lists.TestMerge()

	//lists.TestMergeSort()
}
