package main

import "fmt"

func main () {

	threeD := [2][2][2]int {
		{
			{1, 2},
			{3, 4},
		},
		{
			{5, 6},
			{7, 8},
		},
	}
	threeD[0][1][1] = -1
	for _, number := range threeD {
		fmt.Printf("%d ", number)
	}
	fmt.Printf("\n")
}
