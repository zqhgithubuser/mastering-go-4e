package main

import "fmt"

func main() {
	// 创建空前片
	var aSlice []float64
	fmt.Println(aSlice, len(aSlice), cap(aSlice)) // [] 0 0
	// 添加元素
	aSlice = append(aSlice, 1234.56)
	aSlice = append(aSlice, -34.0)
	fmt.Println(aSlice, "with length", len(aSlice)) // [1234.56 -34] with length 2

	// 初始长度
	t := make([]int, 4)
	t[0] = -1
	t[1] = -2
	t[2] = -3
	t[3] = -4
	t = append(t, -5)
	fmt.Println(t) // [-1 -2 -3 -4 -5]

	// 二维切片
	twoD := [][]int{{1, 2, 3}, {4, 5, 6}}
	// 遍历
	for _, i := range twoD {
		for _, k := range i {
			fmt.Print(k, " ")
		}
		fmt.Println()
	}

	// 未初始化
	make2D := make([][]int, 2)
	fmt.Println(make2D) // [[] []]
	make2D[0] = []int{1, 2, 3, 4}
	make2D[1] = []int{-1, -2, -3, -4}
	fmt.Println(make2D) // [[1 2 3 4] [-1 -2 -3 -4]]
}
