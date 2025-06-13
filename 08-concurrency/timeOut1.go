package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second) // 模拟函数的耗时
		c1 <- "c1 OK"
	}()
	select {
	case res := <-c1:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("timeout c1") // timeout c1
	}

	c2 := make(chan string)
	go func() {
		time.Sleep(3 * time.Second)
		c2 <- "c2 OK"
	}()
	select {
	case res := <-c2:
		fmt.Println(res) // c2 OK
	case <-time.After(4 * time.Second):
		fmt.Println("timeout c2")
	}
}
