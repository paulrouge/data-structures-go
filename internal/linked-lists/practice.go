package linkedlists

import "fmt"

func Example() {

	// Create a new linked list.
	ll := LinkedList{}

	// Create nodes.
	n1 := Node{Value: 1}
	n2 := Node{Value: 21}
	n3 := Node{Value: 15}
	n4 := Node{Value: 307}
	n5 := Node{Value: 88}
	n6 := Node{Value: 36}

	// Append the nodes to the linked list.
	ll.Append(&n1)
	ll.Append(&n2)
	ll.Append(&n3)
	ll.Append(&n4)
	ll.Append(&n5)
	ll.Append(&n6)

	// Print the linked list.
	ll.PrintAll()

	// reversed ll
	err := ll.Reverse()
	if err != nil {
		fmt.Println(err)
	}

	ll.PrintAll()

}

func DoublyExample() {
	// create a new doubly linked list
	dll := DoublyLinkedList{}

	// create nodes
	n1 := DoublyNode{Value: 10}
	n2 := DoublyNode{Value: 4}
	n3 := DoublyNode{Value: 99}
	n4 := DoublyNode{Value: 129}
	n5 := DoublyNode{Value: 7}

	// append to doubly linked list
	dll.Append(&n1)
	dll.Append(&n2)
	dll.Append(&n4)

	// prepend
	dll.Prepend(&n3)

	dll.PrintAll()

	// insert
	dll.Insert(2, &n5)

	dll.PrintAll()

}
