package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSet_FilterFunc(t *testing.T) {
	set := NewSet[uint](defaultSize)
	values := getRandomUints(10*defaultSize, maxRandomValue)
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
