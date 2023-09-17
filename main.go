package main

import (
	try "backend/try/list"
)

func main() {
	ll := new(try.LinkedList)

	ll.Insert(1)
	ll.Insert(4)
	ll.Insert(6)
	ll.Insert(8)
	ll.Insert(9)
	ll.Insert(10)
	ll.Insert(6)
	ll.Insert(3)
	ll.Insert(2)
	ll.Insert(3)
	// ll.ReverseIterative()
	// ll.ReverseRecursive()
	ll.ReverseEven()
	ll.Display()

}


