package data

import "github.com/bits-and-blooms/bitset"

type Mask interface {

	// Len method - returns the number of bits in the BitSet
	Len() uint

	// Test method - returns whether ith bit is set
	Test(index uint) bool

	// Compact method - shrinks BitSet to so that we preserve all set bits, while minimizing memory usage
	Compact()

	// String method - return a string representation of the Bitmap
	String() string

	// NextSet method returns the next bit set from the specified index, including possibly the current index
	// along with an error code (true = valid, false = no set bit found, i.e all bits are clear)
	//
	//	for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
	NextSet(index uint) (uint, bool)

	// NextClear method returns the next bit clear from the specified index, including possibly the current index
	// along with an error code (true = valid, false = no clear bit found, i.e all bits are set)
	//
	//	for i,e := v.NextClear(0); e; i,e = v.NextClear(i + 1) {...}
	NextClear(index uint) (uint, bool)

	// Clone method - returns a new BitSet with the same bits set and same size
	Clone() *bitset.BitSet

	// Copy method - copies bits into a destination BitSet (using the Go array copy semantics).
	//
	// The number of bits copied is the minimum of the number of bits - min( Len(mask), Len(other) )
	Copy(other *bitset.BitSet) uint

	// CopyFull method - copies into a destination BitSet such that the destination is
	// identical to the source after the operation
	CopyFull(other *bitset.BitSet) *bitset.BitSet

	// Count method - returns the number of set bits
	Count() uint

	// Equal method - returns whether the two BitSets are the same (compares both bits and size)
	Equal(other *bitset.BitSet) bool

	// Difference method - performs the difference operation with the given BitSet:
	//
	//	other = other & (~BitMask)
	Difference(other *bitset.BitSet) *bitset.BitSet

	// DifferenceCardinality method - returns the cardinality of the difference operation with the given BitSet:
	//
	//	count( other & (~BitMask) )
	DifferenceCardinality(other *bitset.BitSet) uint

	// Intersection method - performs the intersection operation with the given BitSet:
	//
	//	other = other & BitMask
	Intersection(other *bitset.BitSet) *bitset.BitSet

	// IntersectionCardinality method - returns the cardinality of the intersection operation with the given BitSet:
	//
	//	count( other & BitMask )
	IntersectionCardinality(other *bitset.BitSet) uint

	// Union method - performs the union operation with the given BitSet:
	//
	//	other = other | BitMask
	Union(other *bitset.BitSet) *bitset.BitSet

	// UnionCardinality method - returns the cardinality of the union operation with the given BitSet:
	//
	//	count( other | BitMask )
	UnionCardinality(other *bitset.BitSet) uint

	// SymmetricalDifference method - performs the symmetrical difference operation with the given BitSet:
	//
	//	other = other ^ BitMask
	SymmetricalDifference(other *bitset.BitSet) *bitset.BitSet

	// SymmetricalDifferenceCardinality method - returns the cardinality of the symmetrical difference operation
	// with the given BitSet:
	//
	//	count( other ^ BitMask )
	SymmetricalDifferenceCardinality(other *bitset.BitSet) uint

	// All method - returns true if all bits are set, false otherwise (returns true for empty sets)
	All() bool

	// None method - returns true if no bit is set, false otherwise (returns true for empty sets)
	None() bool

	// Any method - returns true if any bit is set, false otherwise
	Any() bool

	// IsSuperSetOf method - returns true if this is a superset of the other set
	IsSuperSetOf(other *bitset.BitSet) bool

	// IsStrictSuperSetOf method - returns true if this is a strict superset of the other set
	IsStrictSuperSetOf(other *bitset.BitSet) bool

	// IsSubSetOf method - returns true if this is a subset of the other set
	IsSubSetOf(other *bitset.BitSet) bool

	// IsStrictSubSetOf method - returns true if this is a strict subset of the other set
	IsStrictSubSetOf(other *bitset.BitSet) bool

	// Rank method - returns the number of set bits up to and including the index that are set in the bitset
	Rank(index uint) uint

	// Select returns the index of the jth set bit, where j is the argument
	//
	// NOTE: The caller is responsible to ensure that 0 <= j < Count()
	//
	// WARNING: When j is out of range, the function returns the length of the bitset (b.length)
	Select(index uint) uint
}
