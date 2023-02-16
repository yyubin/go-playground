package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSquare1(t *testing.T) {
	assert.Equal(t, 81, square(9), "sqaure(9) should be 81")
	// rst := square(9)
	// if rst != 81 {
	// 	t.Errorf("sqaure(9) should be 81 but returns %d", rst)
	// }
}

func TestSquare2(t *testing.T) {
	assert.Equal(t, 9, square(3), "sqaure(3) should be 9")
	// rst := square(3)
	// if rst != 9 {
	// 	t.Errorf("sqaure(3) should be 9 but returns %d", rst)
	// }
}

func TestSquare3(t *testing.T) {
	assert.Equal(t, 0, square(0), "sqaure(0) should be 0")
	// rst := square(0)
	// if rst != 0 {
	// 	t.Errorf("sqaure(0) should be 0 but returns %d", rst)
	// }
}

// 테스트 주도 개발(TDD)
// 기존 문제 : 테스트가 중요한데 테스트 케이스가 부족하거나 형식적인 경우가 많았음.
// 기존 : 코드 작성 -> 테스트 -> 코드 수정 -> 완성
// 예상 가능한 테스트 케이스만 적용해 보게 됨

// 테스트 케이스가 촘촘할 수록 버그 방지에 유리
// ->> 테스트 주도 개발
// 테스트 작성 -> 테스트 실패 -> 코드 작성 -> 테스트 코드 -> 완성
// 첫 테스트는 무조건 실패하는데 실패한 테스트 케이스를 통과시킬 수 있도록 코드 작성
// 새로운 테스트 코드 적용, 다시 코드 작성 및 리팩토링 반복
