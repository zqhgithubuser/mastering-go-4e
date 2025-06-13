package main

import (
	"fmt"
	"sync"
)

func writeToChannel(c chan int, x int) {
	c <- x
	close(c)
}

func printer(ch chan bool) {
	ch <- true
}

func main() {
	c := make(chan int, 1)

	var wg sync.WaitGroup
	wg.Add(1)
	go func(c chan int) {
		defer wg.Done()
		writeToChannel(c, 10)
		fmt.Println("Exit.")
	}(c)
	fmt.Println("Read:", <-c) // Read: 10

	_, ok := <-c
	if ok {
		fmt.Println("Channel is open!")
	} else {
		fmt.Println("Channel is closed!") // Channel is closed!
	}
	wg.Wait()

	var ch = make(chan bool)
	for i := 0; i < 5; i++ {
		go printer(ch)
	}

	n := 0
	for i := range ch {
		fmt.Println(i)
		if i == true {
			n++
		}
		if n > 2 {
			fmt.Println("n:", n) // n: 3
			close(ch)
			break
		}
	}
}
