package pool

import (
	"github.com/andrei-cosmin/sandata/internal/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
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
