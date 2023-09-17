package try

import (
	"fmt"

)

type Node struct {
	Data int
	Next *Node
}
type LinkedList struct {
	Head *Node
}

func (ll *LinkedList) Insert(data int) {
	newNode := &Node{
		Data: data,
		Next: nil,
	}

	if ll.Head == nil {
		ll.Head = newNode
	} else {
		current := ll.Head
		for current.Next != nil {
			current = current.Next
		}
		current.Next = newNode
	}
}

// func (ll *LinkedList) ReverseIterative() {
// 	var prev *Node 
// 	current := ll.Head

// 	for current != nil {
// 			next := current.Next
// 			current.Next = prev
// 			prev = current
// 			current = next
// 	}
// 	ll.Head = prev
// }

// func (ll *LinkedList) ReverseRecursive() {
// 	ll.Head = ll.reverse(ll.Head, nil)
// }

// func (ll *LinkedList) reverse(current, prev *Node) *Node {
// 	if current == nil {
// 			return prev
// 	}
	
// 	next := current.Next
// 	current.Next = prev
	
// 	return ll.reverse(next, current)
// }

func (ll *LinkedList) ReverseEven() {
	var prev *Node
	current := ll.Head

	for current != nil {
		if current.Data % 2 == 0 {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
			
		} else {
			if prev == nil {
				current = current.Next
			}
		}
	}
	// ll.Head = prev
}

func (ll *LinkedList) Filter() {
	current := ll.Head
	for current != nil {
		if current.Data%2 != 0 {
			fmt.Printf("%d ", current.Data)
		}
		current = current.Next
	}
}

func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
