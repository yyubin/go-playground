package main

import "fmt"

// 패닉은 처리하기 힘든 에러 만났을 때 프로그램 조기 종료
func divide(a, b int) {
	if b == 0 {
		panic("b는 0일 수 없습니다")
		// 어디에서 종료되었는지 확인 가능
	}
	fmt.Printf("%d / %d = %d\n", a, b, a/b)
}

func f() {
	fmt.Println("f() 함수 시작")
	defer func() { // defer + 함수 리터럴
		if r := recover(); r != nil { // 리커버 시도
			fmt.Println("패닉 복구 - ", r)
		}
	}()

	g()
	fmt.Println("f() 함수 끝")
}

func g() {
	fmt.Printf("9/3 = %d\n", h(9, 3))
	fmt.Printf("9/0 = %d\n", h(9, 0))
}

func h(a, b int) int {
	if b == 0 {
		panic("제수는 0일 수 없습니다.")
	}
	return a / b
}

func main() {
	// divide(9, 3)
	// divide(9, 0)
	f()
	fmt.Println("프로그램이 계속 실행됨")

	// 실행 결과
	// f() 함수 시작
	// 9/3 = 3
	// 패닉 복구 -  제수는 0일 수 없습니다.
	// 프로그램이 계속 실행됨
}

// 패닉의 전파
// 콜스택 역순으로 전파된다
// h() 에서 패닉 발생
// g() 복구 확인 -> 전파
// f() 복구 확인 -> 전파
// main() 복구 확인 -> 전파
// 복구 할 시에는 recover()객체 사용
// 패닉 객체 반환
// defer 사용

// 사용시에는 바로 recover 사용하지 않고 오류 고치도록 해야함
