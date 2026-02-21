package lists

import (
	"fmt"
)

func Merge2List(l1, l2 *List) *List {
	if l2.Head == nil {
		return l1
	}

	if l1.Head == nil {
		return l2
	}

	for l2.Head != nil {
		InsertFront(l1, l2.Head.Data)
		l2.Head = l2.Head.Next
	}
	return l1
}

func printList(l *List) {
	if l.Head == nil {
		return
	}

	for l.Head != nil {
		fmt.Print(l.Head.Data, " --> ")
		l.Head = l.Head.Next
	}
	fmt.Println(nil)
}

func TestMerge() {
	l1 := &List{}
	l2 := &List{}

	InsertBack(l1, 1)
	InsertBack(l1, 2)
	InsertFront(l1, 0)
	InsertBack(l1, 4)
	InsertBack(l2, 5)
	InsertFront(l2, 9)
	InsertBack(l2, 8)

	InsertFront(l1, 6)
	InsertBack(l2, 7)

	/* printList(l1)
	printList(l2)
	*/
	l3 := Merge2List(l1, l2)

	printList(l3)
}
