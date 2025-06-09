package main

import (
	"fmt"
)

func change(s []string) {
	s[0] = "Change_function"
}

func main() {
	a := [4]string{"Zero", "One", "Two", "Three"}
	fmt.Println("a:", a)

	var S0 = a[0:1]
	fmt.Println(S0) // [Zero]
	S0[0] = "S0"
	fmt.Println("a:", a) // a: [S0 One Two Three]

	var S12 = a[1:3]
	fmt.Println(S12) // [One Two]
	S12[0] = "S12_0"
	S12[1] = "S12_1"
	fmt.Println("a:", a) // a: [S0 S12_0 S12_1 Three]

	change(S12)
	fmt.Println("a:", a) // a: [S0 Change_function S12_1 Three]

	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0)) // Capacity of S0: 4 Length of S0: 1
	S0 = append(S0, "N1")
	S0 = append(S0, "N2")
	S0 = append(S0, "N3")
	a[0] = "-N1"
	fmt.Println(a) // [-N1 N1 N2 N3]

	// 添加该元素后，S0的底层数组不再是a
	S0 = append(S0, "N4")
	fmt.Println("Capacity of S0:", cap(S0), "Length of S0:", len(S0)) // Capacity of S0: 8 Length of S0: 5
	a[0] = "-N1-"
	a[1] = "-N2-"
	fmt.Println(a) // [-N1- -N2- N2 N3]

	fmt.Println("S0:", S0)   // S0: [-N1 N1 N2 N3 N4]
	fmt.Println("S12:", S12) // [-N2- N2]
}
