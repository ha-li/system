package main

// Golang supports reflection. The API is found in the package 'reflect'.
// The below example takes an unknown struct and outputs it struct name as
// well as each member type, name and value


import (
	"fmt"
	"reflect"
)

type Message struct {
	X     int
	Y     int
	Label string
}

func main () {
	p1 := Message {4, 2, "A Message" }
	p2 := Message {}
	p2.Label = "Message 2"

	s1 := reflect.ValueOf(&p1).Elem()
	fmt.Println("s1=", s1)
	fmt.Println("p1=", p1)

	t := s1.Type()
	fmt.Printf("%v(struct) {\n", t.Name())
	for i:=0; i < s1.NumField(); i++ {
		f := s1.Field(i)

		fmt.Printf("   %s(%s) = %v\n", t.Field(i).Name, f.Type(), f.Interface())
		//fmt.Printf("= %v\n", f.Interface())
	}
	fmt.Println("}")
}
