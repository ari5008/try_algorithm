package try

import "fmt"

type Node2 struct {
	prev *Node2
	data int
	next *Node2
}

type DoubleLinkedList struct {
	head *Node2
	tail *Node2
}

func (list *DoubleLinkedList) Insert(data int) {
	newNode := &Node2{
		prev: nil,
		data: data,
		next: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		currentNode := list.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		newNode.prev = currentNode
		currentNode.next = newNode
		list.tail = newNode
	}
}

func (list *DoubleLinkedList) Append(data int) {
	newNode := &Node2{
		prev: nil,
		data: data,
		next: nil,
	}

	if list.head == nil {
		list.head = newNode
		list.tail = newNode
	} else {
		newNode.prev = list.tail
		list.tail.next = newNode
		list.tail = newNode
	}
}

func (list *DoubleLinkedList) Remove(data int) bool {
	if list.head == nil {
		return false
	}

	currentNode := list.head
	for currentNode != nil {
		if currentNode.data == data {
			if currentNode == list.head {
				list.head = currentNode.next
				if list.head != nil {
					list.head.prev = nil
				}
				if currentNode == list.tail {
					list.tail = nil
				}
			} else if currentNode == list.tail {
				list.tail = currentNode.prev
				list.tail.next = nil
			} else {
				currentNode.prev.next = currentNode.next
				currentNode.next.prev = currentNode.prev
			}
			return true
		}
		currentNode = currentNode.next
	}

	return false
}

func (list *DoubleLinkedList) Display() {
	currentNode := list.head
	for currentNode != nil {
		fmt.Println(currentNode.data)
		currentNode = currentNode.next
	}
}
