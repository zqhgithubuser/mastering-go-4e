package main

import (
	"fmt"
	"sort"
)

type Size struct {
	F1 int
	F2 string
	F3 int
}

type Person struct {
	F1 int
	F2 string
	F3 Size
}

type PersonSlice []Person

func (a PersonSlice) Len() int {
	return len(a)
}

func (a PersonSlice) Less(i, j int) bool {
	return a[i].F3.F1 < a[j].F3.F1
}

func (a PersonSlice) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func main() {
	data := []Person{
		Person{1, "One", Size{1, "Person_1", 10}},
		Person{2, "Two", Size{2, "Person_2", 20}},
		Person{-1, "Two", Size{-1, "Person_3", -20}},
	}

	fmt.Println("Before:", data)
	sort.Sort(PersonSlice(data))
	fmt.Println("After:", data)
	// 逆序
	sort.Sort(sort.Reverse(PersonSlice(data)))
	fmt.Println("Reverse:", data)
}
