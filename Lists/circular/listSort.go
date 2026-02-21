package circular

func SortListAscending(l *CircularList) {
	if l.Head == nil {
		return
	}

	var array []int
	current := l.Head.Next
	previous := l.Head

	array = append(array, previous.Data)

	for current != l.Head {
		array = append(array, current.Data)
		current = current.Next
	}

	// Empty the list
	l.Head = nil

	sortArray(array)

	for i := 0; i < len(array); i++ {
		InsertFront(l, array[i])
	}
}

func sortArray(array []int) {
	if array == nil {
		return
	}

	n := len(array)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j+1] > array[j] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
}

func SortListDescending(l *CircularList) {
	if l.Head == nil {
		return
	}

	var array []int
	current := l.Head.Next
	previous := l.Head

	array = append(array, previous.Data)

	for current != l.Head {
		array = append(array, current.Data)
		current = current.Next
	}

	// Empty the list
	l.Head = nil

	sortArray(array)

	for i := 0; i < len(array); i++ {
		InsertBack(l, array[i])
	}
}
