package main

import (
    "fmt"
    "sort"
)

func main() {
    sInts := []int{1, 0, 2, -3, 4, -20}
    sFloats := []float64{1.0, 0.2, 0.22, -3, 4.1, -0.1}
    sStrings := []string{"aa", "a", "A", "Aa", "aab", "AAa"}
    fmt.Println("sInts original:", sInts) // [1 0 2 -3 4 -20]
    sort.Ints(sInts)
    fmt.Println("sInts:", sInts) // [-20 -3 0 1 2 4]

    sort.Sort(sort.Reverse(sort.IntSlice(sInts)))
    fmt.Println("Reverse:", sInts) // [4 2 1 0 -3 -20]

    fmt.Println("sFloats original:", sFloats)
    sort.Float64s(sFloats)
    fmt.Println("sFloats:", sFloats)
    sort.Sort(sort.Reverse(sort.Float64Slice(sFloats)))
    fmt.Println("Reverse:", sFloats)

    fmt.Println("sStrings original:", sStrings)
    sort.Strings(sStrings)
    fmt.Println("sStrings:", sStrings) //  [A AAa Aa a aa aab]
    sort.Sort(sort.Reverse(sort.StringSlice(sStrings)))
    fmt.Println("Reverse:", sStrings) // [aab aa a Aa AAa A]
}
