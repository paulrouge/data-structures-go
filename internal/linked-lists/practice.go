package linkedlists

func RunExample() {

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

	// remove index 4
	ll.RemoveIndex(2)

	ll.PrintAll()
}