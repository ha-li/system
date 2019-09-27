package main

import (
	"fmt"
)

type Fee struct {
	Amount int
	Description string
}

type People struct {
	Age int
	Name string
}

func main () {
	objs := make([]interface{}, 3)
	objs[0] = "bool"
	objs[1] = true
	objs[2] = 5

	objs = append(objs, Fee {1, "Tax"} )
	objs = append(objs, People {45, "John Paul"} )

	for _,s:=range objs {
		fmt.Println(s)
	}
}
