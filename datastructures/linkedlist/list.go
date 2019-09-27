package linkedlist

import (
	"fmt"
)

type Node struct {
	value interface{}
	next *Node
}

type LinkedList struct {
	Root *Node
	Tail *Node
}

func (ll *LinkedList) Add(v interface{}) {
	// the new node
	t := &Node {v, nil}

	// keep track of the current tail
	if ll.Tail != nil {
		tmp := ll.Tail
		tmp.next = t
	}

	// set the new tail
	ll.Tail = t

	// set the root
	if ll.Root == nil {
		ll.Root = t
	}
}

func (ll *LinkedList) Traverse() {
	if ll.Root == nil {
		fmt.Println("Empty list")
		return
	}
	t := ll.Root
	for t != nil {
		fmt.Printf ("%v ->", t.value)
		t = t.next
	}
	fmt.Println()
}


