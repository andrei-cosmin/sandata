package array

import (
	"github.com/andrei-cosmin/sandata/bit"
	"github.com/andrei-cosmin/sandata/mathutil"
	"slices"
)

// Array struct - a generic array implementation
//   - container []T - the array container
//   - empty T - the empty value for the array
type Array[T any] struct {
	container []T
	empty     T
}

// NewArray method - creates a new array with the given size
func NewArray[T any](size uint) *Array[T] {
	return &Array[T]{
		container: make([]T, size),
	}
}

// Get method - returns the value at the given index
func (a *Array[T]) Get(index uint) T {
	return a.container[index]
}

// Set method - sets the value at the given index (array will automatically grow if the index is out of bounds)
func (a *Array[T]) Set(index uint, value T) {
	a.ensureCapacity(index)
	a.container[index] = value
}

// Size method - returns the size of the array
func (a *Array[T]) Size() uint {
	return uint(len(a.container))
}

// ClearAll method - clears all the values in the array according to the set bits
func (a *Array[T]) ClearAll(bits bit.Mask) {
	for index, hasNext := bits.NextSet(0); hasNext && index < uint(len(a.container)); index, hasNext = bits.NextSet(index + 1) {
		a.container[index] = a.empty
	}
}

// ClearAllFunc method - clears all the values in the array according to the set bits
// and calls `f` for each cleared value
func (a *Array[T]) ClearAllFunc(bits bit.Mask, f func(T)) {
	for index, hasNext := bits.NextSet(0); hasNext && index < uint(len(a.container)); index, hasNext = bits.NextSet(index + 1) {
		f(a.container[index])
		a.container[index] = a.empty
	}
}

// Clear method - clears the value at the given index
func (a *Array[T]) Clear(index uint) {
	a.container[index] = a.empty
}

// ensureCapacity - ensures the array has the capacity to store the value at the given index
func (a *Array[T]) ensureCapacity(index uint) {
	if index >= uint(len(a.container)) {
		a.container = slices.Grow(a.container, int(mathutil.NextPowerOfTwo(index+1))-len(a.container))
		a.container = a.container[:cap(a.container)]
	}
}
