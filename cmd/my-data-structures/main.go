package main

import (
	"fmt"
	"data-structures/internal/linked-lists"
)

func main() {
	fmt.Println("Main ran...")

	// Create a new linked list.
	ll := linkedlists.LinkedList{}

	// Create nodes.
	n1 := linkedlists.Node{Value: 1}
	n2 := linkedlists.Node{Value: 21}
	n3 := linkedlists.Node{Value: 15}
	n4 := linkedlists.Node{Value: 307}
	n5 := linkedlists.Node{Value: 88}
	n6 := linkedlists.Node{Value: 36}
	
	
	// Append the nodes to the linked list.
	ll.Append(&n1)
	ll.Append(&n2)
	ll.Append(&n3)
	ll.Append(&n4)
	ll.Append(&n5)
	ll.Append(&n6)


	// Print the linked list.
	ll.PrintAll()

	// remove index 4
	ll.RemoveIndex(2)

	ll.PrintAll()
}