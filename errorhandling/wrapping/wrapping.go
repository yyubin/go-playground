package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func MultipleFromString(str string) (int, error) {
	scanner := bufio.NewScanner(strings.NewReader(str)) // 인자로 reader 필요
	scanner.Split(bufio.ScanWords)

	pos := 0
	a, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Faild to readNextInt(), pos : %d err : %w", pos, err)
	}

	pos += n + 1
	b, n, err := readNextInt(scanner)
	if err != nil {
		return 0, fmt.Errorf("Faild to readNextInt(), pos : %d err : %w", pos, err)
	}

	return a * b, nil
}

// 다음 단어를 읽어서 숫자로 변환하여 반환
// 변환된 숫자, 읽은 글자 수, 에러 반환
func readNextInt(scanner *bufio.Scanner) (int, int, error) {
	if !scanner.Scan() {
		return 0, 0, fmt.Errorf("Failed to scan")
	}

	word := scanner.Text()
	number, err := strconv.Atoi(word)
	if err != nil {
		return 0, 0, fmt.Errorf("Failed to convert word to int, word:%s, err:%w", word, err)
	}
	return number, len(word), nil
}

func readEq(eq string) {
	rst, err := MultipleFromString(eq)
	if err == nil {
		fmt.Println(rst)
	} else {
		fmt.Println(err)
		var numError *strconv.NumError
		if errors.As(err, &numError) {
			// errors.As에서 wrapping된 err에 numError 객체가 있는지 확인
			// is는 있는지 확인, as는 데이터까지 꺼내옴
			fmt.Println("NumberError", numError)
		}
	}
}

func main() {
	readEq("123 3")
	readEq("123 abc")
}
