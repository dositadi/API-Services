package circular

type Node struct {
	Data int
	Next *Node
}

type CircularList struct {
	Head *Node
	Tail *Node
}
