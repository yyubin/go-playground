package main

import "fmt"

func sqaure(x int) int { // 첫번째 테케만 통과하도록 작성
	// if x == 3 { // 다음 테케만 통과하도록 작성
	// 	return 9
	// }
	// return 81
	return x * x // 이제 세번째 테케도 통과하도록 수정
}

func fibonacci1(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return n
	}
	return fibonacci1(n-1) + fibonacci1(n-2)
}

func fibonacci2(n int) int {
	if n < 0 {
		return 0
	}
	if n < 2 {
		return 2
	}
	one := 1
	two := 0
	rst := 0
	for i := 2; i <= n; i++ {
		rst = one + two
		two = one
		one = rst
	}
	return rst

}

func main() {
	fmt.Println(fibonacci1(13))
	fmt.Println(fibonacci2(13))
}
