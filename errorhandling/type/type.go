package main

import "fmt"

type PasswordError struct {
	Len        int
	RequireLen int
}

func (err PasswordError) Error() string {
	return "암호 길이가 짧습니다."
}

func RegisterAccount(name, password string) error {
	if len(password) < 8 {
		// 필요한 정보 구조체로 만들어서 사용자 타입의 에러 반환
		// 에러는 인터페이스
		return PasswordError{len(password), 8}
	}
	return nil
}

func main() {
	err := RegisterAccount("yubin", "1")
	if err != nil {
		if errInfo, ok := err.(PasswordError); ok {
			fmt.Printf("%v Len:%d RequrieLen:%d \n", errInfo, errInfo.Len, errInfo.RequireLen)
		}
	} else {
		fmt.Println("회원 가입 성공")
	}

}
