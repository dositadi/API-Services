package doubly

func SortListAscending(l *DoublyList) {
	if l.Head == nil {
		return
	}

	var array []int

	for l.Head != nil {
		array = append(array, l.Head.Data)
		l.Head = l.Head.Next
	}

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

func SortListDescending(l *DoublyList) {
	if l.Head == nil {
		return
	}

	var array []int

	for l.Head != nil {
		array = append(array, l.Head.Data)
		l.Head = l.Head.Next
	}

	sortArray(array)

	for i := 0; i < len(array); i++ {
		InsertBack(l, array[i])
	}
}
