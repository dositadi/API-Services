package circular

func ListReverse(l *CircularList) {
	if l.Head == nil {
		return
	}

	temp := &CircularList{}
	previous := l.Head
	current := l.Head.Next

	InsertBack(temp, previous.Data)

	for current != l.Head {
		InsertBack(temp, current.Data)
		current = current.Next
	}

	// Clear/Empty the List
	l.Head = nil
	tempPrevious := temp.Head
	tempCurrent := temp.Head.Next

	InsertFront(l, tempPrevious.Data)

	for tempCurrent != temp.Head {
		InsertFront(l, tempCurrent.Data)
		tempCurrent = tempCurrent.Next
	}
}
