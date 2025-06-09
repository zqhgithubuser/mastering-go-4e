package main

import (
	"io"
	"log"
	"os"
)

func main() {
	flag := os.O_APPEND | os.O_CREATE | os.O_WRONLY
	file, err := os.OpenFile("myLog.log", flag, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	w := io.MultiWriter(file, os.Stderr)
	logger := log.New(w, "myApp: ", log.LstdFlags)
	logger.Printf("BOOK %d", os.Getpid())
}
