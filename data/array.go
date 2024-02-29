package data

import (
	"github.com/andrei-cosmin/sandata/internal/util"
	"github.com/bits-and-blooms/bitset"
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
func (a *Array[T]) ClearAll(set *bitset.BitSet) {
	for index, hasNext := set.NextSet(0); hasNext; index, hasNext = set.NextSet(index + 1) {
		if index >= uint(len(a.container)) {
			return
		}
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
		a.container = slices.Grow(a.container, int(util.NextPowerOfTwo(index+1))-len(a.container))
		a.container = a.container[:cap(a.container)]
	}
}
