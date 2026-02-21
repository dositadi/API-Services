package circular

import (
	"fmt"
)

func DeleteFront(l *CircularList) {
	if l.Head == nil {
		return
	}

	l.Head = l.Head.Next
	l.Tail.Next = l.Head
}

func DeleteBack(l *CircularList) {
	if l.Head == nil {
		return
	}

	previous := l.Head
	current := l.Head.Next.Next

	for current != l.Head {
		previous = previous.Next
		current = current.Next
	}

	l.Tail = previous
	previous.Next = l.Head
}

func DeleteRandom(l *CircularList, position, data int) {
	if l.Head == nil {
		return
	}

	if position == 1 {
		DeleteFront(l)
		return
	}

	previous := l.Head
	current := l.Head.Next.Next
	count := 1

	for current != l.Head {
		if count == position-1 {
			break
		}
		count++
		previous = previous.Next
		current = current.Next
	}
	fmt.Println("Prev: ", previous.Data, "\nCurrent: ", current.Data)

	if position-count > 1 {
		fmt.Println("The position does not exist.")
		return
	}

	if position-count == 1 {
		DeleteBack(l)
		return
	}
	previous.Next = current
}

func DeleteAllSpecific(l *CircularList, data int) {
	if l.Head == nil {
		return
	}

	for l.Head != l.Tail && l.Head.Data == data {
		l.Head = l.Head.Next
		l.Tail.Next = l.Head
	}

	previous := l.Head
	current := l.Head.Next

	for current != l.Head {
		if current.Data == data {
			previous.Next = current.Next
			current = current.Next
		} else {
			previous = current
			current = current.Next
		}
	}
}
