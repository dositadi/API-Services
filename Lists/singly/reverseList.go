package lists

func ListReverse(l *List) {
	if l.Head == nil {
		return
	}

	temp := &List{}

	for l.Head != nil {
		InsertBack(temp, l.Head.Data)
		l.Head = l.Head.Next
	}

	for temp.Head != nil {
		InsertFront(l, temp.Head.Data)
		temp.Head = temp.Head.Next
	}
}
