package data

// Pool - struct for a pool
//   - cursor int - current cursor position
//   - container []T - container for the pool
//   - empty T - empty value for the pool
type Pool[T any] struct {
	cursor    int
	container []T
	empty     T
}

// NewPool method - creates a new pool with the given size
func NewPool[T any](size uint) *Pool[T] {
	return &Pool[T]{
		cursor:    -1,
		container: make([]T, size),
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
