package stack

import (
	"fmt"
)

type Node struct {
	Value interface{}
	Next *Node
}

type Stack struct {
	Top *Node
}

// get a look at the element on the top
// of the stack, without popping it off
func (s *Stack) Peek () interface{} {
	return s.Top
}

// add an element to the top of the stack
func (s *Stack) Push(v interface{}) {
	t := &Node{v, nil}

	// the first node on the stack
	if s.Top == nil {
		s.Top = t
		return
	}

	t.Next = s.Top
	s.Top = t
}

// remove the top of the stack and return it
func (s *Stack) Pop () interface{} {
	tmp := s.Top

	if tmp != nil {
		s.Top = tmp.Next
		return tmp.Value
	}
	return nil
}



func (s *Stack) Traverse() {
	tmp := s.Top

	if s.Top == nil {
		fmt.Println( "Empty Stack")
		return
	}

	for tmp != nil {
		fmt.Printf ("%v -> ", tmp.Value)
		tmp = tmp.Next
	}
	fmt.Println()
}
