package doubly

func ListReverse(l *DoublyList) {
	if l.Head == nil {
		return
	}

	temp := &DoublyList{}

	for l.Head != nil {
		InsertBack(temp, l.Head.Data)
		l.Head = l.Head.Next
	}

	for temp.Head != nil {
		InsertFront(l, temp.Head.Data)
		temp.Head = temp.Head.Next
	}
}
