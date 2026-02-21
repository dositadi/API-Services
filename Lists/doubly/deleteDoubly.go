package doubly

import "fmt"

func DeleteBack(l *DoublyList) {
	if l.Head == nil {
		return
	}

	l.Tail.Prev.Next = nil
	l.Tail = l.Tail.Prev
}

func DeleteFront(l *DoublyList) {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next
}

func DeleteRandom(l *DoublyList, position int) {
	if l.Head == nil {
		return
	}

	if position == 1 {
		l.Head = l.Head.Next
		return
	}

	count := 1
	prev := l.Head
	curr := l.Head.Next.Next

	for curr != nil {
		if count == position-1 {
			prev.Next = curr

			if prev.Next == nil {
				l.Tail = prev.Next
			}
			return
		}
		count++
		prev = prev.Next
		curr = curr.Next
	}

	if position-count == 1 {
		prev.Next = nil
		l.Tail = prev
		return
	} else {
		fmt.Println("The position does not exist")
	}
}

func DeleteSpecificItem(l *DoublyList, value int) {
	if l.Head == nil {
		return
	}

	// Ensure That if the item is at the head of the list it deleted
	for l.Head != nil && l.Head.Data == value {
		l.Head = l.Head.Next
	}

	previousNode := l.Head
	currentNode := l.Head.Next

	for currentNode != nil {
		if currentNode.Data == value {
			previousNode.Next = currentNode.Next

			if currentNode == l.Tail {
				l.Tail = previousNode
			}
			currentNode = currentNode.Next
		} else {
			previousNode = currentNode
			currentNode = currentNode.Next
		}
	}
}
