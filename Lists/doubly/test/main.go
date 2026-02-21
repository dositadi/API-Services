package main

import (
	d "doubly"
	"fmt"
)

func printListForward(l *d.DoublyList) {
	if l.Head == nil {
		return
	}

	for l.Head != nil {
		fmt.Print(l.Head.Data, " --> ")
		l.Head = l.Head.Next
	}
	fmt.Println(nil)
}

func printListBackward(l *d.DoublyList) {
	if l.Tail == nil {
		return
	}

	for l.Tail != nil {
		fmt.Print(l.Tail.Data, " --> ")
		l.Tail = l.Tail.Prev
	}
	fmt.Println(nil)
}

func main() {
	l := &d.DoublyList{}

	d.InsertBack(l, 1)
	d.InsertBack(l, 1)
	d.InsertBack(l, 1)

	d.InsertFront(l, 1)
	d.InsertFront(l, 4)
	d.InsertFront(l, 5)

	d.InsertRandom(l, 1, 1)
	d.InsertRandom(l, 6, 9)
	d.InsertRandom(l, 8, 10)
	/*
		d.DeleteBack(l)
		d.DeleteBack(l)

		d.DeleteFront(l)
		d.DeleteFront(l)

		d.InsertBack(l, 5)

		d.DeleteRandom(l, 6)
		d.DeleteRandom(l, 5) */

	fmt.Println(l.Tail.Data)
	fmt.Println(l.Head.Data)

	//doubly.SortListAscending(l)
	//doubly.SortListDescending(l)
	//doubly.ListReverse(l)
	//fmt.Println(d.SearchList(l, 5))
	d.DeleteSpecificItem(l, 1)

	printListForward(l)
	//printListBackward(l)

	//doubly.TestMerge()
	d.TestMergeSorted()
}
