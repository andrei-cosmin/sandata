package data

// ChainNode struct - represents a node in a double linked list
//   - PreviousChain *ChainNode[T] - reference to the previous node in the list
//   - NextChain *ChainNode[T] - reference to the next node in the list
//   - Data T - the data which is stored in the node
type ChainNode[T any] struct {
	PreviousChain *ChainNode[T]
	NextChain     *ChainNode[T]
	Data          T
}

// RemoveNode method - removes the node from the list, and links the previous and next nodes together
//
// Example (remove `B`):
//
//	A <-> B -> C  =>  A <-> C, B
func (c *ChainNode[T]) RemoveNode() {
	// Get the reference of the previous node
	previous := c.PreviousChain

	// Get the reference of the next node
	next := c.NextChain

	// If the previous node is not nil, link it to the next node
	if previous != nil {
		previous.NextChain = next
	}

	// If the next node is not nil, link it to the previous node
	if next != nil {
		next.PreviousChain = previous
	}

	// Isolate the current node
	c.PreviousChain = nil
	c.NextChain = nil
}

// InsertAsHead method - inserts the given node as the head of the list, and clears out references
//
// NOTE: the current node will become the head of the [other] node, the [other] node will become the tail of the current node
//
// Example (insert `B` as head of `A`):
//
//	List with `B`: X <-> B <-> Y <-> Z <-> T
//	List with `A`: C <-> A <-> D
//	Result:	X <-> B <-> A <-> D, Y <-> Z <-> T, C
func (c *ChainNode[T]) InsertAsHead(other *ChainNode[T]) {
	// Get the reference of the previous head
	previous := other.PreviousChain

	// Get reference of the next node
	next := c.NextChain

	// Disconnects the previous head from the [other] node
	if previous != nil {
		previous.NextChain = nil
	}

	// Disconnects the previous tail from the current node
	if next != nil {
		next.PreviousChain = nil
	}

	// Link the current node as head to the [other] node
	c.NextChain = other

	// Link the [other] node as tail to the current node
	other.PreviousChain = c
}

// Split method - splits the list at the current node
//
// NOTE: the previous head will become a tail, the current node will become the head of a new list
//
// Example (split at `B`):
//
//	X <-> Y <-> B <-> Z <-> T  =>  X <-> Y, B <-> Z <-> T
func (c *ChainNode[T]) Split() {
	// Get the reference of the previous node
	previous := c.PreviousChain

	// Set the previous node as the tail of the list
	if previous != nil {
		previous.NextChain = nil
	}

	// Set the current node as the head of the list
	c.PreviousChain = nil
}

// InsertAsTail method - inserts the given node as the tail of the list, and clears out references
//
// NOTE: the current node will become the tail of the [other] node, the [other] node will become the head of the current node
//
// Example (insert `A` as tail of `B`):
//
//	List with `A`: C <-> A <-> D
//	List with `B`: X <-> B <-> Y <-> Z <-> T
//	Result:	X <-> B <-> A <-> D, Y <-> Z <-> T, C
func (c *ChainNode[T]) InsertAsTail(other *ChainNode[T]) {
	// Get the reference of the previous head
	previous := c.PreviousChain

	// Get reference of the next node
	next := other.NextChain

	// Disconnects the previous head from the current node
	if previous != nil {
		previous.NextChain = nil
	}

	// Disconnects the previous tail from the [other] node
	if next != nil {
		next.PreviousChain = nil
	}

	// Link the current node as tail to the [other] node
	c.PreviousChain = other

	// Link the [other] node as head to the current node
	other.NextChain = c
}

// HasNext method - returns true if the current node has a next node
func (c *ChainNode[T]) HasNext() bool {
	return c.NextChain != nil
}

// HasPrevious method - returns true if the current node has a previous node
func (c *ChainNode[T]) HasPrevious() bool {
	return c.PreviousChain != nil
}

// IsHead method - returns true if the current node is the head of the list
func (c *ChainNode[T]) IsHead() bool {
	return c.PreviousChain == nil
}

// IsTail method - returns true if the current node is the tail of the list
func (c *ChainNode[T]) IsTail() bool {
	return c.NextChain == nil
}
