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

package bit

import "github.com/bits-and-blooms/bitset"

// BitMask represents a read-only wrapper for operations of bitset.BitSet
//   - Bits *bitset.BitSet - the bitset to be handled
//
// NOTE: The bitset owned by the mask is considered a right hand operand in all operations
//
// WARNING: All bitwise operations with effects are destructive and will modify the given bitset (parameter - other)
//
//	BitMask.Difference(other) -> other = other & (~BitMask)
//	BitMask.SymmetricalDifference(other) -> other = other ^ BitMask

type bitSet = *bitset.BitSet
type BitMask struct {
	bitSet
}

// NewMask creates a new mask over the given bitset
func NewMask(bits *bitset.BitSet) *BitMask {
	return &BitMask{
		bitSet: bits,
	}
}

// Bits returns the underlying bitset
func (m *BitMask) Bits() *bitset.BitSet {
	return m.bitSet
}

// Compact shrinks BitSet to so that it preserves all set bits, while minimizing memory usage
func (m *BitMask) Compact() {
	m.bitSet = m.bitSet.Compact()
}

// Difference performs the difference operation with the given BitSet:
//
//	other = other & (~BitMask)
func (m *BitMask) Difference(other *bitset.BitSet) {
	other.InPlaceDifference(m.bitSet)
}

// DifferenceCardinality returns the cardinality of the difference operation with the given BitSet:
//
//	count( other & (~BitMask) )
func (m *BitMask) DifferenceCardinality(other *bitset.BitSet) uint {
	return other.DifferenceCardinality(m.bitSet)
}

// Intersection performs the intersection operation with the given BitSet:
//
//	other = other & BitMask
func (m *BitMask) Intersection(other *bitset.BitSet) {
	other.InPlaceIntersection(m.bitSet)
}

// IntersectionCardinality returns the cardinality of the intersection operation with the given BitSet:
//
//	count( other & BitMask )
func (m *BitMask) IntersectionCardinality(other *bitset.BitSet) uint {
	return other.IntersectionCardinality(m.bitSet)
}

// Union performs the union operation with the given BitSet:
//
//	other = other | BitMask
func (m *BitMask) Union(other *bitset.BitSet) {
	other.InPlaceUnion(m.bitSet)
}

// UnionCardinality returns the cardinality of the union operation with the given BitSet:
//
//	count( other | BitMask )
func (m *BitMask) UnionCardinality(other *bitset.BitSet) uint {
	return other.UnionCardinality(m.bitSet)
}

// SymmetricalDifference performs the symmetrical difference operation with the given BitSet:
//
//	other = other ^ BitMask
func (m *BitMask) SymmetricalDifference(other *bitset.BitSet) {
	other.InPlaceSymmetricDifference(m.bitSet)
}

// SymmetricalDifferenceCardinality returns the cardinality of the symmetrical difference operation
// with the given BitSet:
//
//	count( other ^ BitMask )
func (m *BitMask) SymmetricalDifferenceCardinality(other *bitset.BitSet) uint {
	return other.SymmetricDifferenceCardinality(m.bitSet)
}

// IsSubSetOf returns true if this is a subset of the other set
func (m *BitMask) IsSubSetOf(other *bitset.BitSet) bool {
	return other.IsSuperSet(m.bitSet)
}

// IsStrictSubSetOf returns true if this is a strict subset of the other set
func (m *BitMask) IsStrictSubSetOf(other *bitset.BitSet) bool {
	return other.IsStrictSuperSet(m.bitSet)
}
