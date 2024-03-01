package data

import "github.com/bits-and-blooms/bitset"

// BitMask struct - represents a read-only wrapper for operations of bitset.BitSet
//   - Bits *bitset.BitSet - the bitset to be handled
//
// NOTE: The bitset owned by the mask is considered a right hand operand in all operations
//
// WARNING: All bitwise operations with effects are destructive and will modify the given bitset (parameter - other)
//
//	BitMask.Difference(other) -> other = other & (~BitMask)
//	BitMask.SymmetricalDifference(other) -> other = other ^ BitMask
type BitMask struct {
	Bits *bitset.BitSet
}

// NewMask method - creates a new mask over the given bitset
func NewMask(bits *bitset.BitSet) *BitMask {
	return &BitMask{
		Bits: bits,
	}
}

// Len method - returns the number of bits in the BitSet
func (m *BitMask) Len() uint {
	return uint(m.Bits.Len())
}

// Test method - returns whether ith bit is set
func (m *BitMask) Test(index uint) bool {
	return m.Bits.Test(index)
}

// Compact method - shrinks BitSet to so that we preserve all set bits, while minimizing memory usage
func (m *BitMask) Compact() {
	m.Bits = m.Bits.Compact()
}

// String method - return a string representation of the Bitmap
func (m *BitMask) String() string {
	return m.Bits.String()
}

// NextSet method returns the next bit set from the specified index, including possibly the current index
// along with an error code (true = valid, false = no set bit found, i.e all bits are clear)
//
//	for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
func (m *BitMask) NextSet(index uint) (uint, bool) {
	return m.Bits.NextSet(index)
}

// NextClear method returns the next bit clear from the specified index, including possibly the current index
// along with an error code (true = valid, false = no clear bit found, i.e all bits are set)
//
//	for i,e := v.NextClear(0); e; i,e = v.NextClear(i + 1) {...}
func (m *BitMask) NextClear(index uint) (uint, bool) {
	return m.Bits.NextClear(index)
}

// Clone method - returns a new BitSet with the same bits set and same size
func (m *BitMask) Clone() *bitset.BitSet {
	return m.Bits.Clone()
}

// Copy method - copies bits into a destination BitSet (using the Go array copy semantics).
//
// The number of bits copied is the minimum of the number of bits - min( Len(mask), Len(other) )
func (m *BitMask) Copy(other *bitset.BitSet) uint {
	return m.Bits.Copy(other)
}

// CopyFull method - copies into a destination BitSet such that the destination is
// identical to the source after the operation
func (m *BitMask) CopyFull(other *bitset.BitSet) *bitset.BitSet {
	m.Bits.CopyFull(other)
	return other
}

// Count method - returns the number of set bits
func (m *BitMask) Count() uint {
	return uint(m.Bits.Count())
}

// Equal method - returns whether the two BitSets are the same (compares both bits and size)
func (m *BitMask) Equal(other *bitset.BitSet) bool {
	return other.Equal(m.Bits)
}

// Difference method - performs the difference operation with the given BitSet:
//
//	other = other & (~BitMask)
func (m *BitMask) Difference(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceDifference(m.Bits)
	return other
}

// DifferenceCardinality method - returns the cardinality of the difference operation with the given BitSet:
//
//	count( other & (~BitMask) )
func (m *BitMask) DifferenceCardinality(other *bitset.BitSet) uint {
	return other.DifferenceCardinality(m.Bits)
}

// Intersection method - performs the intersection operation with the given BitSet:
//
//	other = other & BitMask
func (m *BitMask) Intersection(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceIntersection(m.Bits)
	return other
}

// IntersectionCardinality method - returns the cardinality of the intersection operation with the given BitSet:
//
//	count( other & BitMask )
func (m *BitMask) IntersectionCardinality(other *bitset.BitSet) uint {
	return other.IntersectionCardinality(m.Bits)
}

// Union method - performs the union operation with the given BitSet:
//
//	other = other | BitMask
func (m *BitMask) Union(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceUnion(m.Bits)
	return other
}

// UnionCardinality method - returns the cardinality of the union operation with the given BitSet:
//
//	count( other | BitMask )
func (m *BitMask) UnionCardinality(other *bitset.BitSet) uint {
	return other.UnionCardinality(m.Bits)
}

// SymmetricalDifference method - performs the symmetrical difference operation with the given BitSet:
//
//	other = other ^ BitMask
func (m *BitMask) SymmetricalDifference(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceSymmetricDifference(m.Bits)
	return other
}

// SymmetricalDifferenceCardinality method - returns the cardinality of the symmetrical difference operation
// with the given BitSet:
//
//	count( other ^ BitMask )
func (m *BitMask) SymmetricalDifferenceCardinality(other *bitset.BitSet) uint {
	return other.SymmetricDifferenceCardinality(m.Bits)
}

// All method - returns true if all bits are set, false otherwise (returns true for empty sets)
func (m *BitMask) All() bool {
	return m.Bits.All()
}

// None method - returns true if no bit is set, false otherwise (returns true for empty sets)
func (m *BitMask) None() bool {
	return m.Bits.None()
}

// Any method - returns true if any bit is set, false otherwise
func (m *BitMask) Any() bool {
	return m.Bits.Any()
}

// IsSuperSetOf method - returns true if this is a superset of the other set
func (m *BitMask) IsSuperSetOf(other *bitset.BitSet) bool {
	return m.Bits.IsSuperSet(other)
}

// IsStrictSuperSetOf method - returns true if this is a strict superset of the other set
func (m *BitMask) IsStrictSuperSetOf(other *bitset.BitSet) bool {
	return m.Bits.IsStrictSuperSet(other)
}

// IsSubSetOf method - returns true if this is a subset of the other set
func (m *BitMask) IsSubSetOf(other *bitset.BitSet) bool {
	return other.IsSuperSet(m.Bits)
}

// IsStrictSubSetOf method - returns true if this is a strict subset of the other set
func (m *BitMask) IsStrictSubSetOf(other *bitset.BitSet) bool {
	return other.IsStrictSuperSet(m.Bits)
}

// Rank method - returns the number of set Bits up to and including the index that are set in the bitset
func (m *BitMask) Rank(index uint) uint {
	return m.Bits.Rank(index)
}

// Select returns the index of the jth set bit, where j is the argument
//
// NOTE: The caller is responsible to ensure that 0 <= j < Count()
//
// WARNING: When j is out of range, the function returns the length of the bitset (b.length)
func (m *BitMask) Select(index uint) uint {
	return m.Bits.Select(index)
}
