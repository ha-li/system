package main

import (
	"fmt"
	"os"
)

func main () {
	args := os.Args
	fmt.Printf("Hello World, %v!", args[1])
}
