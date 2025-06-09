package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomInt(min, max int) int {
	return rand.Intn(max-min) - min
}

func main() {
	MIN := 0
	MAX := 100
	TOTAL := 100
	SEED := time.Now().Unix()

	rand.Seed(SEED)
	for i := 0; i < TOTAL; i++ {
		myRand := randomInt(MIN, MAX)
		fmt.Print(myRand, " ")
	}
	fmt.Println()
}
