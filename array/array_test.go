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
	"math/rand/v2"
	"slices"
	"testing"

	"github.com/andrei-cosmin/sandata/bit"
	"github.com/andrei-cosmin/sandata/internal/testutil"
	"github.com/bits-and-blooms/bitset"
	"github.com/stretchr/testify/assert"
)

const (
	defaultSize      = 1000
	numArrayElements = 10000
	maxRandomValue   = 100000
)

func TestArray_Set(t *testing.T) {
	keys := testutil.RandomUInts(numArrayElements, maxRandomValue)
	array := New[uint](defaultSize)

	keys[0] = 100
	for index := range numArrayElements {
		array.Set(keys[index], keys[index])
	}

	for index := range numArrayElements {
		assert.Equal(t, keys[index], array.Get(keys[index]))
	}
}

func TestArray_Clear(t *testing.T) {
	keys := testutil.RandomUInts(numArrayElements, maxRandomValue)
	array := New[uint](defaultSize)

	bits := bitset.New(defaultSize)
	mask := bit.NewMask(bits)

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
	keys := testutil.RandomUInts(numArrayElements, maxRandomValue)
	array := New[uint](defaultSize)

	bits := bitset.New(defaultSize)
	mask := bit.NewMask(bits)

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
