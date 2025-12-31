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

package array

import (
	"slices"

	"github.com/andrei-cosmin/sandata/bit"
	"github.com/andrei-cosmin/sandata/mathutil"
)

// Array a generic array implementation
//   - container []T - the array container
//   - empty T - the empty value for the array
type Array[T any] struct {
	container []T
	empty     T
}

// New creates a new array with the given size
func New[T any](size uint) *Array[T] {
	return &Array[T]{
		container: make([]T, size),
	}
}

// Get returns the value at the given index
func (a *Array[T]) Get(index uint) T {
	return a.container[index]
}

// Set sets the value at the given index (array will automatically grow if the index is out of bounds)
func (a *Array[T]) Set(index uint, value T) {
	a.ensureCapacity(index)
	a.container[index] = value
}

// Size returns the size of the array
func (a *Array[T]) Size() uint {
	return uint(len(a.container))
}

// ClearAll clears all the values in the array according to the set bits
func (a *Array[T]) ClearAll(bits bit.Mask) {
	for index, hasNext := bits.NextSet(0); hasNext && index < uint(len(a.container)); index, hasNext = bits.NextSet(index + 1) {
		a.container[index] = a.empty
	}
}

// ClearAllFunc clears all the values in the array according to the set bits
// and calls `f` for each cleared value
func (a *Array[T]) ClearAllFunc(bits bit.Mask, f func(T)) {
	for index, hasNext := bits.NextSet(0); hasNext && index < uint(len(a.container)); index, hasNext = bits.NextSet(index + 1) {
		f(a.container[index])
		a.container[index] = a.empty
	}
}

// Clear clears the value at the given index
func (a *Array[T]) Clear(index uint) {
	a.container[index] = a.empty
}

// ensureCapacity ensures the array has the capacity to store the value at the given index
func (a *Array[T]) ensureCapacity(index uint) {
	if index >= uint(len(a.container)) {
		a.container = slices.Grow(a.container, int(mathutil.NextPowerOfTwo(index+1))-len(a.container))
		a.container = a.container[:cap(a.container)]
	}
}
