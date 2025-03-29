package chain

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var listKeys = []int{1, 2, 3, 4, 5, 6, 7}
var otherKeys = []int{101, 102, 103, 104, 105, 106, 107}

func TestNode_Remove(t *testing.T) {
	node := New[int](listKeys).Next.Next
	assert.Equal(t, listKeys[2], node.Data)

	listNode := node.Next

	node.RemoveNode()
	assert.False(t, node.HasPrevious())
	assert.False(t, node.HasNext())
	assert.True(t, node.IsHead())
	assert.True(t, node.IsTail())

	index := 0
	cursor := HeadOf[int](listNode)
	assert.True(t, cursor.IsHead())
	assert.False(t, cursor.HasPrevious())

	for cursor != nil {
		if index == 2 {
			index++
		}
		assert.Equal(t, listKeys[index], cursor.Data)
		cursor = cursor.Next
		index++
	}

	cursor = TailOf[int](listNode)
	assert.True(t, cursor.IsTail())
	assert.False(t, cursor.HasNext())

	for cursor != nil {
		index--
		if index == 2 {
			index--
		}
		assert.Equal(t, listKeys[index], cursor.Data)
		cursor = cursor.Prev
	}
}

func TestNode_InsertAsHead(t *testing.T) {
	nodeHead := New[int](listKeys).Next.Next
	nodeTail := New[int](otherKeys).Next.Next

	newHead := nodeHead.Next
	newTail := nodeTail.Prev

	nodeHead.InsertAsHead(nodeTail)
	assert.True(t, newHead.IsHead())
	assert.True(t, newTail.IsTail())
	originalHead := HeadOf[int](nodeHead)
	originalTail := TailOf[int](nodeHead)
	assert.Equal(t, originalHead.Data, listKeys[0])
	assert.Equal(t, originalTail.Data, otherKeys[len(otherKeys)-1])
}

func TestNode_Split(t *testing.T) {
	splitNode := New[int](listKeys).Next.Next
	newTail := splitNode.Prev
	splitNode.Split()
	assert.True(t, newTail.IsTail())
	assert.True(t, splitNode.IsHead())
}

func TestNode_InsertAsTail(t *testing.T) {
	nodeHead := New[int](listKeys).Next.Next
	nodeTail := New[int](otherKeys).Next.Next

	newHead := nodeHead.Next
	newTail := nodeTail.Prev

	nodeTail.InsertAsTail(nodeHead)
	assert.True(t, newHead.IsHead())
	assert.True(t, newTail.IsTail())
	originalHead := HeadOf[int](nodeHead)
	originalTail := TailOf[int](nodeHead)
	assert.Equal(t, originalHead.Data, listKeys[0])
	assert.Equal(t, originalTail.Data, otherKeys[len(otherKeys)-1])
}
