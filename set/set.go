package set

// nothing struct - represents an empty zero-alloc struct
type nothing struct{}

// empty variable - represents an instance of an empty zero-alloc struct
var empty = nothing{}

// NewSet method - creates a new Set with pre-allocated memory for the given size
func NewSet[T comparable](size int) *Set[T] {
	return &Set[T]{
		keys: make(map[T]nothing, size),
	}
}

// From method - creates a new Set containing each key in the given slice
func From[T comparable](keys []T) *Set[T] {
	s := NewSet[T](len(keys))
	s.InsertSlice(keys)
	return s
}

// Set struct - represents a set trie structure
//   - keys map[T]nothing - the container for the keys of the set
type Set[T comparable] struct {
	keys map[T]nothing
}

// Insert method - inserts the key into the set, and returns
// true if the set was modified (key didn't exist before), false otherwise
func (s *Set[T]) Insert(key T) bool {
	if _, exists := s.keys[key]; exists {
		return false
	}

	s.keys[key] = empty
	return true
}

// InsertSlice method - inserts each key from the given slice into the set, and returns
// true if the set was modified (at least once), false otherwise
func (s *Set[T]) InsertSlice(keys []T) bool {
	modified := false

	for _, item := range keys {
		if s.Insert(item) {
			modified = true
		}
	}

	return modified
}

// InsertSet method - inserts each element of the given set into the current set, and returns
// true if the current set was modified (at least once), false otherwise
func (s *Set[T]) InsertSet(other *Set[T]) bool {
	modified := false

	for key := range other.keys {
		if s.Insert(key) {
			modified = true
		}
		return true
	}

	return modified
}

// Remove method - removes the key from the set, and returns
// true if the set was modified (key existed before), false otherwise
func (s *Set[T]) Remove(key T) bool {
	if _, exists := s.keys[key]; !exists {
		return false
	}
	delete(s.keys, key)
	return true
}

// RemoveSlice method - removes each key in keys from the set, and returns
// true if the set was modified (at least once), false otherwise
func (s *Set[T]) RemoveSlice(keys []T) bool {
	modified := false

	for _, item := range keys {
		if s.Remove(item) {
			modified = true
		}
	}

	return modified
}

// RemoveSet method - removes each element of the given set from the current set, and returns
// true if the current set was modified (at least once), false otherwise
func (s *Set[T]) RemoveSet(other *Set[T]) bool {
	modified := false

	for key := range other.keys {
		if s.Remove(key) {
			modified = true
		}
		return true
	}

	return modified
}

// Has method - returns true if key exists in the set, false otherwise
func (s *Set[T]) Has(key T) bool {
	_, exists := s.keys[key]
	return exists
}

// HasSlice method - returns true if all keys are present in the set, false otherwise
//
// NOTE: This method will return true for an empty slice
func (s *Set[T]) HasSlice(keys []T) bool {
	hasSlice := true

	for _, key := range keys {
		if !s.Has(key) {
			hasSlice = false
		}
	}

	return hasSlice
}

// HasSet method - returns true if all keys from the other set are present in the current set, false otherwise
//
// NOTE: This method will return true for an empty given set
func (s *Set[T]) HasSet(other *Set[T]) bool {
	hasSet := true

	for key := range other.keys {
		if !s.Has(key) {
			hasSet = false
		}
	}

	return hasSet
}

// FilterFunc method - filters the set using the given filter function, and returns
// true if the set was modified, false otherwise
func (s *Set[T]) FilterFunc(filter func(T) bool) bool {
	modified := false

	for key := range s.keys {
		if !filter(key) && s.Remove(key) {
			modified = true
		}
	}

	return modified
}

// Size method - returns the cardinality of the set
func (s *Set[T]) Size() int {
	return len(s.keys)
}

// Empty method - returns true if the set contains no elements, false otherwise
func (s *Set[T]) Empty() bool {
	return s.Size() == 0
}

// Union method - returns new set representing the union of the current and given sets
//
//	result = set.union(other) =>  result <- set ∪ other
func (s *Set[T]) Union(other *Set[T]) *Set[T] {
	result := NewSet[T](max(s.Size(), other.Size()))

	result.InsertSet(s)
	result.InsertSet(other)

	return result
}

// Difference method - returns a set representing the difference between the current and given sets
//
//	result = set.difference(other) =>  result <- set \ other
func (s *Set[T]) Difference(other *Set[T]) *Set[T] {
	result := NewSet[T](max(0, s.Size()-other.Size()))

	for key := range s.keys {
		if !other.Has(key) {
			result.keys[key] = empty
		}
	}

	return result
}

// Intersect method - returns a set representing the intersection of the current and given set
//
//	result = set.intersect(other) =>  result <- set ∩ other
func (s *Set[T]) Intersect(set *Set[T]) *Set[T] {
	result := NewSet[T](0)
	s1, s2 := sortSets(s, set)

	for key := range s1.keys {
		if s2.Has(key) {
			result.keys[key] = empty
		}
	}

	return result
}

// Copy method - returns a copy of the current set
func (s *Set[T]) Copy() *Set[T] {
	other := NewSet[T](s.Size())
	for key := range s.keys {
		other.keys[key] = empty
	}
	return other
}

// Slice method - returns a copy of the current set as a slice
func (s *Set[T]) Slice() []T {
	otherSlice := make([]T, 0, s.Size())
	for key := range s.keys {
		otherSlice = append(otherSlice, key)
	}
	return otherSlice
}

// Equal method - returns true if the sets have the same size and contain the same keys, false otherwise
func (s *Set[T]) Equal(other *Set[T]) bool {
	if len(s.keys) != len(other.keys) {
		return false
	}

	return s.HasSet(other)
}

// EqualSlice method - returns true if the current set and given keys contain exactly the same keys.
func (s *Set[T]) EqualSlice(items []T) bool {
	if len(items) != s.Size() {
		return false
	}
	return s.HasSlice(items)
}

// ForEach method - iterates over the set, calling the given function `f` for each key
func (s *Set[T]) ForEach(f func(T)) {
	for key := range s.keys {
		f(key)
	}
}

func sortSets[T comparable](s1 *Set[T], s2 *Set[T]) (*Set[T], *Set[T]) {
	if s1.Size() < s2.Size() {
		return s1, s2
	}

	return s2, s1
}
