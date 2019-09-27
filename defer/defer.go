package main

/*
  simple go app to show how defer key work works.

  defer
  - takes a function call and delays its execution until the surrounding function ends
  - is LIFO, so the calls that are deferred last are called first

 */

import (
	"fmt"
)

func a1 () {
	// illustrates the lifo aspect of defer
	for i:=0; i<3; i++ {
		defer fmt.Print (i, " ")
	}
}

func a2 () {
	// illustrates that deferred functions only get called after
	// the rest of the function completes, since the anonymous function
	// get created after the rest of the for loop ends, i will be 3
	// when the anonymous functions are created
	for i:=0; i<3; i++ {
		defer func () { fmt.Print(i, " "); }()
	}
}

func a3 () {
	// this is a combination of a1 and a2,
	// from a2, we know that the anonymous function will be created at the end
	// but i is passed in at each iteration, so we will get 2, 1, 0 like a1
	for i:=0; i<3; i++ {
		defer func(n int) { fmt.Print(n, " ") } (i)
	}
}

func main () {
	a1()
	fmt.Println()
	a2()
	fmt.Println()
	a3()
}
