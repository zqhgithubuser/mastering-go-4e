package main

import "fmt"

func main() {
	willClose := make(chan complex64, 10)
	// Write some data to the channel
	willClose <- -1
	willClose <- 1i
	<-willClose
	<-willClose
	close(willClose)
	//Read again - this is a closed channel
	read := <-willClose
	fmt.Println(read) // (0+0i)
}
