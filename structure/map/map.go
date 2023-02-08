package main

import "fmt"

type Product struct {
	Name  string
	Price int
}

const M = 10

func hash(d int) int {
	return d % M
}

func main() {
	// m := make(map[string]string) // map[string]string

	// m["이화랑"] = "서울시 광진구"
	// m["송화랑"] = "서울시 강남구"
	// m["조화랑"] = "부산시 사하구"
	// m["고화랑"] = "청주시 상당구"
	// m["고화랑"] = "전주시 덕진구"

	// fmt.Printf("송화랑의 주소는 %s입니다\n", m["송화랑"])
	// fmt.Printf("고화랑의 주소는 %s입니다\n", m["고화랑"])

	// m := make(map[int]Product)

	// m[16] = Product{"카드", 100}
	// m[54] = Product{"티슈", 200}
	// m[61] = Product{"펜", 300}
	// m[222] = Product{"마우스", 400}
	// m[345] = Product{"키보드", 5000}

	// for k, v := range m {
	// 	fmt.Println(k, v) // 정렬 보장 X
	// }

	// hash map(unordered map), sorted map 차이
	// sorted map의 경우에는 키가 정렬됨
	// hashmap은 정렬 X

	// 요소 삭제와 존재 여부
	// delete(m, key)
	// v, ok := m[3]
	// fmt.Println(v, ok)

	// map1 := make(map[int]int)
	// map1[1] = 0
	// map1[2] = 2
	// map1[3] = 3

	// delete(map1, 3)
	// delete(map1, 4)
	// fmt.Println(map1[3]) // 없으면 그냥 기본값 출력
	// fmt.Println(map1[1])

	// v, ok := map1[3]
	// fmt.Println(v, ok)

	// 		   배열   리스트  맵
	//    추가  O(n)  O(1)  O(1)
	//	  삭제  O(n)  O(1)  O(1)
	//    읽기  O(1)  O(n)  O(1)
	//	  각각 인덱스, 키로 접근시 기준

	// 맵이 빠르지만 순회에 순서보장 X, 메모리를 더 많이 소모

	// 맵의 원리
	// 해쉬 함수 동작; 해쉬: 잘게 부순다
	// 1. 같은 입력이 들어오면 같은 결과가 나온다.
	// 2. 다른 입력이 들어오면 되도록 다른 결과가 나온다.
	// 3. 입력값의 범위는 무한대이고 결과는 특정 범위를 갖는다.
	// 나머지 연산 주로 사용(해쉬 함수로 활용)

	m := [M]string{}
	m[hash(23)] = "애쉬" // 넣을때도 일정한 시간
	m[hash(259)] = "케틀"

	fmt.Printf("%d = %s\n", 23, m[hash(23)]) // 출력시에도 일정한 시간이 걸린다, 나머지 함수(해시 함수 사용)
	fmt.Printf("%d = %s\n", 259, m[hash(259)])
	// 23을 넣으면 3나오는데 배열의 인덱스 3 위치에 값을 대입

	// 해쉬 충돌
	// ex) hash 23의 값과 hash 33의 값이 같음
	// 하나의 값만을 넣지 않고 리스트를 넣어 키값으로 비교해서 가져올 수 있도록 함(이 때에도 리스트를 순회하기는 함)
	// 각 요소가 리스트를 가짐

	// 해쉬 사용하는 곳
	// file -> checksum

}
