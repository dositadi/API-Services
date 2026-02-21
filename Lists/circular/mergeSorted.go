package circular

func Merge2SortedLists(l1, l2 *CircularList) *CircularList {
	if l1.Head == nil {
		return l2
	}

	if l2.Head == nil {
		return l1
	}

	previous2 := l2.Head
	current2 := l2.Head.Next

	InsertBack(l1, previous2.Data)

	for current2 != l2.Head {
		InsertBack(l1, current2.Data)
		current2 = current2.Next
	}

	var array []int
	previous1 := l1.Head
	current1 := l1.Head.Next
	array = append(array, previous1.Data)

	for current1 != l1.Head {
		array = append(array, current1.Data)
		current1 = current1.Next
	}

	// Empty/Clear the list
	l1.Head = nil

	sort(array)

	for i := 0; i < len(array); i++ {
		InsertBack(l1, array[i])
	}
	return l1
}

func sort(array []int) {
	if array == nil {
		return
	}

	n := len(array)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if array[j+1] < array[j] {
				array[j+1], array[j] = array[j], array[j+1]
			}
		}
	}
}

func TestMergeSorted() {
	l1 := &CircularList{}
	l2 := &CircularList{}

	InsertBack(l1, 1)
	InsertBack(l1, 2)
	InsertFront(l1, 0)
	InsertBack(l1, 4)
	InsertBack(l2, 5)
	InsertFront(l2, 9)
	InsertBack(l2, 8)

	InsertFront(l1, 6)
	InsertBack(l2, 7)

	PrintCircularList(l1)
	PrintCircularList(l2)

	l3 := Merge2SortedLists(l1, l2)

	PrintCircularList(l3)
}
