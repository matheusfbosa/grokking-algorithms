package chapter2

import (
	"fmt"
)

// Node represents a node in the linked list.
type Node struct {
	Data int
	Next *Node
}

// LinkedList represents a singly linked list.
type LinkedList struct {
	Head *Node
}

// NewLinkedList creates a new linked list.
func NewLinkedList() *LinkedList {
	return &LinkedList{}
}

// Append adds a new node with the given value to the end of the linked list.
func (ll *LinkedList) Append(data int) {
	newNode := &Node{Data: data, Next: nil}

	if ll.Head == nil {
		ll.Head = newNode
		return
	}

	current := ll.Head
	for current.Next != nil {
		current = current.Next
	}

	current.Next = newNode
}

// Display prints the elements of the linked list.
func (ll LinkedList) Display() {
	current := ll.Head
	for current != nil {
		fmt.Printf("%d -> ", current.Data)
		current = current.Next
	}
	fmt.Println("nil")
}
