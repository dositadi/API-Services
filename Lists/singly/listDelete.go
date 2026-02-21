package lists

import "fmt"

func DeleteFront(l *List) {
	if l.Head == nil {
		return
	}
	l.Head = l.Head.Next
}

func DeleteBack(l *List) {
	if l.Head == nil {
		return
	}

	// Create a current node that will be ahead of the previous node
	current := l.Head.Next.Next
	// Create a previous node that will be behind the current node
	previous := l.Head

	// Loop through the list, until the current node becomes nil
	for current != nil {
		previous = previous.Next
		current = current.Next
	}
	// let the previous node now point to nil, thus making the current node nil
	previous.Next = nil
	// make the previous node the tail
	l.Tail = previous
}

func DeleteRandom(l *List, position int) {
	if l.Head == nil {
		return
	}

	current := l.Head.Next
	previous := l.Head
	count := 1

	fmt.Println("Current: ", current.Data, "\nPrevious: ", previous.Data)

	for previous != nil {
		if count == position-1 {
			previous.Next = current.Next

			if previous.Next == nil {
				l.Tail = previous
			}
			return
		}
		count++
		current = current.Next
		previous = previous.Next
	}
}

func DeleteSpecificItem(l *List, value int) {
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
