package lists

func SearchList(l *List, value int) (bool, int) {
	if l.Head == nil {
		return false, 0
	}

	var count int = 1

	for l.Head != nil {
		if l.Head.Data == value {
			return true, count
		}
		count++
		l.Head = l.Head.Next
	}
	return false, 0
}
