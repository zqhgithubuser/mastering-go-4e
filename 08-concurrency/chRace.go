package main

import "fmt"

func printer2(ch chan<- bool, times int) {
	for i := 0; i < times; i++ {
		ch <- true
	}
	close(ch)
}

func main() {
	var ch = make(chan bool)
	go printer2(ch, 5)

	for val := range ch {
		fmt.Print(val, " ") // true true true true true
	}
	fmt.Println()

	for i := 0; i < 15; i++ {
		fmt.Print(<-ch, " ")
	}
	fmt.Println()
}
