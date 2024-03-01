package data

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var listKeys = []int{1, 2, 3, 4, 5, 6, 7}
var otherKeys = []int{101, 102, 103, 104, 105, 106, 107}

func TestChainNode_Remove(t *testing.T) {
	node := newChain[int](listKeys).NextChain.NextChain
	assert.Equal(t, listKeys[2], node.Data)

	listNode := node.NextChain

	node.RemoveNode()
	assert.False(t, node.HasPrevious())
	assert.False(t, node.HasNext())
	assert.True(t, node.IsHead())
	assert.True(t, node.IsTail())

	index := 0
	cursor := getHead[int](listNode)
	assert.True(t, cursor.IsHead())
	assert.False(t, cursor.HasPrevious())

	for cursor != nil {
		if index == 2 {
			index++
		}
		assert.Equal(t, listKeys[index], cursor.Data)
		cursor = cursor.NextChain
		index++
	}

	cursor = getTail[int](listNode)
	assert.True(t, cursor.IsTail())
	assert.False(t, cursor.HasNext())

	for cursor != nil {
		index--
		if index == 2 {
			index--
		}
		assert.Equal(t, listKeys[index], cursor.Data)
		cursor = cursor.PreviousChain
	}
}

func TestChainNode_InsertAsHead(t *testing.T) {
	nodeHead := newChain[int](listKeys).NextChain.NextChain
	nodeTail := newChain[int](otherKeys).NextChain.NextChain

	newHead := nodeHead.NextChain
	newTail := nodeTail.PreviousChain

	nodeHead.InsertAsHead(nodeTail)
	assert.True(t, newHead.IsHead())
	assert.True(t, newTail.IsTail())
	originalHead := getHead[int](nodeHead)
	originalTail := getTail[int](nodeHead)
	assert.Equal(t, originalHead.Data, listKeys[0])
	assert.Equal(t, originalTail.Data, otherKeys[len(otherKeys)-1])
}

func TestChainNode_Split(t *testing.T) {
	splitNode := newChain[int](listKeys).NextChain.NextChain
	newTail := splitNode.PreviousChain
	splitNode.Split()
	assert.True(t, newTail.IsTail())
	assert.True(t, splitNode.IsHead())
}

func TestChainNode_InsertAsTail(t *testing.T) {
	nodeHead := newChain[int](listKeys).NextChain.NextChain
	nodeTail := newChain[int](otherKeys).NextChain.NextChain

	newHead := nodeHead.NextChain
	newTail := nodeTail.PreviousChain

	nodeTail.InsertAsTail(nodeHead)
	assert.True(t, newHead.IsHead())
	assert.True(t, newTail.IsTail())
	originalHead := getHead[int](nodeHead)
	originalTail := getTail[int](nodeHead)
	assert.Equal(t, originalHead.Data, listKeys[0])
	assert.Equal(t, originalTail.Data, otherKeys[len(otherKeys)-1])
}
