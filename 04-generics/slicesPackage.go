package main

import (
	"fmt"
	"slices"
)

func main() {
	s1 := []int{1, 2, -1, -2}
	s2 := slices.Clone(s1)
	s3 := slices.Clone(s1[2:])
	fmt.Println(s1, s2, s3) // [1 2 -1 -2] [1 2 -1 -2] [-1 -2]

	s1[2] = 0
	s2[3] = 0
	fmt.Println(s1, s2, s3) // [1 2 0 -2] [1 2 -1 0] [-1 -2]

	s1 = slices.Compact(s1)
	fmt.Println("s1 (compact):", s1)
	fmt.Println(slices.Contains(s1, 2), slices.Contains(s1, -2)) // true true

	s4 := make([]int, 10, 100)
	fmt.Println("Len:", len(s4), "Cap:", cap(s4)) // Len: 10 Cap: 100
	// 删除未使用的 Capability
	s4 = slices.Clip(s4)
	fmt.Println("Len:", len(s4), "Cap:", cap(s4)) // Len: 10 Cap: 10

	fmt.Println("Min:", slices.Min(s1), "Max:", slices.Max(s1))

	fmt.Println("s2:", s2)                  // [1 2 -1 0]
	s2 = slices.Replace(s2, 1, 3, 100, 200) // [1 100 200 0]
	fmt.Println("s2 (replaced):", s2)
	slices.Sort(s2)
	fmt.Println("s2 (sorted):", s2)
}
