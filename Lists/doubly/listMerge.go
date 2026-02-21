package doubly

import "fmt"

func Merge2Lists(l1, l2 *DoublyList) *DoublyList {
	if l1.Head == nil {
		return l2
	}

	if l2.Head == nil {
		return l1
	}

	for l2.Head != nil {
		InsertBack(l1, l2.Head.Data)
		l2.Head = l2.Head.Next
	}
	return l1
}

func printList(l *DoublyList) {
	if l.Tail == nil {
		return
	}

	for l.Tail != nil {
		fmt.Print(l.Tail.Data, " --> ")
		l.Tail = l.Tail.Prev
	}
	fmt.Println(nil)
}

func TestMerge() {
	l1 := &DoublyList{}
	l2 := &DoublyList{}

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

	l3 := Merge2Lists(l1, l2)

	printList(l3)
}
