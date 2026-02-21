package lists

func MergeSortedList(l1, l2 *List) *List {
	if l2.Head == nil {
		return l1
	}

	if l1.Head == nil {
		return l2
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

	sortArray(array)

	for i := 0; i < len(array); i++ {
		InsertBack(l1, array[i])
	}

	return l1
}

func sortArray(arr []int) {
	if arr == nil {
		return
	}

	n := len(arr)

	for i := 0; i < n; i++ {
		for j := 0; j < n-i-1; j++ {
			if arr[j+1] < arr[j] {
				arr[j+1], arr[j] = arr[j], arr[j+1]
			}
		}
	}
}

func TestMergeSort() {
	l1 := &List{}
	l2 := &List{}

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
	l3 := MergeSortedList(l1, l2)

	printList(l3)
}
