package lists

import "fmt"

type Node struct {
	Data int
	Next *Node
}

type List struct {
	Head *Node
	Tail *Node
}

func InsertBack(l *List, data int) {
	newNode := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head = newNode
		l.Tail = newNode
		return
	}
	l.Tail.Next = newNode
	l.Tail = newNode
}

func InsertFront(l *List, data int) {
	newNode := &Node{Data: data, Next: nil}

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	newNode.Next = l.Head
	l.Head = newNode
}

func InsertRandom(l *List, position int, data int) {
	newNode := &Node{Data: data, Next: nil}

	count := 1

	if l.Head == nil {
		l.Head, l.Tail = newNode, newNode
		return
	}

	if position == 1 {
		InsertFront(l, data)
		return
	}

	//create two pointers to track the node before the position and the node at the position
	previous := l.Head
	current := l.Head.Next

	for previous != nil {
		if count == position-1 {
			// The node before the position points to the new node
			previous.Next = newNode
			// The new node points to the node at the position
			newNode.Next = current

			if newNode.Next == nil {
				l.Tail = newNode
			}

			fmt.Println("tail:  ", l.Tail.Data)

			return
		}
		count++
		previous = previous.Next
		current = current.Next
	}

}

func ListSize(l *List) int {
	if l.Head == nil {
		return 1
	}

	count := 0
	for l.Head != nil {
		count++
		l.Head = l.Head.Next
	}
	return count
}
