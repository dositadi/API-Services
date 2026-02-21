package doubly

func InsertBack(l *DoublyList, data int) {
	newNode := &DoublyNode{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}
	l.Tail.Next = newNode
	newNode.Prev = l.Tail
	l.Tail = newNode
}

func InsertFront(l *DoublyList, data int) {
	newNode := &DoublyNode{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	newNode.Next = l.Head
	l.Head.Prev = newNode
	l.Head = newNode
}

func InsertRandom(l *DoublyList, position int, data int) {
	newNode := &DoublyNode{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	if position == 1 {
		InsertFront(l, data)
		return
	}

	count := 0
	prev := l.Head
	curr := l.Head.Next

	for curr != nil {
		if count == position-2 {
			newNode.Next = curr
			newNode.Prev = prev
			prev.Next = newNode
			curr.Prev = newNode

			if newNode.Next == nil {
				l.Tail = newNode
			}
			return
		}
		count++
		prev = prev.Next
		curr = curr.Next
	}
}
