package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 20; i++ {
		i := i
		go func() {
			fmt.Print(i, " ")
		}()
	}
	time.Sleep(time.Second)
	fmt.Println()

	// 更好的做法
	for i := 0; i < 20; i++ {
		go func(x int) {
			fmt.Print(x, " ")
		}(i)
	}
	time.Sleep(time.Second)
	fmt.Println()
}
