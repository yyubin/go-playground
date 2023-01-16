// go slice
package main

import "fmt"

// func addNum(slice []int) []int {
// 	slice = append(slice, 4)
// 	return slice
// }

// func main() {
// 	slice := []int{1, 2, 3}
// 	slice = addNum(slice)

// 	fmt.Println(slice)
// }

// func main() {
// 	array := [100]int{1: 1, 2: 2, 99: 100}
// 	slice1 := array[1:10]
// 	slice2 := array[2:99]

// 	fmt.Println(slice1)
// 	fmt.Println(slice2)
// }

// func main()  {
// 	slice1 := []int{1, 2, 3, 4, 5}
// 	slice2 := append([]int{}, slice1...)

// 	slice2[1] = 100
// 	fmt.Println("slice1", slice1)
// 	fmt.Println("slice2", slice2)

// }

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))
	copy(slice2, slice1)

	// case 2
	// slice = append(slick[:idx], append([]int{100}, slice[idx:]...)...)

	// case 3
	// slice = append(slice, 0)
	// copy(slice[idx+1:], slice[idx:])
	// slice[idx] = 100

	slice2[1] = 100
	fmt.Println("slice1", slice1) //[1, 2, 3, 4, 5]
	fmt.Println("slice2", slice2) //[1, 100, 3, 4, 5]

}
