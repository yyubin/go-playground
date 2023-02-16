package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSqaure(t *testing.T) { // 테스트부터 작성
	assert.Equal(t, 81, sqaure(9), "should be 81")
	assert.Equal(t, 9, sqaure(3), "should be 9")
	assert.Equal(t, 16, sqaure(4), "should be 16")
	assert.Equal(t, 16, sqaure(-4), "should be 16") // 테스트 먼저 작성할 시 더 촘촘한 테케 만들 수 있음
}

// TDD 장점
// 1. 테이스케이스가 자연스럽게 늘어난다
// 2. 테스트가 촘촘해진다
// 3. 자연스러운 회기 테스트가 가능하다
// 4. 리팩토링이 쉬워진다
//		개선전 코드의 동작과 개선후 코드 동작이 같아야함
// 5. 개발이 즐겁다
// 6. 코드 커버리지가 자연히 증가된다
// 		테스트 코드가 커버하는 영역이 증가한다

// TDD 단점
// 1. 모듈간 의존성 높을 경우 테스트 케이스 만들기 힘들다
//		 의존성을 끊거나 mock 데이터(테스트용 객체)로 테스트
// 2. 동시성 테스트에 취약하다
// 3. 진정한 TDD가 아닌 형식적인 테스트로 전락할 수 있다
// 4. 지속적인 모니터링과 관리가 필요하다
//

func BenchmarkFibonacci1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci1(20)
	}
	// 38122             31345 ns/op
}

func BenchmarkFibonacci2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fibonacci2(20)
	}
	// 166420030                7.206 ns/op
}
