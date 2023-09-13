package try

import "fmt"

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

func (ll *LinkedList) Filter() {
	current := ll.Head
	for current != nil {
		if current.Data%2 != 0 {
			fmt.Printf("%d ", current.Data)
		}
		current = current.Next
	}
}

// func (ll *LinkedList) Reverse() {
// 	current := ll.Head
// 	for current == nil{
// 		current = current.Next
// 	}
// 	for current 
// }


func (ll *LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d ", current.Data)
		current = current.Next
	}
	fmt.Println()
}
