package linkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPushBack(t *testing.T) {
	var l LinkedList[int]
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushBack(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushBack(2)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 2, l.Back().Value)

	l.PushBack(3)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)

	assert.Equal(t, 3, l.Count())

	assert.Equal(t, 1, l.GetAt(0).Value)
	assert.Equal(t, 2, l.GetAt(1).Value)
	assert.Nil(t, l.GetAt(4))
}

func TestPushFront(t *testing.T) {
	var l LinkedList[int]
	assert.Nil(t, l.root)
	assert.Nil(t, l.tail)
	l.PushFront(1)

	assert.NotNil(t, l.root)
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(2)
	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	l.PushFront(3)
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 3, l.Count())

	l.PushFront(4)
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)

	assert.Equal(t, 4, l.Count())
	assert.Equal(t, 4, l.Count2())
}

func TestInsertAfter(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertAfter(node, 4) // 1, 2, 4, 3

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)
}

func TestInsertBefore(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(1)
	l.InsertBefore(node, 4) // 1, 4, 2, 3

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.GetAt(1).Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{Value: 5}
	l.InsertBefore(tempNode, 200)
	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.Count())
}

func TestInsertRootBefore(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)

	node := l.GetAt(0)
	l.InsertBefore(node, 4) // 4, 1, 2, 3

	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.Front().Value)
	assert.Equal(t, 2, l.GetAt(2).Value)
	assert.Equal(t, 3, l.Back().Value)

	tempNode := &Node[int]{Value: 5}
	l.InsertBefore(tempNode, 200)
	assert.Equal(t, 4, l.Count2())
	assert.Equal(t, 4, l.Count())
}

func TestRemove(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.Remove(l.GetAt(1))

	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())

	l.Remove(l.GetAt(1))
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())

	l.Remove(&Node[int]{Value: 100})
	assert.Equal(t, 1, l.Front().Value)
	assert.Equal(t, 1, l.Back().Value)
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())

	l.Remove(l.GetAt(0))
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 0, l.Count2())
}

func TestPopFront(t *testing.T) {
	var l LinkedList[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.PopFront()

	assert.Equal(t, 2, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 2, l.Count())
	assert.Equal(t, 2, l.Count2())

	l.PopFront()
	assert.Equal(t, 3, l.Front().Value)
	assert.Equal(t, 3, l.Back().Value)
	assert.Equal(t, 1, l.Count())
	assert.Equal(t, 1, l.Count2())

	l.PopFront()
	assert.Nil(t, l.Front())
	assert.Nil(t, l.Back())
	assert.Equal(t, 0, l.Count())
	assert.Equal(t, 0, l.Count2())
}
