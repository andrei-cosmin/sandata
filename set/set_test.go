package set

import (
	"github.com/andrei-cosmin/sandata/internal/testutil"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	defaultSize    = 1000
	maxRandomValue = 100000
)

func TestSet_FilterFunc(t *testing.T) {
	set := NewSet[uint](defaultSize)
	values := testutil.RandomUInts(10*defaultSize, maxRandomValue)
	for _, value := range values {
		set.Insert(value & 1)
	}

	modified := set.FilterFunc(func(value uint) bool {
		return value%2 == 0
	})

	assert.True(t, modified)

	set.ForEach(func(value uint) {
		assert.True(t, value%2 == 0)
	})
}
