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

func (ll *LinkedList) ReverseIterative() {
	var prev *Node
	current := ll.Head

	for current != nil {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
	}
	ll.Head = prev
}

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
	ll.Head = reverseEven(ll.Head, nil)
}

func reverseEven(head *Node, prev *Node) *Node {
	if head == nil {
			return nil
	}
	
	current := head
	for current != nil && current.Data%2 == 0 {
			next := current.Next
			current.Next = prev
			prev = current
			current = next
	}

	if current != head {
			head.Next = current
			reverseEven(current, nil)
			return prev
	} else {
			head.Next = reverseEven(head.Next, head)
			return head
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
