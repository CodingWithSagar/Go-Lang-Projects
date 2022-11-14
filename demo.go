package main

import "fmt"

func main() {
	var a chan int
	if a == nil {
		fmt.Println("channel a is nil, going to define it")
		a = make(chan int)
		fmt.Printf("Type of a is %T value %v", a, a)
	}
	// a := make(chan int)
	//a <-11
	//fmt.Println(a)
}
