package main

import (
	"container/list"
	"fmt"
)

type Queue struct {
	v *list.List
}

func (q *Queue) Push(val interface{}) { //
	q.v.PushBack(val)
}

func (q *Queue) Pop() interface{} { // 배열로 구현시 O(n) 리스트로 구현시 O(1)
	front := q.v.Front()
	if front != nil {
		return q.v.Remove(front)
	}
	return nil
}

func NewQueue() *Queue {
	return &Queue{list.New()}
}

func main() {
	v := list.New()
	e4 := v.PushBack(4)
	e1 := v.PushFront(1)

	v.InsertBefore(3, e4)
	v.InsertAfter(2, e1)

	for e := v.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()

	for e := v.Back(); e != nil; e = e.Prev() {
		fmt.Print(e.Value, " ")
	}
	// https://pkg.go.dev/container/list

	// Big-O
	// O(1) < O(N) < O(NlogN) < O(N^2) < O(N^3)

	// 큐
	queue := NewQueue()

	for i := 1; i < 5; i++ {
		queue.Push(i)
	}

	// v := queue.Pop()
	// for v != nil {
	// 	fmt.Printf("%v -> ", v)
	// 	v = queue.Pop()
	// }

}
