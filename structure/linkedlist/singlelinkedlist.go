package linkedlist

import "fmt"

type Node[T any] struct {
	next  *Node[T] // 다음 노드를 가리키는 포인터
	Value T        // 노드의 값
}

type LinkedList[T any] struct {
	root  *Node[T] // 첫 번째 노드를 가리키는 포인터
	tail  *Node[T] // 마지막 노드를 가리키는 포인터
	count int
}

func (l *LinkedList[T]) PushBack(v T) {
	node := &Node[T]{Value: v}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}

	l.tail.next = node
	l.tail = node
}

func (l *LinkedList[T]) PushFront(v T) {
	node := &Node[T]{Value: v}
	l.count++
	if l.root == nil {
		l.root = node
		l.tail = node
		return
	}
	node.next = l.root
	l.root = node
}

func (l *LinkedList[T]) Front() *Node[T] {
	return l.root
}

func (l *LinkedList[T]) Back() *Node[T] {
	return l.tail
}

func (l *LinkedList[T]) Count() int {
	node := l.root
	cnt := 0

	for ; node != nil; node = node.next {
		cnt++
	}
	return cnt
}

func (l *LinkedList[T]) Count2() int {
	return l.count
}

func (l *LinkedList[T]) GetAt(idx int) *Node[T] {
	if idx >= l.Count2() {
		return nil
	}
	i := 0
	for node := l.root; node != nil; node = node.next {
		if i == idx {
			return node
		}
		i++
	}
	return nil
}

func (l *LinkedList[T]) InsertAfter(node *Node[T], value T) {
	if !l.isIncluded(node) {
		return
	}
	newNode := &Node[T]{Value: value}
	newNode.next, node.next = node.next, newNode
	l.count++
}

func (l *LinkedList[T]) isIncluded(node *Node[T]) bool {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner == node {
			return true
		}
	}
	return false
}

func (l *LinkedList[T]) InsertBefore(node *Node[T], value T) {
	if node == l.root {
		l.PushFront(value)
		return
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}
	newNode := &Node[T]{Value: value}
	prevNode.next, newNode.next = newNode, node
	l.count++
}

func (l *LinkedList[T]) findPrevNode(node *Node[T]) *Node[T] {
	inner := l.root
	for ; inner != nil; inner = inner.next {
		if inner.next == node {
			return inner
		}
	}
	return nil
}

func (l *LinkedList[T]) PopFront() {
	if l.root == nil {
		return
	}
	l.root.next, l.root = nil, l.root.next
	if l.root == nil {
		l.tail = nil
	}
	l.count--
}

func (l *LinkedList[T]) Remove(node *Node[T]) {
	if node == l.root {
		l.PopFront()
		return
	}
	prevNode := l.findPrevNode(node)
	if prevNode == nil {
		return
	}
	prevNode.next, node.next = node.next, nil
	if node == l.tail {
		l.tail = prevNode
	}
	l.count--
}

func main() {
	fmt.Println("Hello, world!")
}
