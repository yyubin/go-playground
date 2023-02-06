package main

import (
	"fmt"
	"os"
)

func sum(nums ...int) int { // slice 타입, 빈 인터페이스를 파라미터로 받으면 모든 타입을 받을 수 있다(가변 인수)
	sum := 0

	fmt.Printf("nums 타입: %T\n", nums)
	for _, v := range nums {
		sum += v
	}
	return sum
}

func main() {
	fmt.Println(sum(1, 2, 3, 4, 5))
	fmt.Println(sum(10, 29))
	fmt.Println(sum()) //0

	// defer : 함수 종료 직전 실행
	// 주로 os 자원 반납할때 사용
	f, err := os.Create("test.txt")
	if err != nil {
		fmt.Println("Failed to create a file", err)
		return
	}

	defer fmt.Println("반드시 호출됩니다") // 3
	defer f.Close()                // 파일 자원이 os에 반환됨
	defer fmt.Println("파일을 닫았습니다") // 2
	// defer는 맨 마지막부터 호출 (스택)

	fmt.Println("파일에 Hello World 씁니다") // 1
	fmt.Fprintln(f, "Hello World")     // 출력 스트림 정할 수 있음 api 정보 : https://pkg.go.dev/fmt#Fprintf

	var operator OpFn // 타입이 함수 타입, 인트를 2개 받고 출력으로 인트를 냄
	operator = getOperator("+")

	var result = operator(3, 4) // add(3, 4)와 같음
	fmt.Println(result)         // 7

	// 함수 리터럴(람다)
	var operator2 OpFn
	operator2 = getOperator("-")

	var result2 = operator2(5, 3)
	fmt.Println(result2) // 2

	// 함수 리터럴 내부 상태
	// 일반 함수는 상태를 가질 수 없지만
	// 함수 리터럴은 내부 상태를 가질 수 있다.

	i := 0

	fun := func() { // function임에도 입력값 x, 출력값x
		i += 10 // 외부 변수를 캡쳐해와서 내부에서 사용가능하다 -> 내부 상태의 의미, 내부 상태에서 사용할 수 있다
	}

	i++ // i가 1이 먼저 더해지고

	fun() // 10이 더해짐

	fmt.Println(i) // 11

	// ***캡쳐는 값복사가 아닌 레퍼런스 복사 (go에서는 레퍼런스 없음)
	// -> 포인터 복사

	CaptureLoop()  // 3 3 3
	CaptureLoop2() // 0 1 2

	// 의존성 주입
	f2, err := os.Create("test2.txt")
	if err != nil {
		fmt.Println("Failed to create a file")
		return
	}
	defer f2.Close()

	writeHello(func(msg string) { // 인자로 함수 넣어주기
		fmt.Fprintln(f2, msg)
	})

}

type Writer func(string)
type WriterInterface interface {
	Write(string)
}

func writeHello(writer Writer) { // 파일에 쓰는지, 콘솔에 쓰는지 이 시점에선 알 수 없음
	writer("Hello World") // 외부에서 로직을 주입함 --> 의존성 주입
}

func writeHello2(Writer WriterInterface) { // 인터페이스로 구현시
	Writer.Write("Hello World") // 이 문자열이 어디에 쓰여질지는 모름
}

func CaptureLoop() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop")
	for i := 0; i < 3; i++ {
		f[i] = func() {
			fmt.Println(i)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}
}

func CaptureLoop2() {
	f := make([]func(), 3)
	fmt.Println("ValueLoop2")
	for i := 0; i < 3; i++ {
		v := i
		f[i] = func() {
			fmt.Println(v)
		}
	}

	for i := 0; i < 3; i++ {
		f[i]()
	}

}

type OpFn func(int, int) int // 별칭타입 만들어서 사용

// 함수 타입 변수
// 함수를 값으로 갖는 변수
// 함수 타입은 함수 시그니쳐로 표현

func add(a, b int) int {
	return a + b
}

// 위 함수 시그니처
// func (int, int) int

func mul(a, b int) int {
	return a * b
}

func getOperator(op string) OpFn { // 별칭타입으로 사용
	if op == "+" {
		return add // 함수 add의 주소를 반환
	} else if op == "*" {
		return mul
	} else if op == "-" {
		return func(a, b int) int { // 함수 리터럴
			return a - b
		}
	} else {
		return nil
	}
}
