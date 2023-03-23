package linkedlists

// Append adds a node to the end of the linked list.
func (l *LinkedList) Append(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		l.Tail.Next = n
		l.Tail = n
	}
	l.Length++
}

// Prepend adds a node to the beginning of the linked list.
func (l *LinkedList) Prepend(n *Node) {
	if l.Head == nil {
		l.Head = n
		l.Tail = n
	} else {
		n.Next = l.Head
		l.Head = n
	}
	l.Length++
}

