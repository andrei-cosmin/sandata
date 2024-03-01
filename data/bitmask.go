package data

import "github.com/bits-and-blooms/bitset"

// BitMask struct - represents a read-only wrapper for operations of bitset.BitSet
//   - bits *bitset.BitSet - the bitset to be handled
//
// NOTE: The bitset owned by the mask is considered a right hand operand in all operations
//
// WARNING: All bitwise operations with effects are destructive and will modify the given bitset (parameter - other)
//
//	BitMask.Difference(other) -> other = other & (~BitMask)
//	BitMask.SymmetricalDifference(other) -> other = other ^ BitMask
type BitMask struct {
	bits *bitset.BitSet
}

// NewMask method - creates a new mask over the given bitset
func NewMask(bits *bitset.BitSet) *BitMask {
	return &BitMask{
		bits: bits,
	}
}

// Len method - returns the number of bits in the BitSet
func (m *BitMask) Len() uint {
	return uint(m.bits.Len())
}

// Test method - returns whether ith bit is set
func (m *BitMask) Test(index uint) bool {
	return m.bits.Test(index)
}

// Compact method - shrinks BitSet to so that we preserve all set bits, while minimizing memory usage
func (m *BitMask) Compact() {
	m.bits = m.bits.Compact()
}

// String method - return a string representation of the Bitmap
func (m *BitMask) String() string {
	return m.bits.String()
}

// NextSet method returns the next bit set from the specified index, including possibly the current index
// along with an error code (true = valid, false = no set bit found, i.e all bits are clear)
//
//	for i,e := v.NextSet(0); e; i,e = v.NextSet(i + 1) {...}
func (m *BitMask) NextSet(index uint) (uint, bool) {
	return m.bits.NextSet(index)
}

// NextClear method returns the next bit clear from the specified index, including possibly the current index
// along with an error code (true = valid, false = no clear bit found, i.e all bits are set)
//
//	for i,e := v.NextClear(0); e; i,e = v.NextClear(i + 1) {...}
func (m *BitMask) NextClear(index uint) (uint, bool) {
	return m.bits.NextClear(index)
}

// Clone method - returns a new BitSet with the same bits set and same size
func (m *BitMask) Clone() *bitset.BitSet {
	return m.bits.Clone()
}

// Copy method - copies bits into a destination BitSet (using the Go array copy semantics).
//
// The number of bits copied is the minimum of the number of bits - min( Len(mask), Len(other) )
func (m *BitMask) Copy(other *bitset.BitSet) uint {
	return m.bits.Copy(other)
}

// CopyFull method - copies into a destination BitSet such that the destination is
// identical to the source after the operation
func (m *BitMask) CopyFull(other *bitset.BitSet) *bitset.BitSet {
	m.bits.CopyFull(other)
	return other
}

// Count method - returns the number of set bits
func (m *BitMask) Count() uint {
	return uint(m.bits.Count())
}

// Equal method - returns whether the two BitSets are the same (compares both bits and size)
func (m *BitMask) Equal(other *bitset.BitSet) bool {
	return other.Equal(m.bits)
}

// Difference method - performs the difference operation with the given BitSet:
//
//	other = other & (~BitMask)
func (m *BitMask) Difference(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceDifference(m.bits)
	return other
}

// DifferenceCardinality method - returns the cardinality of the difference operation with the given BitSet:
//
//	count( other & (~BitMask) )
func (m *BitMask) DifferenceCardinality(other *bitset.BitSet) uint {
	return other.DifferenceCardinality(m.bits)
}

// Intersection method - performs the intersection operation with the given BitSet:
//
//	other = other & BitMask
func (m *BitMask) Intersection(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceIntersection(m.bits)
	return other
}

// IntersectionCardinality method - returns the cardinality of the intersection operation with the given BitSet:
//
//	count( other & BitMask )
func (m *BitMask) IntersectionCardinality(other *bitset.BitSet) uint {
	return other.IntersectionCardinality(m.bits)
}

// Union method - performs the union operation with the given BitSet:
//
//	other = other | BitMask
func (m *BitMask) Union(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceUnion(m.bits)
	return other
}

// UnionCardinality method - returns the cardinality of the union operation with the given BitSet:
//
//	count( other | BitMask )
func (m *BitMask) UnionCardinality(other *bitset.BitSet) uint {
	return other.UnionCardinality(m.bits)
}

// SymmetricalDifference method - performs the symmetrical difference operation with the given BitSet:
//
//	other = other ^ BitMask
func (m *BitMask) SymmetricalDifference(other *bitset.BitSet) *bitset.BitSet {
	other.InPlaceSymmetricDifference(m.bits)
	return other
}

// SymmetricalDifferenceCardinality method - returns the cardinality of the symmetrical difference operation
// with the given BitSet:
//
//	count( other ^ BitMask )
func (m *BitMask) SymmetricalDifferenceCardinality(other *bitset.BitSet) uint {
	return other.SymmetricDifferenceCardinality(m.bits)
}

// All method - returns true if all bits are set, false otherwise (returns true for empty sets)
func (m *BitMask) All() bool {
	return m.bits.All()
}

// None method - returns true if no bit is set, false otherwise (returns true for empty sets)
func (m *BitMask) None() bool {
	return m.bits.None()
}

// Any method - returns true if any bit is set, false otherwise
func (m *BitMask) Any() bool {
	return m.bits.Any()
}

// IsSuperSetOf method - returns true if this is a superset of the other set
func (m *BitMask) IsSuperSetOf(other *bitset.BitSet) bool {
	return m.bits.IsSuperSet(other)
}

// IsStrictSuperSetOf method - returns true if this is a strict superset of the other set
func (m *BitMask) IsStrictSuperSetOf(other *bitset.BitSet) bool {
	return m.bits.IsStrictSuperSet(other)
}

// IsSubSetOf method - returns true if this is a subset of the other set
func (m *BitMask) IsSubSetOf(other *bitset.BitSet) bool {
	return other.IsSuperSet(m.bits)
}

// IsStrictSubSetOf method - returns true if this is a strict subset of the other set
func (m *BitMask) IsStrictSubSetOf(other *bitset.BitSet) bool {
	return other.IsStrictSuperSet(m.bits)
}

// Rank method - returns the number of set bits up to and including the index that are set in the bitset
func (m *BitMask) Rank(index uint) uint {
	return m.bits.Rank(index)
}

// Select returns the index of the jth set bit, where j is the argument
//
// NOTE: The caller is responsible to ensure that 0 <= j < Count()
//
// WARNING: When j is out of range, the function returns the length of the bitset (b.length)
func (m *BitMask) Select(index uint) uint {
	return m.bits.Select(index)
}

// Bits method - returns the bits of the Mask
func (m *BitMask) Bits() *bitset.BitSet {
	return m.bits
}
