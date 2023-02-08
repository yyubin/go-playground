// 마지막 요소가 다시 처음 요소로 연결된 링 형태의 자료구조 -> 환형 버퍼
package main

import (
	"container/ring"
	"fmt"
)

func main() {
	r := ring.New(5)
	n := r.Len()

	for i := 0; i < n; i++ {
		r.Value = 'A' + i
		r = r.Next()
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%c ", r.Value) // A B C D E
		r = r.Next()
	}
	fmt.Println()

	for i := 0; i < n; i++ {
		fmt.Printf("%c ", r.Value) // A E D C B
		r = r.Prev()
	}

	// 일정한 갯수만 사용하고 오래된 요소가 지워져도 되는 경우에 사용
	// 실행 취소 기능
	// 고정 크기 버퍼 기능
	// 리플레이 기능
}
