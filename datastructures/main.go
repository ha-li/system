package main

import (
	"fmt"
	list "github.com/system/datastructures/linkedlist"
	"github.com/system/datastructures/stack"
)

func reverse () {
	//var myStack *
}

func testLinkedList () {

	myList := new(list.LinkedList)
	myList.Traverse()
	myList.Add(4)
	//myList.Traverse()
	myList.Add(5)
	//myList.Traverse()
	myList.Add(6)
	//myList.Traverse()
	myList.Add(0)
	myList.Add(100)
	//myList.Traverse()
	myList.Add("bruce")
	myList.Traverse()

	myList2 := new(list.LinkedList)
	myList2.Add(1)
	//myList2.Traverse()
	myList2.Add(2)
	//myList2.Traverse()
	myList2.Add(3)
	//myList2.Traverse()
	myList2.Add(4)
	myList2.Add(5)
	myList2.Traverse()

}

func testStack() {

	myStack := new (stack.Stack)

	//var myStack *stack.Node
	myStack.Push(4)
	myStack.Traverse()
	myStack.Push(6)
	myStack.Traverse()
	myStack.Push(8)
	myStack.Traverse()

	fmt.Println(myStack.Pop())
	myStack.Traverse()

	fmt.Println(myStack.Pop())
	myStack.Traverse()

	fmt.Println(myStack.Pop())
	myStack.Traverse()

	fmt.Println(myStack.Pop())
	fmt.Println(myStack.Peek())
	// push onto tmpStck
	/*
	fmt.Println("Push onto tmpStack")
	for tmp != nil {
		fmt.Println(tmp.Value)
		tmpStack.Push(tmp.Value)
		tmp = tmp.Next
	}
	fmt.Println("End push onto tmpStack")

	var revList *list.Node

	fmt.Println("get back values")
	for tmpValue := tmpStack.Pop(); tmpValue != nil; tmpValue = tmpStack.Pop() {
		fmt.Println(tmpValue)
		revList.Add(tmpValue)
		revList.Traverse()
	}
	fmt.Println ("reverse list")
	revList.Traverse()
	*/

}

func main () {
	//testLinkedList()
	testStack()
}