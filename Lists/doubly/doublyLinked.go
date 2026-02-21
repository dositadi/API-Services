package doubly

type DoublyNode struct {
	Data int
	Prev *DoublyNode
	Next *DoublyNode
}

type DoublyList struct {
	Head *DoublyNode
	Tail *DoublyNode
}
