package main

import (
    "fmt"
    "maps"
)

// 删除所有奇数
func del(k string, v int) bool {
    return v%2 != 0
}

func equal(v1 int, v2 float64) bool {
    return float64(v1) == v2
}

func main() {
    m := map[string]int{
        "one": 1, "two": 2,
        "three": 3, "four": 4,
    }
    maps.DeleteFunc(m, del)
    fmt.Println(m) // map[four:4 two:2]

    n := maps.Clone(m)
    if maps.Equal(m, n) {
        fmt.Println("Equal!") // Equal!
    } else {
        fmt.Println("Not equal!")
    }

    n["three"] = 3
    n["two"] = 22
    fmt.Println("Before n:", n, "m:", m) // Before n: map[four:4 three:3 two:22] m: map[four:4 two:2]
    maps.Copy(m, n)
    fmt.Println("After n:", n, "m:", m) // After n: map[four:4 three:3 two:22] m: map[four:4 three:3 two:22]

    t := map[string]int{
        "one": 1, "two": 2,
        "three": 3, "four": 4,
    }
    mFloat := map[string]float64{
        "one": 1.00, "two": 2.00,
        "three": 3.00, "four": 4.00,
    }
    eq := maps.EqualFunc(t, mFloat, equal)
    fmt.Println("Is t equal to mFloat?", eq)
}
