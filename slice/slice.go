package main

import "fmt"

// slices are passed by reference automatically for you
// so it is very fast to use slices
func printSlice(s1 []int) {
	for _,v:=range s1 {
		fmt.Println (v);
	}
}

func main () {
	aSlice := []int{-1, 1, 2, 3, 4,6}
	printSlice(aSlice)
}
