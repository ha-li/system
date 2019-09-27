package main

import (
	"fmt"
	"os"
	"strings"
)

func main () {
	args := os.Args
	scani := false
	for i := 0; i < len(args); i++ {
		if strings.Compare(args[i], "-i") == 0 {
			scani = true
			break
		}
	}

	if scani {
		fmt.Println("Got the -i parameter")
		fmt.Print ("y/n: ")
		var input string
		fmt.Scanln(&input)
		fmt.Println ("you entered ", input)
	}
}
