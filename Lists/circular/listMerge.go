package circular

func Merge2Lists(l1, l2 *CircularList) *CircularList {
	if l1.Head == nil {
		return l2
	}

	if l2.Head == nil {
		return l1
	}

	for l2.Head != l2.Tail {
		InsertBack(l1, l2.Head.Data)
		l2.Head = l2.Head.Next
	}
	InsertBack(l1, l2.Tail.Data)
	return l1
}

func TestMerge() {
	l1 := &CircularList{}
	l2 := &CircularList{}

	InsertBack(l1, 1)
	InsertBack(l1, 2)
	InsertFront(l1, 0)
	InsertBack(l1, 4)
	InsertBack(l2, 5)
	InsertFront(l2, 9)
	InsertBack(l2, 8)

	InsertFront(l1, 6)
	InsertBack(l2, 7)

	PrintCircularList(l1)
	PrintCircularList(l2)

	l3 := Merge2Lists(l1, l2)

	PrintCircularList(l3)
}
