package main

// log will automatically print the date and time
import (
	"errors"
	"fmt"
	"log"
)
// errors is a special type
func division(x, y int) (int, error, error) {
	if y == 0 {
		return 0, nil, errors.New("Division by zero error.")
	}

	if x%y == 0 {
		remainer := errors.New("Remainder error.")
		return x/y, remainer, nil
	}

	return x/y, nil, nil
}

func main() {
	result, remainder, err := division(2, 0)
	if err != nil {
		log.Fatal (err)
	} else {
		fmt.Println ("Result ", result)
	}
	if remainder != nil {
		fmt.Print(remainder)
	}
}