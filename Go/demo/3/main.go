package main

import (
	"container/list"
	"fmt"
)

func main() {
	s3 := []int{1, 2, 3, 4, 5, 6, 7, 8}
	s4 := s3[3:6]

	fmt.Println("s3", len(s3))
	fmt.Println("s3", &s3)
	fmt.Println("s4", len(s3))
	fmt.Println("s4", &s4)
	fmt.Println("s3", cap(s3))
	fmt.Println("s4", cap(s4)) // = cap(s3) - 3


	var l list.List
	fmt.Println("l.Len()", l.Len())

	fmt.Println(l.PushBack(0))
	fmt.Println(l.PushBack("12"))
	fmt.Println(l.PushFront(2))
	fmt.Println(l.PushFront(3))

	fmt.Println("l", l)
	fmt.Println("l.Len()", l.Len())

	fmt.Println("Front", l.Front().Value)
	fmt.Println("Front.Next", l.Front().Next().Value)
	fmt.Println("Back", l.Back().Value)

	var e list.Element
	e.Value = "test"

	fmt.Println(l.InsertAfter("testFront", &e))
	fmt.Println(l.InsertAfter("testFront", l.Back()))
}
