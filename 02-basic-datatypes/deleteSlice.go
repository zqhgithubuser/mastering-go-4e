package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Need an integer value.")
		return
	}

	index := arguments[1]
	i, err := strconv.Atoi(index)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Using index", i)
	aSlice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("Original slice:", aSlice)

	// 索引超过总长度
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}
	aSlice = append(aSlice[:i], aSlice[i+1:]...)
	fmt.Println("After 1st deletion:", aSlice)

	// Delete element at index i
	if i > len(aSlice)-1 {
		fmt.Println("Cannot delete element", i)
		return
	}
	aSlice[i] = aSlice[len(aSlice)-1]
	aSlice = aSlice[:len(aSlice)-1]
	fmt.Println("After 2nd deletion:", aSlice)
}
