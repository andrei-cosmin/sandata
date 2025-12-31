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

package pool

import (
	"testing"

	"github.com/andrei-cosmin/sandata/internal/testutil"
	"github.com/stretchr/testify/assert"
)

const (
	defaultSize    = 1000
	maxRandomValue = 100000
)

func TestPool_SizeCapacity(t *testing.T) {
	pool := New[uint](defaultSize)
	assert.True(t, pool.Empty())
	assert.Equal(t, 0, pool.Size())
	assert.Equal(t, defaultSize, pool.Capacity())

	values := testutil.RandomUInts(10*defaultSize, maxRandomValue)
	for _, value := range values {
		pool.Push(value)
	}
	assert.False(t, pool.Empty())
	assert.Equal(t, defaultSize, pool.Size())
	assert.Equal(t, defaultSize, pool.Capacity())
}

func TestPool_PopEmpty(t *testing.T) {
	pool := New[uint](defaultSize)
	value, ok := pool.Pop()
	assert.False(t, ok)
	assert.Equal(t, uint(0), value)
}

func TestPool_PushPop(t *testing.T) {
	pool := New[uint](defaultSize)

	values := testutil.RandomUInts(defaultSize/2, maxRandomValue)
	for _, value := range values {
		pool.Push(value)
	}

	cursor := defaultSize/2 - 1
	for !pool.Empty() {
		value, ok := pool.Pop()
		assert.True(t, ok)
		assert.Equal(t, values[cursor], value)
		cursor--
	}
}

func TestPool_PushPop_Full(t *testing.T) {
	pool := New[uint](defaultSize)

	values := testutil.RandomUInts(10*defaultSize, maxRandomValue)
	for _, value := range values {
		pool.Push(value)
	}

	cursor := defaultSize - 1
	for !pool.Empty() {
		value, ok := pool.Pop()
		assert.True(t, ok)
		assert.Equal(t, values[cursor], value)
		cursor--
	}
}
