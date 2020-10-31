package main

import "fmt"

var n = "hello"

func main() {
	var n = "hello2"
	{
		var n = "hello3"
		fmt.Println("hello3:", n)
	}
	fmt.Println("hello2:", n)

	//n := "hello2"
	{
		n := 1
		fmt.Println("hello3:", n)
	}
	fmt.Println("hello2:", n)
	fmt.Println("--------------------------")
	test()
}

func test() {
	var i *int
	i = new(int)
	*i = 10
	fmt.Println(*i)
}
