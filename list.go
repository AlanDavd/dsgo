/* Linked list. */
package dsgo

/* List container. */
type List struct {
	Head   *Node
	Length int
}

/* Linked list node. */
type Node struct {
	Value  int
	Next *Node
}

/* Initialize linked list. */
func New() List {
	return List{}
}

/*
 * Get the value of the index-th node in the linked list.
 * Return -1 if the index is invalid.
 * List is zero-based.
 */
func (l *List) Get(index int) int {
	if index < 0 || index > l.Length-1 {
		return -1
	}

	currentNode := l.Head
	for i := 0; i < index && currentNode != nil; i++ {
		currentNode = currentNode.Next
	}

	if currentNode == nil {
		return -1
	}

	return currentNode.Value
}

/*
 * Add a node before the first element of the linked list.
 * After the insertion, the new node will be the first node of the linked list.
 */
func (l *List) AddAtHead(val int) {
	var newNode Node
	newNode.Value = val
	newNode.Next = l.Head
	l.Head = &newNode
	l.Length++
}

/*
 * Append a node to the last element of the linked list.
 */
func (l *List) AddAtTail(val int) {
	if l.Head == nil {
		l.AddAtHead(val)
		return
	}

	currentNode := l.Head
	for i := 0; currentNode.Next != nil; i++ {
		currentNode = currentNode.Next
	}

	var newNode Node
	newNode.Value = val
	currentNode.Next = &newNode

	l.Length++
}

/*
 * Add a node before the index-th node in the linked list.
 * If index equals to the length of linked list, the node will be appended
 * to the end of linked list. If index is greater than the length, the node
 * will not be inserted.
 */
func (l *List) AddAtIndex(index int, val int) {
	if index == 0 {
		l.AddAtHead(val)
		return
	} else if index == l.Length {
		l.AddAtTail(val)
		return
	} else if index < 0 || index > l.Length {
		return
	}

	currentNode := l.Head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	var newNode Node
	newNode.Value = val
	newNode.Next = currentNode.Next
	currentNode.Next = &newNode

	l.Length++
}

/*
 * Delete the index-th node in the linked list, if the index is valid.
 */
func (l *List) DeleteAtIndex(index int) {
	if index < 0 || index >= l.Length {
		return
	} else if index == 0 {
		l.Head = l.Head.Next
		l.Length--
		return
	}

	currentNode := l.Head
	for i := 0; i < index-1; i++ {
		currentNode = currentNode.Next
	}

	currentNode.Next = currentNode.Next.Next
	l.Length--
}
