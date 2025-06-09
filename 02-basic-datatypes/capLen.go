package main

import "fmt"

func main() {
	// Length=4 Capacity=4
	a := make([]int, 4)
	fmt.Println("L:", len(a), "C:", cap(a))
	// Length=5 Capacity=5
	b := []int{0, 1, 2, 3, 4}
	fmt.Println("L:", len(b), "C:", cap(b))

	// 显式指定容量
	aSlice := make([]int, 4, 4)
	fmt.Println(aSlice) // [0 0 0 0]

	aSlice = append(aSlice, 5)
	fmt.Println(aSlice)                               // [0 0 0 0 5]
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice)) // L: 5 C: 8
	// 增加多个元素
	aSlice = append(aSlice, []int{-1, -2, -3, -4}...)
	fmt.Println(aSlice)                               // [0 0 0 0 5 -1 -2 -3 -4]
	fmt.Println("L:", len(aSlice), "C:", cap(aSlice)) // L: 9 C: 16
}
