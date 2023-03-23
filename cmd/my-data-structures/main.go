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

	n3 := linkedlists.Node{Value:192}
	ll.Append(&n3)

	// Print the linked list.
	// fmt.Println(ll)
	ll.PrintAll()
	
	// new node
	n4 := linkedlists.Node{Value: 66}

	// insert new node 
	ll.Insert(2,&n4)

	// search
	fmt.Println("---------")
	ll.PrintAll()
	
}