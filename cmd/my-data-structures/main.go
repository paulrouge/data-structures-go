package main

import (
	"fmt"
	"data-structures/internal/linked-lists"
)

func main() {
	fmt.Println("Main ran...")

	// Create a new linked list.
	ll := linkedlists.LinkedList{}

	// Create a new node.
	n := linkedlists.Node{Value: 1}

	// Append the node to the linked list.
	ll.Append(&n)

	// Create a new node.
	n2 := linkedlists.Node{Value: 21}

	// Prepend the node to the linked list.
	ll.Prepend(&n2)

	// Print the linked list.
	fmt.Println(ll.Length)

	// Find a node in the linked list.
	fmt.Println(ll.Find(100))

}