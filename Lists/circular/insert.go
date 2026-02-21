package circular

import "fmt"

func InsertBack(l *CircularList, data int) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	l.Tail.Next = newNode
	l.Tail = newNode
	newNode.Next = l.Head
}

func InsertFront(l *CircularList, data int) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	newNode.Next = l.Head
	l.Tail.Next = newNode
	l.Head = newNode
}

func InsertRandom(l *CircularList, position, data int) {
	newNode := &Node{Data: data}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	if position == 1 {
		InsertFront(l, data)
		return
	}

	previous := l.Head
	current := l.Head.Next
	count := 1

	for current != l.Head {
		if count == position-1 {
			break
		}

		count = count + 1
		previous = previous.Next
		current = current.Next
	}

	if position-count > 1 {
		fmt.Println("This position does not exixt")
		return
	}

	if position-count == 1 {
		InsertBack(l, data)
		return
	}

	newNode.Next = current
	previous.Next = newNode
}
