package circular

func CountNodes(l *CircularList) int {
	if l.Head == nil {
		return 0
	}

	count := 1
	current := l.Head

	for current != l.Tail {
		count++
		current = current.Next
	}
	return count
}
