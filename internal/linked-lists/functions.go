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

// Find returns the first node with the given value.
func (l *LinkedList) Find(v int) *Node {
	current := l.Head
	for current != nil {
		if current.Value == v {
			return current
		} else {
			current = current.Next
		}
	}
	return nil
}

// Delete removes the first node with the given value.
func (l *LinkedList) Delete(v int) {
	current := l.Head
	for current != nil {
		if current.Value == v {
			current.Value = current.Next.Value
			current.Next = current.Next.Next
			l.Length--
		}
	}

}
