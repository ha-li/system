package main

import (
	"fmt"
	"os"
	"strconv"
)


func main () {
	args := os.Args
	sum := 0
	for i := 1; i < len(args); i++ {
		val,err := strconv.Atoi(args[i])
		if err != nil {
			fmt.Printf ("%v is not an integer, skipping\n", args[i])
			continue;
		}
		sum += val
	}
	fmt.Println (sum)
}