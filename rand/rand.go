package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init () {
	// this does not need to be called multiple times,
	// just once per execution of program will suffice to
	// seed the random number generator
	rand.Seed(time.Now().Unix())
}

func random(min, max int) int {

	return rand.Intn(max-min) + min
}

func main () {
	for i := 0; i < 100; i++ {
		fmt.Print( random(1, 100))
		fmt.Print( " ")
	}
}
