package main

import (
	c "circular"
)

func main() {
	l := &c.CircularList{}

	c.InsertBack(l, 5)
	c.InsertBack(l, 4)
	c.InsertBack(l, 3)
	c.InsertBack(l, 1)
	c.InsertBack(l, 10)
	c.InsertBack(l, 6)
	c.InsertBack(l, 9)
	c.InsertBack(l, 0)
	c.InsertBack(l, 7)
	c.InsertFront(l, 8)
	c.InsertFront(l, 2)

	/* c.DeleteFront(l)
	c.DeleteFront(l)

	c.DeleteBack(l)
	c.DeleteBack(l)
	c.DeleteBack(l) */

	c.InsertRandom(l, 7, 11)

	//c.DeleteRandom(l, 6, 3)
	//c.DeleteRandom(l, 6, 3)

	//c.DeleteRandom(l, 6, 3)

	//c.DeleteAllSpecific(l, 1)

	//fmt.Println(c.CountNodes(l))

	//fmt.Println(c.SearchList(l, 8))

	c.SortListAscending(l)

	c.ListReverse(l)

	c.PrintCircularList(l)

	//c.TestMerge()
	c.TestMergeSorted()
}
