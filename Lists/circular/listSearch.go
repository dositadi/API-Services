package circular

func SearchList(l *CircularList, value int) (bool, int) {
	if l.Head == nil {
		return false, 0
	}

	var count int = 1

	current := l.Head.Next
	previous := l.Head

	if previous.Data == value {
		return true, count
	}

	for current != l.Head {
		if current.Data == value {
			return true, count + 1
		}
		count++
		current = current.Next
	}
	return false, 0
}
