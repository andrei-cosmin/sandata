/*
 * MIT License
 *
 * Copyright (c) 2025 Andrei Casu-Pop
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software and associated
 * documentation files (the "Software"), to deal in the Software without restriction, including without limitation the
 * rights to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies of the Software, and to
 * permit persons to whom the Software is furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial portions of the
 * Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE
 * WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR
 * COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR
 * OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package pool

// Pool representation of a pool
//   - cursor int - current cursor position
//   - container []T - container for the pool
//   - empty T - empty value for the pool
type Pool[T any] struct {
	cursor    int
	container []T
	empty     T
}

// New creates a new pool with the given capacity
func New[T any](capacity uint) *Pool[T] {
	return &Pool[T]{
		cursor:    -1,
		container: make([]T, capacity),
	}
}

// Push pushes a value to the pool
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

// Pop pops a value from the pool
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

// Size returns the size of the pool (current fill of the pool)
func (p *Pool[T]) Size() int {
	return p.cursor + 1
}

// Capacity returns the capacity of the pool
func (p *Pool[T]) Capacity() int {
	return len(p.container)
}

// Empty returns true if the pool is empty, false otherwise
func (p *Pool[T]) Empty() bool {
	return p.cursor == -1
}
