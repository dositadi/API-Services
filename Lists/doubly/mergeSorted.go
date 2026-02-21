package doubly

func Merge2SortedLists(l1, l2 *DoublyList) *DoublyList {
	if l1.Head == nil {
		return l2
	}

	if l2.Head == nil {
		return l1
	}

	for l2.Head != nil {
		InsertBack(l1, l2.Head.Data)
		l2.Head = l2.Head.Next
	}

	var array []int

	for l1.Head != nil {
		array = append(array, l1.Head.Data)
		l1.Head = l1.Head.Next
	}

	sort(array)

	for i := 0; i < len(array); i++ {
		InsertFront(l1, array[i])

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
	l1 := &DoublyList{}
	l2 := &DoublyList{}

	InsertBack(l1, 1)
	InsertBack(l1, 2)
	InsertFront(l1, 0)
	InsertBack(l1, 4)
	InsertBack(l2, 5)
	InsertFront(l2, 9)
	InsertBack(l2, 8)

	InsertFront(l1, 6)
	InsertBack(l2, 7)

	/* printList(l1)
	printList(l2)
	*/
	l3 := Merge2SortedLists(l1, l2)

	printList(l3)
}
