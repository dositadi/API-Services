package circular

import "fmt"

func PrintCircularList(l *CircularList) {
	if l.Head == nil {
		fmt.Println("No Item in the List given.")
		return
	}

	current := l.Head.Next

	fmt.Print(l.Head.Data, " --> ")

	for current != l.Head {
		fmt.Print(current.Data, " --> ")
		current = current.Next
	}
	fmt.Println(l.Tail.Next.Data)
}
