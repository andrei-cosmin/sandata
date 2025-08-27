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

type Mask interface {

	// Len - returns the number of bits in the BitSet
	Len() uint

	// Test - returns whether ith bit is set
	Test(index uint) bool

	// String - returns a string representation of the Bitmap
	String() string

	// NextSet - returns the next bit set from the specified index, including possibly the current index
	// along with an error code (true = valid, false = no set bit found, i.e all bits are clear)
	//
	//	for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
	NextSet(index uint) (uint, bool)

	// NextClear - returns the next bit clear from the specified index, including possibly the current index
	// along with an error code (true = valid, false = no clear bit found, i.e all bits are set)
	//
	//	for i,e := v.NextClear(0); e; i,e = v.NextClear(i + 1) {...}
	NextClear(index uint) (uint, bool)

	// Clone - returns a new BitSet with the same bits set and same size
	Clone() *bitset.BitSet

	// Copy - copies bits into a destination BitSet (using the Go array copy semantics).
	//
	// The number of bits copied is the minimum of the number of bits - min( Len(mask), Len(other) )
	Copy(other *bitset.BitSet) uint

	// CopyFull - copies into a destination BitSet such that the destination is
	// identical to the source after the operation
	CopyFull(other *bitset.BitSet)

	// Count - returns the number of set bits
	Count() uint

	// Equal - returns whether the two BitSets are the same (compares both bits and size)
	Equal(other *bitset.BitSet) bool

	// Difference - performs the difference operation with the given BitSet:
	//
	//	other = other & (~BitMask)
	Difference(other *bitset.BitSet)

	// DifferenceCardinality - returns the cardinality of the difference operation with the given BitSet:
	//
	//	count( other & (~BitMask) )
	DifferenceCardinality(other *bitset.BitSet) uint

	// Intersection - performs the intersection operation with the given BitSet:
	//
	//	other = other & BitMask
	Intersection(other *bitset.BitSet)

	// IntersectionCardinality - returns the cardinality of the intersection operation with the given BitSet:
	//
	//	count( other & BitMask )
	IntersectionCardinality(other *bitset.BitSet) uint

	// Union - performs the union operation with the given BitSet:
	//
	//	other = other | BitMask
	Union(other *bitset.BitSet)

	// UnionCardinality - returns the cardinality of the union operation with the given BitSet:
	//
	//	count( other | BitMask )
	UnionCardinality(other *bitset.BitSet) uint

	// SymmetricalDifference - performs the symmetrical difference operation with the given BitSet:
	//
	//	other = other ^ BitMask
	SymmetricalDifference(other *bitset.BitSet)

	// SymmetricalDifferenceCardinality - returns the cardinality of the symmetrical difference operation
	// with the given BitSet:
	//
	//	count( other ^ BitMask )
	SymmetricalDifferenceCardinality(other *bitset.BitSet) uint

	// All - returns true if all bits are set, false otherwise (returns true for empty sets)
	All() bool

	// None - returns true if no bit is set, false otherwise (returns true for empty sets)
	None() bool

	// Any - returns true if any bit is set, false otherwise
	Any() bool

	// IsSuperSet - returns true if this is a superset of the other set
	IsSuperSet(other *bitset.BitSet) bool

	// IsStrictSuperSet - returns true if this is a strict superset of the other set
	IsStrictSuperSet(other *bitset.BitSet) bool

	// IsSubSetOf - returns true if this is a subset of the other set
	IsSubSetOf(other *bitset.BitSet) bool

	// IsStrictSubSetOf - returns true if this is a strict subset of the other set
	IsStrictSubSetOf(other *bitset.BitSet) bool

	// Rank - returns the number of set bits up to and including the index that are set in the bitset
	Rank(index uint) uint

	// Select returns the index of the jth set bit, where j is the argument
	//
	// NOTE: The caller is responsible to ensure that 0 <= j < Count()
	//
	// WARNING: When j is out of range, the function returns the length of the bitset (b.length)
	Select(index uint) uint
}
