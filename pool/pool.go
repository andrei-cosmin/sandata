package pool

// Pool - struct for a pool
//   - cursor int - current cursor position
//   - container []T - container for the pool
//   - empty T - empty value for the pool
type Pool[T any] struct {
	cursor    int
	container []T
	empty     T
}

// New method - creates a new pool with the given capacity
func New[T any](capacity uint) *Pool[T] {
	return &Pool[T]{
		cursor:    -1,
		container: make([]T, capacity),
	}
}

// Push method - pushes a value to the pool
func (p *Pool[T]) Push(value T) {
	// If the pool is full, return
	if p.cursor+1 == len(p.container) {
		return
	} else {
		// Increment the cursor and push the value
		p.cursor++
		p.container[p.cursor] = value
	}
}

// Pop method - pops a value from the pool
func (p *Pool[T]) Pop() (T, bool) {
	if p.cursor == -1 {
		// If the pool is empty, return the (empty value, false)
		return p.empty, false
	} else {
		// Decrement the cursor and pop the value
		value := p.container[p.cursor]
		p.container[p.cursor] = p.empty
		p.cursor--

		// Return the value and true
		return value, true
	}
}

// Size method - returns the size of the pool (current fill of the pool)
func (p *Pool[T]) Size() int {
	return p.cursor + 1
}

// Capacity method - returns the capacity of the pool
func (p *Pool[T]) Capacity() int {
	return len(p.container)
}

// Empty method - returns true if the pool is empty, false otherwise
func (p *Pool[T]) Empty() bool {
	return p.cursor == -1
}
