package data

import (
	"github.com/bits-and-blooms/bitset"
	"github.com/stretchr/testify/assert"
	"math/rand/v2"
	"slices"
	"testing"
)

func TestArray_Set(t *testing.T) {
	keys := getRandomUints(numArrayElements, maxRandomValue)
	array := NewArray[uint](defaultSize)

	keys[0] = 100
	for index := range numArrayElements {
		array.Set(keys[index], keys[index])
	}

	for index := range numArrayElements {
		assert.Equal(t, keys[index], array.Get(keys[index]))
	}
}

func TestArray_Clear(t *testing.T) {
	keys := getRandomUints(numArrayElements, maxRandomValue)
	array := NewArray[uint](defaultSize)

	bits := bitset.New(defaultSize)
	mask := NewMask(bits)

	for index := range numArrayElements {
		array.Set(keys[index], keys[index])
		bits.Set(keys[index])
	}

	array.ClearAll(mask)

	for index := range numArrayElements {
		keys[index] = rand.UintN(numArrayElements)
		assert.Equal(t, uint(0), array.Get(keys[index]))
	}
}

func TestArray_ClearAllFunc(t *testing.T) {
	keys := getRandomUints(numArrayElements, maxRandomValue)
	array := NewArray[uint](defaultSize)

	bits := bitset.New(defaultSize)
	mask := NewMask(bits)

	for index := range numArrayElements {
		array.Set(keys[index], keys[index])
		bits.Set(keys[index])
	}
	slices.Sort(keys)

	cursor := 0
	array.ClearAllFunc(mask, func(value uint) {
		assert.Equal(t, keys[cursor], value)
		cursor++
	})

	for index := range numArrayElements {
		keys[index] = rand.UintN(numArrayElements)
		assert.Equal(t, uint(0), array.Get(keys[index]))
	}
}
